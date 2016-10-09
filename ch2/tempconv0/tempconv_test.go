package tempconv

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100"
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
}

func Test2(t *testing.T) {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f)          // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!
}

func Test3(t *testing.T) {
	c := FToC(212.0)        // "100°"
	fmt.Println(c.String()) // "100°"; no need to call String explicitly
	fmt.Printf("%v\n", c)   // "100°"
	fmt.Printf("%s\n", c)   // "100°"
	fmt.Println(c)          // "100°"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100" does not call String
}
