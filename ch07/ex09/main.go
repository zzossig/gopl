/*
	Use the `html/template` package (§4.6) to replace `printTracks` with a function that displays the tracks as an HTML table.
	Use the solution to the previous exercise to arrange that each click on a column head makes an HTTP request to sort the table.

	func printTracks(tracks []*Track) {
		const format = "%v\t%v\t%v\t%v\t%v\t\n"
		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
		fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
		for _, t := range tracks {
			fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
		}
		tw.Flush() // calculate column widths and print table
	}
*/

package main

import (
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
	"time"
)

type MovieReview struct {
	Title string
	Rating int
	RunningTime time.Duration
}

type Data struct {
	Items []MovieReview
}

var reviews = []MovieReview{
	{"Soul", 8, time.Hour + (time.Minute * 40)},
	{"Tenet", 7, (time.Hour * 2) + (time.Minute * 40)},
	{"Gambit", 7, (time.Hour * 1) + (time.Minute * 25)},
	{"Bridgerton", 8, (time.Hour * 1) + (time.Minute * 35)},
	{"Escape", 9, (time.Hour * 2) + (time.Minute * 15)},
	{"Mira", 6, (time.Hour * 2) + (time.Minute * 2)},
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		s := strings.TrimSpace(r.FormValue("sort"))
		if strings.EqualFold(s, "title") {
			clickSimulation("title", reviews)
		} else if strings.EqualFold(s, "rating") {
			clickSimulation("rating", reviews)
		} else if strings.EqualFold(s, "time") {
			clickSimulation("runningTime", reviews)
		}
		printReviews(rw, Data{reviews})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func printReviews(wr io.Writer, data Data) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Movie Review</title>
	</head>
	<body>
		<table>
			<tr>
				<th><a href="/?sort=title">Title</a></th>
				<th><a href="/?sort=rating">Rating</a></th>
				<th><a href="/?sort=time">Time</a></th>
			</tr>
			{{range .Items}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Rating}}</td>
					<td>{{.RunningTime}}</td>
				</tr>
			{{end}}
		</table>
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	err = t.Execute(wr, data)
	check(err)
}

var sortOrder []string = []string{"title","rating","runningTime"}

func clickSimulation(byWhat string, data []MovieReview) {
	if !isContain(sortOrder, byWhat) {
		// do nothing
	} else if strings.EqualFold(byWhat, "title") {
		sortOrder = moveToFront(byWhat, sortOrder)
		sort.Slice(data, func(i, j int) bool {
			return compareTitle(data[i].Title, data[j].Title)
		})
	} else if strings.EqualFold(byWhat, "rating") {
		sortOrder = moveToFront(byWhat, sortOrder)
		sort.Slice(data, func(i, j int) bool {
			if data[i].Rating == data[j].Rating {
				if sortOrder[1] == "title" {
					compareTitle(data[i].Title, data[j].Title)
				} else {
					compareRating(data[i].Rating, data[j].Rating)
				}
			}
			return compareRating(data[i].Rating, data[j].Rating)
		})
	} else if strings.EqualFold(byWhat, "runningTime") {
		sortOrder = moveToFront(byWhat, sortOrder)
		sort.Slice(data, func(i, j int) bool {
			return compareTime(data[i].RunningTime, data[j].RunningTime)
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