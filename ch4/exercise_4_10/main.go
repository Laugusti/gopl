package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/Laugusti/gopl/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	issues := issuesCreatedInRange(result, math.Inf(-1), 24*30)
	fmt.Printf("%d issues (<1 month):\n", len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println()

	issues = issuesCreatedInRange(result, 24*30, 24*365)
	fmt.Printf("%d issues (>1 month, <1 year):\n", len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println()

	issues = issuesCreatedInRange(result, 24*365, math.Inf(1))
	fmt.Printf("%d issues (>1 year):\n", len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func issuesCreatedInRange(searchResult *github.IssuesSearchResult, hoursStart float64, hoursEnd float64) []*github.Issue {
	var result []*github.Issue
	for _, item := range searchResult.Items {
		d := time.Since(item.CreatedAt).Hours()
		if d >= hoursStart && d < hoursEnd {
			result = append(result, item)
		}
	}
	return result
}
