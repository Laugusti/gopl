package treesort

import (
	"testing"
)

func TestSort(t *testing.T) {
	data := []int{8, 2, 1, 9, 0, 5, 3, 7, 4, 6}
	Sort(data)
	for i, v := range data {
		if i != v {
			t.Fail()
		}
	}
}
