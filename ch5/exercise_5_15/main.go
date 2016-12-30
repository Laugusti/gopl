package main

import "fmt"
import "math"

func main() {
	fmt.Printf("min(): %d\n", min())
	fmt.Printf("min(1): %d\n", min(1))
	fmt.Printf("min(1, 2): %d\n", min(1, 2))
	fmt.Printf("max(): %d\n", max())
	fmt.Printf("max(1): %d\n", max(1))
	fmt.Printf("max(1, 2): %d\n", max(1, 2))

	fmt.Printf("min2(1, 2): %d\n", min2(1, 2))
	fmt.Printf("max2(1, 2): %d\n", max2(1, 2))
}

func max(vals ...int) int {
	max := math.MinInt32
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func max2(val int, vals ...int) int {
	return max(append(vals, val)...)
}

func min(vals ...int) int {
	min := math.MaxInt32
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func min2(val int, vals ...int) int {
	return min(append(vals, val)...)
}
