/*
	Modify the `Lissajous` server to read parameter values from the URL.
	For example, you might arrange it so that a URL like `http://localhost:8000/?cycles=20`
	sets the number of cycles to 20 instead of the default 5.
	Use the `strconv.Atoi` function to convert the string parameter into an integer.
	You can see its documentation with `go doc strconv.Atoi`.
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, r.URL.RawQuery)
		}
		http.HandleFunc("/", handler)
		http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Println("prevent running handler twice")
		})
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	
	fmt.Println("usage: go run main.go web")
}

func lissajous(out io.Writer, rawquery string) {
	// parse raw query
	queries := make(map[string]string)
	rawQueries := strings.Split(rawquery, "&")

	if len(rawQueries) > 0 {
		for _, rawQuery := range rawQueries {
			query := strings.Split(rawQuery, "=")
			if len(query) == 2 {
				queries[query[0]] = query[1]
			}
		}
	}

	// set variable from parsed query
	var cycles int
	qCycles, ok := queries["cycles"]
	if ok {
		aCycles, err := strconv.Atoi(qCycles)
		if err != nil {
			cycles = 5
		} else {
			cycles = aCycles
		}
	} else {
		cycles = 5
	}

	var res float64
	qRes, ok := queries["res"]
	if ok {
		aRes, err := strconv.ParseFloat(qRes, 64)
		if err != nil {
			res = 0.001
		} else {
			res = aRes
		}
	} else {
		res = 0.001
	}

	var size int
	qSize, ok := queries["size"]
	if ok {
		aSize, err := strconv.Atoi(qSize)
		if err != nil {
			size = 100
		} else {
			size = aSize
		}
	} else {
		size = 100
	}

	var nframes int
	qNframes, ok := queries["nframes"]
	if ok {
		aNframes, err := strconv.Atoi(qNframes)
		if err != nil {
			nframes = 64
		} else {
			nframes = aNframes
		}
	} else {
		nframes = 64
	}

	var delay int
	qDelay, ok := queries["delay"]
	if ok {
		aDelay, err := strconv.Atoi(qDelay)
		if err != nil {
			delay = 8
		} else {
			delay = aDelay
		}
	} else {
		delay = 8
	}

	// check if all variables set correctly
	fmt.Printf("cycles: %d, res: %f, size: %d, nframes: %d, delay: %d\n", cycles, res, size, nframes, delay)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}