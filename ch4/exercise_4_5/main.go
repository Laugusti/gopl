// eliminateAdjacentDups eliminates adjacent duplicates (in-place) in a []string slice.
package main

import "fmt"

func main() {
	data := []string{"one", "two", "two", "three"}
	fmt.Printf("%q\n", eliminateAdjacentDups(data)) // `["one" "two" "three"]`
	fmt.Printf("%q\n", data)                        // `["one" "two" "three" "three"]`
}

func eliminateAdjacentDups(strings []string) []string {
	i := 0
	var prevStr string
	for _, s := range strings {
		if i == 0 || s != prevStr {
			prevStr = s
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
