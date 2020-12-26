/*
	Modify `issues` to report the results in age categories,
	say less than a month old, less than a year old, and more than a year old.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Result is a result type for json decode
type Result struct {
	TotalCount `json:"total_count"`
	Items []Item
}

// Item represents an issue.
type Item struct {
	Number int
	Title string
	User User
	CreatedAt time.Time `json:"created_at"`
}

// User contains user info
type User struct {
	Login string
}

// TotalCount total issue number
type TotalCount int

// IssuesURL is a github api url
const IssuesURL = "https://api.github.com/search/issues"

func main() {
	args := os.Args[1:]

	result, err := searchIssues(args)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	// not efficient, but I think this is enough. I just need to know how to parse json, how to use structure ebeded structure.
	fmt.Println("less than a month old")
	for _, item := range result.Items {
		days := time.Since(item.CreatedAt).Hours() / 24
		if days < 30 {
			fmt.Printf("#%-5d %-9.9s %-12s %-.55s\n", item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
	}

	fmt.Println("less than a year old")
	for _, item := range result.Items {
		days := time.Since(item.CreatedAt).Hours() / 24
		if days < 365 {
			fmt.Printf("#%-5d %-9.9s %-12s %-.55s\n", item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
	}

	fmt.Println("more than a year old")
	for _, item := range result.Items {
		days := time.Since(item.CreatedAt).Hours() / 24
		if days > 365 {
			fmt.Printf("#%-5d %-9.9s %-12s %-.55s\n", item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
	}
}

func searchIssues(args []string) (Result, error) {
	q := url.QueryEscape(strings.Join(args, " "))
	r := Result{}

	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return r, err
	}
	
	return r, nil
}