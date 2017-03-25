// Write a concurrent program that creates a local mirror of a web site, fetchin
// each reachable page and writing it to a directory on the local disk. Only pages
// within the original domain (for instance, golang.org should be fetched. URLs within
// mirrored pages should be altered as needed so that they refer to the mirrored page,
// no the original.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/Laugusti/gopl/ch5/links"
	"github.com/Laugusti/gopl/ch8/exercise_8_7/htmlutil"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{} // acquire a token

	err := htmlutil.SaveUrl(url)
	if err != nil {
		log.Print(err)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	<-tokens // release the token
	return htmlutil.WithSameDomain(url, list)
}

func main() {
	flag.Parse()
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:2] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			// crawl link if unseen
			if !seen[link] {
				seen[link] = true
				n++
				// crawl the link, and add the result to the worklist with depth+1
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
