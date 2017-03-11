// Add depth-limiting to the concurrent crawler. That is, if the user sets -depth =3,
// then only URLs reachable by at most three links will be fetched
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Laugusti/gopl/ch5/links"
)

var maxDepth = flag.Int("depth", 3, "maximum depth")

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()
	worklist := make(chan []linkdepth)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- makeLDSlice(flag.Args(), 0) }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, ld := range list {
			// if link is unseen and depth is less than max, add to worklist
			if !seen[ld.link] {
				seen[ld.link] = true
				fmt.Println(ld.link)
				if ld.depth < *maxDepth {
					n++
					// crawl the link, and add the result to the worklist with depth+1
					go func(link string, depth int) {
						worklist <- makeLDSlice(crawl(link), depth)
					}(ld.link, ld.depth+1)
				}
			}
		}
	}
}
