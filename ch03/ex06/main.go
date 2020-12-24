/*
	Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average.
	The simplestmethod is to divide each pixel into four "subpixels". Implement it.
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	offsetX := 0
	offsetY := 0

	for {
		// check if arrived edge
		if offsetX + 2 >= width && offsetY + 2 >= height {
			break
		}

		// extract colors
		colors := []color.Color{}
		for py := 0 + offsetY; py < 0 + offsetY + 2; py++ {
			for px := 0 + offsetX; px < 0 + offsetX + 2; px++ {
				colors = append(colors, img.At(px, py))
			}
		}

		// apply avg color
		for py := 0 + offsetY; py < 0 + offsetY + 2; py++ {
			for px := 0 + offsetX; px < 0 + offsetX + 2; px++ {
				img.Set(px, py, avgRGBA(colors))
			}
		}

		// move to next pixel
		if offsetX + 2 >= width {
			offsetY += 2
			offsetX = 0
		} else {
			offsetX += 2
		}
	}
	
	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}

	png.Encode(f, img)
	f.Close()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
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

// calculate average color
func avgRGBA(colors []color.Color) color.Color {
	var rsum uint16
	var gsum uint16
	var bsum uint16
	var asum uint16

	for _, c := range colors {
		r, g, b, a := c.RGBA()
		rsum += uint16(r)
		gsum += uint16(g)
		bsum += uint16(b)
		asum += uint16(a)
	}

	return color.RGBA64{(rsum / 4), (gsum / 4), (bsum / 4), (asum / 4)}
}