package main

import "fmt"

func main() {
	fmt.Printf("The meaning of life: %d\n", theMeaningOfLife())
}

func theMeaningOfLife() (result int) {
	type quit struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case quit{}:
			result = 42
		default:
			panic(p)
		}
	}()
	panic(quit{})
}
