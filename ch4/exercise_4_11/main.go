package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/Laugusti/gopl/ch4/exercise_4_11/github"
)

var mode = flag.String("mode", "read", "[read|create|update|lock|unlock|open|close]")
var owner = flag.String("owner", "golang", "owner of the repository")
var repo = flag.String("repo", "go", "name of repository")
var issueNumber = flag.Int("issue", 1, "Issue #")
var editor = flag.String("editor", "vim", "Text editor")

func main() {
	flag.Parse()
	checkArgs()
	switch *mode {
	case "create":
		createIssue()
	case "read":
		readIssue()
	case "update":
		updateIssue()
	case "lock":
		lockIssue()
	case "unlock":
		unlockIssue()
	case "open":
		openIssue()
	case "close":
		closeIssue()
	default:
		fmt.Println("not implemented")
	}
}

func getBody(current string) (string, error) {
	fmt.Print("Opening editor for body (press enter to continue)")
	in := bufio.NewReader(os.Stdin)
	_, _ = in.ReadByte()
	// create tmp file
	sep := string(os.PathSeparator)
	fname := os.TempDir() + sep + "842984020"
	tmpFile, err := os.Create(fname)
	if err != nil {
		return "", err
	}

	// populate tmp file with current
	w := bufio.NewWriter(tmpFile)
	_, err = w.WriteString(current)
	if err != nil {
		return "", err
	}
	err = w.Flush()
	if err != nil {
		return "", err
	}

	// open tmp file
	cmd := exec.Command(*editor, fname)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// read contents of tmp file and return
	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		return "", err
	}
	r := bufio.NewReader(tmpFile)
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func createIssue() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Title: ")
	title, err := in.ReadString('\n')
	checkErr(err)
	body, err := getBody("")
	checkErr(err)

	issue, err := github.CreateIssue(*owner, *repo, title[:len(title)-1], body)
	checkErr(err)
	fmt.Printf("Issue #%d created\n", issue.Number)
}

func updateIssue() {
	issue, err := github.GetIssue(*owner, *repo, *issueNumber)
	checkErr(err)

	fmt.Printf("Current Title: %s\n", issue.Title)
	fmt.Print("New Title (empty for unchanged): ")
	in := bufio.NewReader(os.Stdin)
	newTitle, err := in.ReadString('\n')
	checkErr(err)
	if len(newTitle) > 1 {
		issue.Title = newTitle[:len(newTitle)-1]
	}

	issue.Body, err = getBody(issue.Body)
	checkErr(err)

	issue, err = github.EditIssue(*owner, *repo, *issueNumber, issue.Title, issue.Body)
	checkErr(err)
	fmt.Printf("Issue #%d updated\n", issue.Number)
}

func lockIssue() {
	issue, err := github.LockIssue(*owner, *repo, *issueNumber)
	checkErr(err)
	fmt.Printf("Issue #%d locked\n", issue.Number)
}

func unlockIssue() {
	issue, err := github.UnlockIssue(*owner, *repo, *issueNumber)
	checkErr(err)
	fmt.Printf("Issue #%d unlocked\n", issue.Number)
}

func readIssue() {
	issue, err := github.GetIssue(*owner, *repo, *issueNumber)
	checkErr(err)
	fmt.Printf("Issue #%d\n", issue.Number)
	fmt.Printf("Title: %s\n", issue.Title)
	fmt.Println("Body:")
	fmt.Printf(issue.Body)
}

func closeIssue() {
	issue, err := github.CloseIssue(*owner, *repo, *issueNumber)
	checkErr(err)
	fmt.Printf("Issue #%d closed\n", issue.Number)
}

func openIssue() {
	issue, err := github.OpenIssue(*owner, *repo, *issueNumber)
	checkErr(err)
	fmt.Printf("Issue #%d opened\n", issue.Number)
}

func checkArgs() {
	if *mode != "read" && *mode != "create" && *mode != "update" && *mode != "lock" && *mode != "unlock" && *mode != "close" && *mode != "open" {
		flag.Usage()
		os.Exit(1)
	}
	if *owner == "" || *repo == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
