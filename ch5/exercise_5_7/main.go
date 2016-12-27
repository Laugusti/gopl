package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Laugusti/gopl/ch5/exercise_5_7/htmlprint"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		doc, err := getDocument(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Print(htmlprint.PrettyPrint(doc))
	}
}

func getDocument(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}
