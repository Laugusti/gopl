package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	s := strings.Split(os.Args[1], "")

	if IsPalindrome(sort.StringSlice(s)) {
		fmt.Printf("%s IS a palindrome\n", s)
	} else {
		fmt.Printf("%s IS NOT a palindrome\n", s)
	}
}
