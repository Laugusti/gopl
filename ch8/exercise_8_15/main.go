// Failure of any client program to read data in a timely manner ultimately causes all
// clients to get stuck. Modify the broadcaster to skip a message rather than wait if a
// client writer is not ready to accept it. Alternatively, add buffering to each client's
// outgoing message channel so that most messages are not dropped; the broadcaster should
// use a non-blocking send to this channel.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	name string
	ch   chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

const (
	maxInactivity    = 5 * time.Minute
	clientQueueDepth = 3
)

func announceClients(ch chan<- string, clients map[client]bool) {
	msg := "Current clients: ["
	var comma bool
	for cli := range clients {
		if comma {
			msg += ", "
		} else {
			comma = true
		}
		msg += cli.name
	}
	msg += "]"
	ch <- msg
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming messages to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
					// client is not ready for message, skip
				}
			}
		case cli := <-entering:
			announceClients(cli.ch, clients)
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignore network errors
	}
}

func getName(conn net.Conn) (string, error) {
	fmt.Fprint(conn, "Please provide a name:") // NOTE: ignore network errors
	input := bufio.NewScanner(conn)
	input.Scan()
	if input.Err() != nil {
		return "", input.Err()
	}
	return input.Text(), nil

}

func handleConn(conn net.Conn) {
	ch := make(chan string, clientQueueDepth) // outgoing client messages
	who, err := getName(conn)
	if err != nil {
		log.Print(err)
		return
	}
	go clientWriter(conn, ch)

	cli := client{who, ch}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)

	// create channel/goroutine for input.Scan()
	msgChannel := make(chan string)
	go func() {
		for input.Scan() {
			msgChannel <- input.Text()
		}
		// NOTE: ignoring potential errors from input.Err()
		close(msgChannel)
	}()

loop:
	for {
		// wait until maxInactivity for goroutine to send data to channel
		select {
		case msg, ok := <-msgChannel:
			if ok {
				messages <- who + ": " + msg
			} else {
				break loop
			}
		case <-time.After(maxInactivity):
			fmt.Fprintf(conn, "timeout after %v of inactivity", maxInactivity) // NOTE: ignore network errors
			break loop
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
