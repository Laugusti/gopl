package main

import (
	"log"
	"net/http"

	"github.com/Laugusti/gopl/ch3/exercise_3_4/surfacePlot"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		surfaceplot.SurfacePlot(w, r)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
