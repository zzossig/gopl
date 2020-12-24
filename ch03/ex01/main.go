/*
	If the function `f` returns a non-finite float64 value,
	the SVG file will contain invalid `<polygon> ` elements
	(although many SVG renderers handle this gracefully).
	Modify the program to skip invalid polygons.
*/

package main

import (
	"fmt"
	"math"
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
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*sin30*xyscale
	sy := height/2 + (x+y)*cos30*xyscale - z*zscale

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