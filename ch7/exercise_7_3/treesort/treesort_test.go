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

func TestString(t *testing.T) {
	root := makeTree()
	if root.String() != "[]" {
		t.Errorf("TestString failed")
	}

	root = makeTree(8, 2, 1, 9, 0, 5, 3, 7, 4, 6)
	if root.String() != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Errorf("TestString failed")
	}
}

func makeTree(vals ...int) *tree {
	var root *tree
	for _, val := range vals {
		root = add(root, val)
	}
	return root
}
