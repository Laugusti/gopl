package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Laugusti/gopl/ch7/exercise_7_4/stringhelper"
	"golang.org/x/net/html"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	//fmt.Printf("%s\n\n", b)
	if err != nil {
		log.Fatal(err)
	}
	r := stringhelper.NewReader(string(b))
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
