package github

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func readGithubCreds() (string, string, error) {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("GitHub username: ")
	username, err := in.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	fmt.Print("GitHub password: ")
	passBytes, err := terminal.ReadPassword(0)
	fmt.Println()
	if err != nil {
		return "", "", err
	}
	return username[:len(username)-1], string(passBytes), nil
}

func updateIssue(url string, method string, expectedStatus int, payload io.Reader) (*Issue, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	username, password, err := readGithubCreds()
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != expectedStatus {
		var message struct{ Message string }
		json.NewDecoder(resp.Body).Decode(&message)

		resp.Body.Close()
		return nil, fmt.Errorf("failed to create/update issue: %s\n\tStatus: %s", message.Message, resp.Status)
	}

	if expectedStatus == http.StatusNoContent {
		resp.Body.Close()
		return nil, nil
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
