package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func incomingURLs() []string {
	return []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}
}

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)

}

func waitFor(duration string, done <-chan struct{}) (interface{}, error) {
	d, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}
	select {
	case <-done:
		return "cancelled", nil
	case <-time.After(d):
		return "completed", nil
	}
}

func TestMemo(t *testing.T) {
	m := New(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

func TestCancel(t *testing.T) {
	tests := []struct {
		waitTime   time.Duration
		cancelTime time.Duration
		cancelled  bool
	}{
		{200 * time.Millisecond, 100 * time.Millisecond, true},
		{200 * time.Millisecond, 300 * time.Millisecond, false},
		{200 * time.Millisecond, 100 * time.Millisecond, false},
	}

	m := New(waitFor)
	for _, test := range tests {
		cancel := make(chan struct{})
		go func(cancel chan struct{}) {
			<-time.After(test.cancelTime)
			close(cancel)
		}(cancel)
		_, err := m.Get(test.waitTime.String(), cancel)
		if err == CancelledError && !test.cancelled || test.cancelled && err != CancelledError {
			t.Fail()
		}
	}
}
