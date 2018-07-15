package main

import (
	"fmt"
	"log"

	"github.com/Laugusti/gopl/ch12/exercise_12_10"
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
	var strangelove interface{} = Movie{
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

	b, err := sexpr.Marshal(&strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshalled output:")
	fmt.Println(string(b))

	strangelove = nil
	sexpr.RegisterCustomType("main.Movie", Movie{})
	err = sexpr.Unmarshal(b, &strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unmarshalled output:")
	fmt.Printf("%v\n", strangelove)

}
