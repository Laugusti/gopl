package tempconv

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Printf("Absolute Zero: %v\n", CToK(AbsoluteZeroC))
	fmt.Printf("Absolute Zero: %v\n", AbsoluteZeroC)
	fmt.Printf("Absolute Zero: %v\n", CToF(AbsoluteZeroC))

	fmt.Printf("Freezing: %v\n", CToK(FreezingC))
	fmt.Printf("Freezing: %v\n", FreezingC)
	fmt.Printf("Freezing: %v\n", CToF(FreezingC))

	fmt.Printf("Boiling: %v\n", CToK(BoilingC))
	fmt.Printf("Boiling: %v\n", BoilingC)
	fmt.Printf("Boiling: %v\n", CToF(BoilingC))
}
