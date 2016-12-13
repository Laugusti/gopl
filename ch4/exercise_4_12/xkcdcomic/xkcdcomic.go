package xkcdcomic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const url = "https://xkcd.com/:num:/info.0.json"

type Xkcd struct {
	Number     int `json:"num"`
	Title      string
	Transcript string
	Alt        string
	Image      string `json:"img"`
	Day        string
	Month      string
	Year       string
	Url        string
}

func FetchAllComics() ([]*Xkcd, error) {
	var result []*Xkcd
	for i := 1; ; i++ {
		if i == 404 {
			continue
		}
		url := strings.Replace(url, ":num:", strconv.Itoa(i), 1)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusNotFound {
				resp.Body.Close()
				break
			}
			resp.Body.Close()
			return nil, fmt.Errorf("failed to retrieve comic #%d; status: %s", i, resp.Status)
		}
		var comic Xkcd
		if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()
		comic.Url = url[:len(url)-len("/info.0.json")]
		result = append(result, &comic)
	}
	return result, nil
}

func LoadComicsFromFile(filename string) ([]*Xkcd, error) {
	var result []*Xkcd
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		data := s.Bytes()
		var comic Xkcd
		if err := json.Unmarshal(data, &comic); err != nil {
			return nil, err
		}
		result = append(result, &comic)
	}
	return result, nil
}

func SaveComicsToFile(comics []*Xkcd, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	for _, c := range comics {
		data, err := json.Marshal(*c)
		if err != nil {
			return err
		}
		data = append(data, '\n')
		_, err = w.Write(data)
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}
