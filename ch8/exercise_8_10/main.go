// HTTP requests may be cancelled by closing the optional Cancel channel in the
// http.Request struct. Modify the web crawler of Section 8.6 to support cancellation.
package main

import (
	"fmt"
	"log"
	"os"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

var cancel = make(chan struct{})

func cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}
func crawl(url string) []string {
	if cancelled() {
		return nil
	}
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := extract(url, cancel)
	<-tokens // release the token
	if err != nil && !cancelled() {
		log.Print(err)
	}
	return list
}

func main() {
	// Cancel when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancel)
	}()

	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
loop:
	for ; n > 0; n-- {
		select {
		case <-cancel:
			// Drain worklist
			for ; n > 0; n-- {
				<-worklist
			}
			break loop
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
}
