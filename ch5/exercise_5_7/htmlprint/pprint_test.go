package htmlprint

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const url = "https://gopl.io"

func TestPrettyPrint(t *testing.T) {
	doc, err := getDocument(url)
	if err != nil {
		t.Fatal(err)
	}
	formatted := PrettyPrint(doc)
	_, err = html.Parse(strings.NewReader(formatted))
	if err != nil {
		t.Fail()
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
