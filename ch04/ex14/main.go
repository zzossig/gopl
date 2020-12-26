// Create a web server that queries GitHub once and then allows navigation of the list of bug reports, milestones, and users.

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const targetURL = "https://api.github.com/repos/zzossig/hugo-theme-zzo/issues" // post
const token = "Your Github API Token"

// MyJSON wrap Issue
type MyJSON struct {
	Type string
	Issues []Issue
}

// Issue contains some fields of the returned result from the api
type Issue struct {
	URL string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL string `json:"labels_url"`
	ID int `json:"id"`
	NodeID string `json:"node_id"`
	Title string `json:"title"`
	User User `json:"user"`
}

// User contains some fields
type User struct {
	Login string `json:"login"`
	ID int `json:"id"`
	AvatarURL string `json:"avatar_url"`
}

func main() {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	issues, err := read();
	check(err)

	result := MyJSON{Issues: issues}

	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Test</title>
	</head>
	<body>
		<a href="/">Home</a>&nbsp;&nbsp;<a href="/title">Title</a>&nbsp;&nbsp;<a href="/url">URL</a>
		{{$type := .Type}}
		{{range .Issues}}
			{{if eq $type "title" }}
				<div>{{ .Title }}</div>
			{{else if eq $type "url"}}
				<div>{{ .URL }}</div>
			{{else if eq $type "home"}}
				<div>{{ .ID }}</div>
			{{end}}
		{{end}}
	</body>
</html>`

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result.Type = "home"
		err = t.Execute(w, result)
		check(err)
	})
	http.HandleFunc("/title", func(w http.ResponseWriter, r *http.Request) {
		result.Type = "title"
		err = t.Execute(w, result)
		check(err)
	})
	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		result.Type = "url"
		err = t.Execute(w, result)
		check(err)
	})
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("prevent running handler twice")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func read() ([]Issue, error) {
	client := &http.Client{}
	
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "bearer " + token)
	
	resp, err := client.Do(req);

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var issues []Issue

	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return issues, nil
}