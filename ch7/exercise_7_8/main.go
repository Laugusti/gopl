package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/Laugusti/gopl/ch7/exercise_7_8/tracks"
)

var t = []*tracks.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	s := tracks.NewTrackSorter(t, tracks.SortByYear)

	sort.Sort(sort.Reverse(s))
	tracks.PrintTracks(t)
	fmt.Println("\n")

	s.SortBy = tracks.SortByTitle
	sort.Sort(s)
	tracks.PrintTracks(t)
}
