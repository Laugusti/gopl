package main

import (
	"flag"
	"log"

	"github.com/Laugusti/gopl/ch4/exercise_4_12/xkcdcomic"
)

var filename = flag.String("filename", "index.json", "Offline index")

func main() {
	flag.Parse()
	comics, err := xkcdcomic.FetchAllComics()
	checkErr(err)
	err = xkcdcomic.SaveComicsToFile(comics, *filename)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
