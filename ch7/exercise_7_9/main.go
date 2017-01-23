package main

import (
	"html/template"
	"log"
	"net/http"
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

var trackList = template.Must(template.New("trackList").Parse(`
<table style="width:100%">
<tr style='text-align: left'>
	<th><a href="?sort=title">Title</a></th>
	<th><a href="?sort=artist">Artist</a></th>
	<th><a href="?sort=album">Album</a></th>
	<th><a href="?sort=year">Year</a></th>
	<th><a href="?sort=length">Length</a></th>
</tr>
{{range .}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
`))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Query().Get("sort") {
	case "title":
		sort.Sort(tracks.NewTrackSorter(t, tracks.SortByTitle))
	case "artist":
		sort.Sort(tracks.NewTrackSorter(t, tracks.SortByArtist))
	case "album":
		sort.Sort(tracks.NewTrackSorter(t, tracks.SortByAlbum))
	case "year":
		sort.Sort(tracks.NewTrackSorter(t, tracks.SortByYear))
	case "length":
		sort.Sort(tracks.NewTrackSorter(t, tracks.SortByLength))
	}

	if err := trackList.Execute(w, t); err != nil {
		log.Print(err)
	}
}
