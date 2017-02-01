// Implements a FTP server.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Laugusti/gopl/ch8/exercise_8_2/ftpsession"
)

var port = flag.Uint("port", 8000, "port for ftp server")

func main() {
	flag.Parse()
	laddr := fmt.Sprintf(":%d", *port)
	listener, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go ftpsession.NewFTPSession(conn).HandleConn()
	}
}
