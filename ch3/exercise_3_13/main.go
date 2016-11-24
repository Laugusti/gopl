package main

import "fmt"

const (
	_, iota = iota, iota * 1000
	KiB
	GiB
)

func main() {
	fmt.Println("KiB: ", KiB)
	fmt.Println("GiB: ", GiB)
}
