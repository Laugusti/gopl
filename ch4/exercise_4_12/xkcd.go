package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Laugusti/gopl/ch4/exercise_4_12/xkcdcomic"
)

var filename = flag.String("filename", "buildindex/index.json", "Index file")
var title = flag.String("title", "", "Comic title")
var transcript = flag.String("transcript", "", "Comic transcript")
var number = flag.String("number", "", "Comic number")

func main() {
	flag.Parse()
	checkArgs()

	comics, err := xkcdcomic.LoadComicsFromFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	if *title != "" {
		comics = filterByTitle(comics, *title)
	}
	if *transcript != "" {
		comics = filterByTranscript(comics, *transcript)
	}
	if isNumber(*number) {
		num, _ := strconv.Atoi(*number)
		comics = filterByNumber(comics, num)
	}

	if len(comics) == 0 {
		fmt.Println("No matching comics found")
	} else {
		fmt.Println("URL\tTranscript")
	}
	for _, c := range comics {
		fmt.Printf("%s\t%s\n", c.Url, c.Transcript)
	}
}

func filterByTitle(comics []*xkcdcomic.Xkcd, title string) []*xkcdcomic.Xkcd {
	var result []*xkcdcomic.Xkcd
	for _, c := range comics {
		if strings.Contains(strings.ToLower(c.Title), strings.ToLower(title)) {
			result = append(result, c)
		}
	}
	return result
}

func filterByTranscript(comics []*xkcdcomic.Xkcd, transcript string) []*xkcdcomic.Xkcd {
	var result []*xkcdcomic.Xkcd
	for _, c := range comics {
		if strings.Contains(strings.ToLower(c.Transcript), strings.ToLower(transcript)) {
			result = append(result, c)
		}
	}
	return result
}

func filterByNumber(comics []*xkcdcomic.Xkcd, number int) []*xkcdcomic.Xkcd {
	var result []*xkcdcomic.Xkcd
	for _, c := range comics {
		if c.Number == number {
			result = append(result, c)
		}
	}
	return result
}

func checkArgs() {

	if *title == "" && *transcript == "" && *number == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *number != "" && !isNumber(*number) {
		fmt.Println("Invalid Number")
		flag.Usage()
		os.Exit(1)
	}
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
