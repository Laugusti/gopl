// Make the chat server disconnect idle clients, such as those that have sent no
// messages in the last five minutes.
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

const maxInactivity = 5 * time.Minute

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
				cli.ch <- msg
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

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
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
			ch <- fmt.Sprintf("timeout after %v of inactivity", maxInactivity)
			// send empty string to client message channel to ensure last message
			// is received and processed before connection is closed
			ch <- ""
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
