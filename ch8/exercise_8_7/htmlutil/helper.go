package htmlutil

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

// lock used to only allow a single thread to create a file at a time
var lock = make(chan struct{}, 1)

// WithSameDomain returns elements of list with the same domain as url.
func WithSameDomain(url string, list []string) []string {
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

// SaveUrl saves the contents of url locally.
func SaveUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	var isHtml bool
	if ct := resp.Header.Get("Content-Type"); ct == "text/html" || strings.HasPrefix(ct, "text/html;") {
		isHtml = true
	}

	return writeToFile(url, resp.Body, isHtml)
}

func writeToFile(url string, r io.Reader, isHtml bool) error {
	// acquire lock at start and release at end
	lock <- struct{}{}
	defer func() { <-lock }()

	fpath, fname := getPathAndNameFromURL(url, isHtml)
	file, err := createFile(fpath, fname)

	if err != nil {
		return fmt.Errorf("creating file %s: %s", url, err)
	}
	defer file.Close()

	// check Content.Type is HTML (e.g., "text/html; charset=utf-8").
	if isHtml {
		doc, err := html.Parse(r)
		if err != nil {
			return fmt.Errorf("creating html doc for %s: %s", url, err)
		}

		_, err = file.WriteString(PrettyPrint(url, doc))
		if err != nil {
			return fmt.Errorf("writing %s: %s", url, err)
		}
	} else {
		_, err := io.Copy(file, r)
		if err != nil {
			return fmt.Errorf("writing %s: %s", url, err)
		}
	}
	return nil
}

// create a file using the filepath and filename
func createFile(fpath, fname string) (*os.File, error) {
	err := os.MkdirAll(fpath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(fpath, fname))
}

// removes the scheme from the url (inluding www.)
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

// returns the domain of the url
func getDomainFromURL(url string) string {
	url = stripSchemeFromURL(url)
	index := strings.Index(url, "/")
	if index != -1 {
		return url[:index]
	} else {
		return url
	}
}

// splits a url into a filepath and filename
func getPathAndNameFromURL(url string, isHtml bool) (string, string) {
	url = strings.TrimRight(stripSchemeFromURL(url), "/")

	// if content type is html and not end in html (use index.html)
	if isHtml && !strings.HasSuffix(url, ".html") {
		return url, "index.html"
	}

	fpath, fname := filepath.Split(url)
	if fpath == "" {
		fpath, fname = fname, "index.html"
	}
	if fname == "" {
		fname = "index.html"
	}
	return fpath, fname
}
