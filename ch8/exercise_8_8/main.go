// Using a select statement, add a timeout to the echo server from Section 8.3 so
// that it disconnects any client that shouts nothing within 10 seconds.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scanInput(input *bufio.Scanner, out chan<- bool) {
}

func handleConn(c net.Conn) {
	// NOTE: ignoring potential errors from input.Err()
	defer c.Close()

	input := bufio.NewScanner(c)

	// create channel/goroutine for input.Scan()
	ch := make(chan bool)
	go func() {
		for input.Scan() {
			ch <- true
		}
		close(ch)
	}()

	for {
		// wait 10 seconds for goroutine to send data to channel
		select {
		case hasData := <-ch:
			if hasData {
				go echo(c, input.Text(), 1*time.Second)
			} else {
				return
			}
		case <-time.After(10 * time.Second):
			fmt.Println("timeout after 10 seconds of inactivity")
			return
		}
	}

}
