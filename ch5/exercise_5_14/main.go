package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Laugusti/gopl/ch5/exercise_5_14/explorefs"
)

func main() {
	// Crawl the web breadth-first.
	// Starting form the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(path string) []string {
	list, err := explorefs.GetSubDirectories(path)
	if err != nil {
		log.Print(err)
	} else {
		printDir(path)
	}
	return list
}

func printDir(path string) {
	fmt.Println(path)
	files, err := explorefs.GetFiles(path)
	if err != nil {
		log.Print(err)
	} else {
		for _, f := range files {
			fmt.Printf("\t%s\n", f)
		}
	}
}
