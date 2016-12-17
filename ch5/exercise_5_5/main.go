package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("Words: %d, Images: %d\n", words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words an images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (word, images int) {
	word = countWords(n)
	images = countImages(n)
	return
}

func countImages(n *html.Node) (count int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		count++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count += countImages(c)
	}
	return
}

func countWords(n *html.Node) (count int) {
	if n.Type == html.TextNode {
		count += countWordsInString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count += countWords(c)
	}
	return
}
func countWordsInString(s string) (count int) {
	in := bufio.NewScanner(strings.NewReader(s))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		count++
	}
	return
}
