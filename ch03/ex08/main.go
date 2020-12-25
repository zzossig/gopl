/*
	Rendering fractals at high zoom levels demands great arithmetic precision.
	Implement the same fractal using four different representations of numbers:
	`complex64`, `complex128`, `big.Float`, and `big.Rat`.
	(The latter two types are found in the `math/big` package.
	`Float` uses arbitrary but bounded-precision floating-point;
	`Rat` uses unbounded-precision rational numbers)
	How do they compare in performance and memory usage?
	At what zoom levels do rendering artifacts become visible?
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
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
			img.Set(px, py, mandelBigRat(z))
		}
	}
	
	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}

	png.Encode(f, img)
	f.Close()
}

// func newtonBigRat(z complex128) color.Color {
// 	const iterations = 37
// 	const contrast = 7
// 	for i := uint8(0); i < iterations; i++ {
// 		z -= (z - 1/(z*z*z)) / 4
// 		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
// 			return color.Gray{255 - contrast*i}
// 		}
// 	}
// 	return color.Black
// }

// func newtonBigFloat(z complex128) color.Color {
// 	const iterations = 37
// 	const contrast = 7

// 	re := new(big.Float).SetFloat64(real(z)) // real number
// 	im := new(big.Float).SetFloat64(imag(z)) // imagin number

// 	for i := uint8(0); i < iterations; i++ {
// 		re2, im2 := calcBigFloatMulti(*re, *im, *re, *im)
// 		re3, im3 := calcBigFloatMulti(*re2, *im2, *re, *im)

// 		var denoL *big.Float
// 		denoL = denoL.Mul(big.NewFloat(4), (denoL.Mul(re3, re3)))
// 		var denoR *big.Float
// 		denoR = denoR.Mul(big.NewFloat(4), (denoR.Mul(im3, im3)))
// 		var denoSum *big.Float
// 		denoSum = denoSum.Add(denoL, denoR) // denoSum is real number

// 		var im3Minus *big.Float
// 		var re3Minus *big.Float

// 		im3Minus = im3Minus.Mul(big.NewFloat(-1), im3)
// 		re3Minus = re3Minus.Mul(big.NewFloat(-1), re3)
// 		moleLReal, moleLImag := calcBigFloatMulti(*re, *im, *re3, *im3Minus)

// 		var moleReal *big.Float
// 		var moleImag *big.Float
// 		moleReal = moleReal.Add(moleLReal, re3Minus)
// 		moleImag = moleImag.Add(moleLImag, im3)

// 		z = z - (z - 1/(z*z*z)) / 4
// 		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
// 			return color.Gray{255 - contrast*i}
// 		}
// 	}
// 	return color.Black
// }

func mandelBigRat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	zReal := new(big.Rat).SetFloat64(real(z)) // real number
	zImag := new(big.Rat).SetFloat64(imag(z)) // imagin number

	vReal := new(big.Rat)
	vImag := new(big.Rat)

	for n := uint8(0); n < iterations; n++ {
		vReal2 := new(big.Rat)
		vImag2 := new(big.Rat)

		vReal2.Mul(vReal, vReal).Sub(vReal2, new(big.Rat).Mul(vImag, vImag)).Add(vReal2, zReal) // real
		vImag2.Mul(vReal, vImag).Mul(vImag2, big.NewRat(2, 1)).Add(vImag2, zImag) // imag
		vReal, vImag = vReal2, vImag2

		vsqrt := new(big.Rat)
		vsqrt.Mul(vReal, vReal).Add(vsqrt, new(big.Rat).Mul(vImag, vImag))

		if vsqrt.Cmp(big.NewRat(4, 1)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	zReal := new(big.Float).SetFloat64(real(z)) // real number
	zImag := new(big.Float).SetFloat64(imag(z)) // imagin number

	vReal := new(big.Float)
	vImag := new(big.Float)

	for n := uint8(0); n < iterations; n++ {
		vReal2 := new(big.Float)
		vImag2 := new(big.Float)

		vReal2.Mul(vReal, vReal).Sub(vReal2, new(big.Float).Mul(vImag, vImag)).Add(vReal2, zReal) // real
		vImag2.Mul(vReal, vImag).Mul(vImag2, big.NewFloat(2)).Add(vImag2, zImag) // imag
		vReal, vImag = vReal2, vImag2

		vsqrt := new(big.Float)
		vsqrt.Mul(vReal, vReal).Add(vsqrt, new(big.Float).Mul(vImag, vImag))

		if vsqrt.Cmp(big.NewFloat(4)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// What can I do except type casting?
func mandel64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	z64 := complex64(z)

	var v64 complex64
	for n := uint8(0); n < iterations; n++ {
		v64 = v64*v64 + z64
		if cmplx.Abs(complex128(v64)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// original version
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