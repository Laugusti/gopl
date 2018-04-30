package display

import (
	"testing"
)

func TestMapOfStruct(t *testing.T) {
	m := map[struct{ a int }]bool{{1}: true, {2}: false}
	Display("m", m)
}

func TestMapOfArray(t *testing.T) {
	m := map[[2]int]bool{[2]int{1, 2}: true, [2]int{3, 4}: false}
	Display("m", m)
}
