package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Laugusti/gopl/ch12/methods"
)

func main() {
	methods.Print(time.Duration(0))
	fmt.Println()
	methods.Print(&strings.Replacer{})
}
