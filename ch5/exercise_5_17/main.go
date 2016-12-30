package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Laugusti/gopl/ch5/exercise_5_17/htmldom"
	"github.com/Laugusti/gopl/ch5/exercise_5_7/htmlprint"
	"golang.org/x/net/html"
)

func main() {
	doc, err := getDocument(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	images := htmldom.ElementsByTagName(doc, "img")
	headings := htmldom.ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	fmt.Println("Images:")
	for _, image := range images {
		i := strings.Replace(htmlprint.PrettyPrint(image), "\n", "", -1)
		fmt.Printf("\t%s\n", i)
	}

	fmt.Println("\n\nHeadings:")
	for _, heading := range headings {
		h := strings.Replace(htmlprint.PrettyPrint(heading), "\n", "", -1)
		fmt.Printf("\t%s\n", h)
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
