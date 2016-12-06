// rotate rotates a slice of ints in place
package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5}
	rotate(data[:], 2)
	fmt.Println(data) // "[2 3 4 5 0 1]"
}

func rotate(arr []int, pos int) {
	tmp := make([]int, pos)
	copy(tmp, arr[:pos])
	copy(arr, arr[pos:])
	copy(arr[len(arr)-pos:], tmp)
}
