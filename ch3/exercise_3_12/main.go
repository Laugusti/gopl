// anagram reports whether two strings are anagrams of each other.
package main

import "fmt"
import "os"
import "strings"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: anagram STRING1 STRING2")
		os.Exit(1)
	}

	if anagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s and %s are anagrams\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s and %s are not anagrams\n", os.Args[1], os.Args[2])
	}
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, v := range s1 {
		if index := strings.Index(s2, string(v)); index >= 0 {
			s2 = strings.Replace(s2, string(v), "", 1)
		} else {
			return false
		}
	}
	return true
}
