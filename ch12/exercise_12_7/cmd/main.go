package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Laugusti/gopl/ch12/exercise_12_7"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     0,
		Color:    true,
		Actor: map[string]string{
			"Dr. Srangelove":             "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	buf := &bytes.Buffer{}
	if err := sexpr.NewEncoder(buf).Encode(strangelove); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
