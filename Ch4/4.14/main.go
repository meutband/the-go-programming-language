package main

import (
	"fmt"
	"gobook/Ch4/4.14/github"
	"html/template"
	"net/http"
	"os"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func handler(issue *github.IssuesSearchResults) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := issueList.Execute(w, issue)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {

	issues, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler(issues))
	err = http.ListenAndServe(":8080", nil)
	panic(err)
}
