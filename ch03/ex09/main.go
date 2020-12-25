/*
	Write a web server that renders fractals and writes tye image data to the client.
	Allow the client to specify the `x,y`, and `zoom` values as parameters to the HTTP request.
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			xmin, ymin, xmax, ymax float64
			width, height        float64
		)

		rawquery := r.URL.RawQuery
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

		x, ok := queries["x"]
		if ok {
			xx, err := strconv.ParseFloat(x, 64)
			if err == nil {
				width = xx
			} else {
				width = 1024
			}
		} else {
			width = 1024
		}

		y, ok := queries["y"]
		if ok {
			yy, err := strconv.ParseFloat(y, 64)
			if err == nil {
				height = yy
			} else {
				height = 1024
			}
		} else {
			height = 1024
		}

		zoom, ok := queries["zoom"]
		if ok {
			zz, err := strconv.ParseFloat(zoom, 64)
			if err == nil {
				xmin = -zz
				ymin = -zz
				xmax = zz
				ymax = zz
			} else {
				xmin = -1
				ymin = -1
				xmax = 1
				ymax = 1
			}
		} else {
			xmin = -1
			ymin = -1
			xmax = 1
			ymax = 1
		}

		img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
		for py := 0; py < int(height); py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < int(width); px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}

		png.Encode(w, img)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}