/*
	Following the approach of the `Lissajous` example in Section 1.7,
	construct a web server that computes surfaces and writes SVG data to the client.
	The server must set the `Content-Type` header like this:
	```golang
    w.Header().Set("Content-Type", "image/svg+xml")
	```
*/

package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var b strings.Builder
			
			fmt.Fprintf(&b, "<svg xmlns='http://www.w3.org/2000/svg' "+
				"style='stroke: grey; fill: white; stroke-width: 0.7' "+
				"width='%d' height='%d'>", width, height)
			for i := 0; i < cells; i++ {
				for j := 0; j < cells; j++ {
					ax, ay := corner(i+1, j)
					bx, by := corner(i, j)
					cx, cy := corner(i, j+1)
					dx, dy := corner(i+1, j+1)
					fmt.Fprintf(&b, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				}
			}
			
			fmt.Fprintf(&b, "</svg>")
			w.Header().Set("Content-Type", "image/svg+xml")

			fmt.Fprintf(w, b.String())
		})
		http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Println("prevent running handler twice")
		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	}

	fmt.Println("usage: go run main.go web")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(sx, 1) {
		sx = +0
	}
	if math.IsInf(sx, -1) {
		sx = -0
	}
	if math.IsNaN(sx) {
		sx = 0
	}

	if math.IsInf(sy, 1) {
		sy = +0
	}
	if math.IsInf(sy, -1) {
		sy = -0
	}
	if math.IsNaN(sy) {
		sy = 0
	}

	return sx, sy
}

// the x, y params never going to be zero at the same time so do not need to check
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}