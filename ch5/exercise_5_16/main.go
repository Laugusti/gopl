package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(joinStrings("||", os.Args[1:]...))
}

func joinStrings(sep string, list ...string) (result string) {
	for i, s := range list {
		if i != 0 {
			result += sep
		}
		result += s
	}
	return
}
