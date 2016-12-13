package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"net/url"
	"os"
)

const omdbUrl = "https://omdbapi.com/"

var filename = flag.String("filename", "poster.jpg", "filename for poster")

type SearchResult struct {
	Movies       []*Movie `json:"Search"`
	TotalResults int      `json:",string"`
	Resonse      bool     `json:",string"`
}

type Movie struct {
	Title  string
	Year   int `json:",string"`
	Type   string
	Poster string
}

func main() {
	flag.Parse()
	title := getTitleArgument()

	posterUrl, err := getPosterUrl(title)
	checkErr(err)

	poster, err := getPosterImage(posterUrl)
	checkErr(err)

	file, err := os.Create(*filename)
	checkErr(err)

	err = jpeg.Encode(file, *poster, nil)
	checkErr(err)
}

func getPosterImage(url string) (*image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get poster image")
	}
	image, _, err := image.Decode(resp.Body)
	resp.Body.Close()
	return &image, err
}

func getPosterUrl(title string) (string, error) {
	query := "?s=" + url.QueryEscape(title)
	resp, err := http.Get(omdbUrl + query)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return "", err
	}
	resp.Body.Close()
	if result.TotalResults < 1 {
		return "", fmt.Errorf("no matching movie found")
	} else {
		if result.Movies[0].Poster == "N/A" {
			return "", fmt.Errorf("poster image is not available")
		}
		return result.Movies[0].Poster, nil
	}
}

func getTitleArgument() string {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s movie-title\n", os.Args[0])
		os.Exit(1)
	}
	return os.Args[1]
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
