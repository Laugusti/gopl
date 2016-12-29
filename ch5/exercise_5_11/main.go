package main

import "fmt"
import "log"
import "sort"

// prereqs maps ocmputer science courses totheir prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"linear algebra":        {"calculus"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	sorted, err := topoSort(prereqs)
	if err != nil {
		log.Fatalf("error sorting courses: %s", err)
	}
	for i, course := range sorted {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)

	cycleChk := make(map[string]bool)
	var cycleFound bool

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if cycleChk[item] {
				cycleFound = true
			}
			cycleChk[item] = true
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
			delete(cycleChk, item)
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	if cycleFound {
		return nil, fmt.Errorf("found cycle")
	}
	return order, nil
}
