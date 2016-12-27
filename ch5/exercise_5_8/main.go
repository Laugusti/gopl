package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Laugusti/gopl/ch5/exercise_5_7/htmlprint"
	"github.com/Laugusti/gopl/ch5/exercise_5_8/htmldom"

	"golang.org/x/net/html"
)

func main() {
	doc, err := getDocument(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	el := htmldom.ElementByID(doc, os.Args[2])
	if el == nil {
		fmt.Println("matching element not found", os.Args[2])
		os.Exit(1)
	} else {
		fmt.Println(htmlprint.PrettyPrint(el))
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
