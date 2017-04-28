// Construct a tool that reports the set of all packages in the workspace
// that transitively depend on the packages specified by the arguments.
// Hint: you will need to run go list twice, once for the initial packages
// and once for all packages. You may want to parse to JSON output using
// the encoding/json package (4.5.)
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// check Args
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./godeps packages...")
		os.Exit(1)
	}

	// get all workspace packags and their dependencies
	packages, err := getPackages(".")
	if err != nil {
		log.Fatal(err)
	}
	pdSlice, err := getAllPackageDependencies(packages)
	if err != nil {
		log.Fatal(err)
	}

	for _, pd := range pdSlice {
		// if all packages in the args are in the dependencies,
		// then print the package name
		if isSuperSet(pd.Dependencies, os.Args[1:]) {
			fmt.Println(pd.ImportPath)
		}
	}
}

// isSuperSet returns true if s1 is a superset of s2.
func isSuperSet(s1, s2 []string) bool {
	// store s1 in map
	m := make(map[string]bool)
	for _, s := range s1 {
		m[s] = true
	}

	// return false if value in s2 is not in s1
	for _, s := range s2 {
		if !m[s] {
			return false
		}
	}
	return true
}
