package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
	"unicode/utf8"
)

// randomPalindrom returns a palindrome whose length and contents
// are derived form the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		var r rune
		// get a random letter
		for !unicode.IsLetter(r) {
			r = rune(rng.Intn(0x1000)) // random run up to '\u0999'
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// randomNonPalindrom returns a non-palindrome whose length and contents
// are derived form the pseudo-random number generator rng.
func randomNonPalindrome(rng *rand.Rand) string {
	var n int
	// generate a random length until >2
	for n < 2 {
		n = rng.Intn(25) // random length up to 24
	}
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		var r rune
		// get a random letter
		for !unicode.IsLetter(r) {
			r = rune(rng.Intn(0x1000)) // random run up to '\u0999'
		}
		runes[i] = r
		// get another random letter that is not equal the previous
		for unicode.ToLower(r) == unicode.ToLower(runes[i]) || !unicode.IsLetter(r) {
			r = rune(rng.Intn(0x1000)) // random run up to '\u0999'
		}
		runes[n-1-i] = r
	}
	return string(runes)

}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true %d", p, utf8.RuneCountInString(p))
		}
	}
}
