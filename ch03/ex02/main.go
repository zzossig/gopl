/*
	Experiment with visualizations of other functions from the math package.
	Can you produce an `egg box`, `moguls`, or a `saddle`?
*/

// I added a server to check the visualizations that is the question of ex04

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
	xyrange       = 50.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var (
	sin = math.Sin(angle)
	cos = math.Cos(angle)
	acos = math.Acos(angle)
	acosh = math.Acosh(angle)
	asin = math.Asin(angle)
	asinh = math.Asinh(angle)
	atan = math.Atan(angle)
	atanh = math.Atanh(angle)
	cosh = math.Cosh(angle)
	sinh = math.Sinh(angle)
	tan = math.Tan(angle)
	tanh = math.Tanh(angle)
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var b strings.Builder

			w.Header().Set("Content-Type", "image/svg+xml")
			
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

	sx := width/2 + (x+y)*cos*xyscale
	sy := height/2 + (x-y)*sin*xyscale - z*zscale

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

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	// return (math.Sin(x*r) + math.Sin(y*r)) / r
	
	// egg box
	return (math.Sin(x/1.5) + math.Sin(y/1.5)) / r
}