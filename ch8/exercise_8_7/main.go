package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Laugusti/gopl/ch5/links"
	"github.com/Laugusti/gopl/ch8/exercise_8_7/htmlutil"
)

var maxDepth = flag.Int("depth", 3, "crawl depth")

type linkdepth struct {
	link  string
	depth int
}

func linkdepthSlice(links []string, depth int) []linkdepth {
	list := make([]linkdepth, len(links))
	for i := range links {
		list[i] = linkdepth{links[i], depth}
	}
	return list
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	// TODO: uncomment
	err = htmlutil.SaveUrl(url)
	if err != nil {
		log.Print(err)
	}
	return htmlutil.WithSameDomain(url, list)
}

func main() {
	flag.Parse()
	worklist := make(chan []linkdepth)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- linkdepthSlice(flag.Args(), 0) }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, ld := range list {
			if ld.depth < *maxDepth && !seen[ld.link] {
				seen[ld.link] = true
				n++
				go func(link string, depth int) {
					worklist <- linkdepthSlice(crawl(link), depth)
				}(ld.link, ld.depth+1)
			}
		}
	}
}
