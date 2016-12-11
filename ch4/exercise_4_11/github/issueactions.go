package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetIssue(owner string, repo string, issue int) (*Issue, error) {
	resp, err := http.Get(GetIssueURL(owner, repo, issue))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get issue: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CreateIssue(owner string, repo string, title string, body string) (*Issue, error) {
	url := GetIssueURL(owner, repo, 0)

	var issue Issue
	issue.Title = title
	issue.Body = body

	payload, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}

	return updateIssue(url, "POST", http.StatusCreated, bytes.NewReader(payload))
}

func EditIssue(owner string, repo string, issueNumber int, title string, body string) (*Issue, error) {
	url := GetIssueURL(owner, repo, issueNumber)

	issue, err := GetIssue(owner, repo, issueNumber)
	if err != nil {
		return nil, err
	}
	issue.Title = title
	issue.Body = body

	payload, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	return updateIssue(url, "PATCH", http.StatusOK, bytes.NewReader(payload))
}

func LockIssue(owner string, repo string, issue int) (*Issue, error) {
	url := GetIssueURL(owner, repo, issue) + "/lock"
	_, err := updateIssue(url, "PUT", http.StatusNoContent, nil)
	if err != nil {
		return nil, err
	} else {
		return &Issue{Number: issue}, nil
	}
}

func UnlockIssue(owner string, repo string, issue int) (*Issue, error) {
	url := GetIssueURL(owner, repo, issue) + "/lock"
	_, err := updateIssue(url, "DELETE", http.StatusNoContent, nil)
	if err != nil {
		return nil, err
	} else {
		return &Issue{Number: issue}, nil
	}
}

func OpenIssue(owner string, repo string, issue int) (*Issue, error) {
	url := GetIssueURL(owner, repo, issue)
	payload := strings.NewReader(`{"state":"open"}`)
	return updateIssue(url, "PATCH", http.StatusOK, payload)
}

func CloseIssue(owner string, repo string, issue int) (*Issue, error) {
	url := GetIssueURL(owner, repo, issue)
	payload := strings.NewReader(`{"state":"closed"}`)
	return updateIssue(url, "PATCH", http.StatusOK, payload)
}
