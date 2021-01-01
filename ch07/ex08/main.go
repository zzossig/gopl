/*
	Many GUIs provide a table widget with a stateful multi-tier sort:
	the primary sort key is the most recently clicked column head,
	the secondary sort key is the second-most recently clicked column head, and so on.
	Define an implementation of `sort.Interface` for use by such a table.
	Compare that approach with repeated sorting using `sort.Stable`.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type movieReview struct {
	title string
	rating int
	runningTime time.Duration
}

var sortOrder []string = []string{"title","rating","runningTime"}

func main() {
	reviews := []movieReview{
		{"Soul", 8, time.Hour + (time.Minute * 40)},
		{"Tenet", 7, (time.Hour * 2) + (time.Minute * 40)},
		{"Gambit", 7, (time.Hour * 1) + (time.Minute * 25)},
		{"Bridgerton", 8, (time.Hour * 1) + (time.Minute * 35)},
		{"Escape", 9, (time.Hour * 2) + (time.Minute * 15)},
		{"Mira", 6, (time.Hour * 2) + (time.Minute * 2)},
	}
	
	clickSimulation("runningTime", reviews)
	clickSimulation("rating", reviews)

	for _, r := range reviews {
		fmt.Printf("%v\n", r)
	}
}

func clickSimulation(byWhat string, data []movieReview) {
	if !isContain(sortOrder, byWhat) {
		// do nothing
	} else if strings.EqualFold(byWhat, "title") {
		sortOrder = moveToFront(byWhat, sortOrder)
		sort.Slice(data, func(i, j int) bool {
			return compareTitle(data[i].title, data[j].title)
		})
	} else if strings.EqualFold(byWhat, "rating") {
		sortOrder = moveToFront(byWhat, sortOrder)
		sort.Slice(data, func(i, j int) bool {
			if data[i].rating == data[j].rating {
				if sortOrder[1] == "title" {
					compareTitle(data[i].title, data[j].title)
				} else {
					compareRating(data[i].rating, data[j].rating)
				}
			}
			return compareRating(data[i].rating, data[j].rating)
		})
	} else if strings.EqualFold(byWhat, "runningTime") {
		sortOrder = moveToFront(byWhat, sortOrder)
		sort.Slice(data, func(i, j int) bool {
			return compareTime(data[i].runningTime, data[j].runningTime)
		})
	}
}

func compareTitle(t1 string, t2 string) bool {
	return strings.Compare(t1, t2) < 0
}

func compareRating(r1 int, r2 int) bool {
	return r1 < r2
}

func compareTime(t1 time.Duration, t2 time.Duration) bool {
	return t1 < t2
}

func isContain(ss []string, s string) bool {
	for _, str := range ss {
		if strings.EqualFold(str, s) {
			return true
		}
	}
	return false
}

func moveToFront(needle string, haystack []string) []string {
	if len(haystack) == 0 || haystack[0] == needle {
		return haystack
	}
	var prev string
	for i, elem := range haystack {
		switch {
		case i == 0:
			haystack[0] = needle
			prev = elem
		case elem == needle:
			haystack[i] = prev
			return haystack
		default:
			haystack[i] = prev
			prev = elem
		}
	}
	return append(haystack, prev)
}