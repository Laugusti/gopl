// Following the approach of mirroredQuery in Section 8.4.4, implement a variant
// of fetch that requests several URLs concurrently. As soon as the first response
// arrives, cancel the othe requests
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var cancel = make(chan struct{})

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	// create GET request using cancel channel
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	req.Cancel = cancel

	// Read request response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	// close cancel channel
	select {
	case <-cancel:
	default:
		close(cancel)
	}

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
