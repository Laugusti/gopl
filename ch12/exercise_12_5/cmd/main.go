package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Laugusti/gopl/ch12/exercise_12_5"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func (m1 Movie) equals(m2 Movie) bool {
	return m1.Title == m2.Title && m1.Subtitle == m2.Subtitle && m1.Year == m2.Year && m1.Color == m2.Color && strMapEquals(m1.Actor, m2.Actor) && strListEquals(m1.Oscars, m2.Oscars) && strPtrEquals(m1.Sequel, m2.Sequel)
}

func strMapEquals(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}
	return true
}

func strListEquals(l1, l2 []string) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}

func strPtrEquals(s1, s2 *string) bool {
	return s1 != nil && s2 != nil && *s1 == *s2 || s1 == s2
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
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

	b, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("marshalled value:\n%s\n", string(b))

	var m Movie
	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("unmarshalled value equals: %t\n", strangelove.equals(m))
}
