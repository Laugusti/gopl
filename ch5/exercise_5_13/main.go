package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Laugusti/gopl/ch5/links"
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

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	} else {
		err = saveLocally(url)
		if err != nil {
			log.Print(err)
		}
	}
	return withSameDomain(url, list)
}

func stripSchemeFromURL(url string) string {
	// only http and https
	switch {
	case strings.HasPrefix(url, "http://"):
		url = url[7:]
	case strings.HasPrefix(url, "https://"):
		url = url[8:]
	}
	if strings.HasPrefix(url, "www.") {
		url = url[4:]
	}
	return url
}

func getDomainFromURL(url string) string {
	url = stripSchemeFromURL(url)
	index := strings.Index(url, "/")
	if index != -1 {
		return url[:index]
	} else {
		return url
	}
}

func withSameDomain(url string, list []string) []string {
	i := 0
	domain := getDomainFromURL(url)
	for _, u := range list {
		if strings.HasPrefix(stripSchemeFromURL(u), domain) {
			list[i] = u
			i++
		}
	}
	return list[:i]
}

func saveLocally(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	fpath, fname := getPathAndNameFromURL(url)
	err = os.MkdirAll(fpath, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(fpath + string(os.PathSeparator) + fname)
	if err != nil {
		return err
	}

	io.Copy(file, resp.Body)
	file.Close()
	resp.Body.Close()
	return nil
}

func getPathAndNameFromURL(url string) (fpath, fname string) {
	url = stripSchemeFromURL(url)
	fpath, fname = filepath.Split(url)
	if fpath == "" {
		fpath, fname = fname, "index.html"
	}
	return
}
