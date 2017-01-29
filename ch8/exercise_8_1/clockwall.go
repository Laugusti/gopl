// Clockwall acts as a client of serveral clock servers at once, reading the
// times from each one and displaying the results in a table.
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func exitWithUsage() {
	fmt.Println("Usage: clockwall NAME=SERVER...")
	os.Exit(1)
}

// parseArgs parse command line argumes to a map of server names to server.
func parseArgs() map[string]string {
	m := make(map[string]string)

	for _, arg := range os.Args[1:] {
		s := strings.Split(arg, "=")
		if len(s) != 2 {
			exitWithUsage()
		}
		m[s[0]] = s[1]
	}
	if len(m) < 1 {
		exitWithUsage()
	}
	return m
}

func main() {
	serverAddrMap := parseArgs()

	// get server names (maps are random order)
	var serverNames []string
	for k := range serverAddrMap {
		serverNames = append(serverNames, k)
	}

	// create a seperate buffer for each server to write data
	serverBufMap := make(map[string]*bytes.Buffer)
	for _, server := range serverNames {
		serverBufMap[server] = &bytes.Buffer{}
		fmt.Printf("%-8s\t", server)
		// create a goroutine for each clock server with a
		// seperate buffer
		go dial(serverBufMap[server], serverAddrMap[server])
	}
	fmt.Println()

	for {
		fmt.Printf("\r")
		// for each clock server, read from its buffer and
		// print to stdout
		for _, server := range serverNames {
			s, err := serverBufMap[server].ReadString(byte('\n'))
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			if s != "" {
				fmt.Printf("%s\t", s[:len(s)-1])
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// dial connects to the address and writes the data from the connection to the
// provided buffer
func dial(buf *bytes.Buffer, laddr string) {
	conn, err := net.Dial("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(buf, conn)
}

// mustCopy copies the data from the reader to the writer and exits if there is an error
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
