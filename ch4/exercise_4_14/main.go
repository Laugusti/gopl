package main

import (
	"html/template"
	"log"
	"net/http"
)

var issueList = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"getMilestoneURL": getMilestoneURL, "getMilestoneTitle": getMilestoneTitle}).
	Parse(`
<body>
<form>
  <span>Search:</span>
  <input type="text" name="query">
  <input type="submit">
</form>
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>User</th>
  <th>Title</th>
  <th>Milestone</th>
</tr>
{{range .Issues}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td><a href='{{. | getMilestoneURL}}'>{{. | getMilestoneTitle}}</a></td>
</tr>
{{end}}
</table>
`))

func getMilestoneURL(issue *Issue) string {
	if issue.Milestone == nil {
		return ""
	} else {
		return issue.Milestone.HTMLURL
	}
}

func getMilestoneTitle(issue *Issue) string {
	if issue.Milestone == nil {
		return ""
	} else {
		return issue.Milestone.Title
	}
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	result, err := SearchIssues(r.URL.Query().Get("query"))
	if err != nil {
		log.Print(err)
		if err = issueList.Execute(w, IssuesSearchResult{}); err != nil {
			log.Print(err)
		}
	} else {
		if err = issueList.Execute(w, result); err != nil {
			log.Print(err)
		}
	}
}
