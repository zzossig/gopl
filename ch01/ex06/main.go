/*
	Modify the `Lissajous` program to produce images in multiple colors by adding more values to `palette`
	and then displaying them by changing the third argument of `SetColorIndex` in some interesting way.
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
	"time"
)

var palette = []color.Color{
	color.White, 
	color.Black, 
	color.RGBA{0x00, 0xff, 0x00, 0xff}, 
	color.RGBA{0xff, 0x00, 0x00, 0xff}, 
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x11, 0x22, 0x33, 0xff},
	color.RGBA{0x44, 0x55, 0x66, 0xff},
	color.RGBA{0x77, 0x88, 0x99, 0xff},
	color.RGBA{0xaa, 0xbb, 0xcc, 0xff},
	color.RGBA{0xdd, 0xee, 0xff, 0xff},
	color.RGBA{0xff, 0xee, 0xdd, 0xff},
	color.RGBA{0xcc, 0xbb, 0xaa, 0xff},
	color.RGBA{0x99, 0x88, 0x77, 0xff},
	color.RGBA{0x66, 0x55, 0x44, 0xff},
	color.RGBA{0x33, 0x22, 0x11, 0xff},
	color.RGBA{0x12, 0x34, 0x56, 0xff},
	color.RGBA{0x78, 0x9a, 0xbc, 0xff},
	color.RGBA{0xde, 0xef, 0xf0, 0xff},
	color.RGBA{0x0f, 0xfe, 0xdc, 0xff},
	color.RGBA{0xba, 0x98, 0x76, 0xff},
	color.RGBA{0x54, 0x32, 0x10, 0xff},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	
	fmt.Println("usage: go run main.go web")
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(21)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}