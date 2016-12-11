// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues
package github

import (
	"strconv"
	"time"
)

const urlPrefix = "https://api.github.com/repos/"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"` // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func GetIssueURL(owner string, repo string, issue int) string {
	if issue > 0 {
		return urlPrefix + owner + "/" + repo + "/issues/" + strconv.Itoa(issue)
	} else {
		return urlPrefix + owner + "/" + repo + "/issues"
	}
}
