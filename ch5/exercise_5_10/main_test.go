package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	sorted := topoSort(prereqs)
	for i, course := range sorted {
		if isPreReq(course, sorted[:i]) {
			t.Fail()
		}
	}
}

func isPreReq(course string, courses []string) bool {
	for _, c := range courses {
		for _, prereq := range prereqs[c] {
			if course == prereq {
				return true
			}
		}
	}
	return false
}
