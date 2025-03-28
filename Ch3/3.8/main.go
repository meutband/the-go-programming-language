// This file emits a PNG mage of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"math/rand"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax            = -2, -2, 2, 2
		width, height                     = 1024, 1024
		cmplx64, cmplx128, bigFlt, bigRat = 0, 1, 2, 3
	)

	val := rand.Intn(4)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			switch val {
			case cmplx64:
				img.Set(px, py, mandelbrot64(z))
			case cmplx128:
				img.Set(px, py, mandelbrot128(z))
			case bigFlt:
				img.Set(px, py, mandelbrotFlt(z))
			case bigRat:
				img.Set(px, py, mandelbrotRat(z))
			}
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
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

func mandelbrotFlt(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	zReal := (&big.Float{}).SetFloat64(real(z))
	zImag := (&big.Float{}).SetFloat64(imag(z))

	real, imag := &big.Float{}, &big.Float{}
	for n := uint8(0); n < iterations; n++ {
		real2, imag2 := &big.Float{}, &big.Float{}
		real2.Mul(real, real).Sub(real2, (&big.Float{}).Mul(imag, imag)).Add(real2, zReal)
		imag2.Mul(real, imag).Mul(imag2, big.NewFloat(2)).Add(imag2, zImag)
		real, imag = real2, imag2

		sum := &big.Float{}
		sum.Mul(real, real).Add(sum, (&big.Float{}).Mul(imag, imag))
		if sum.Cmp(big.NewFloat(4)) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotRat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	zReal := (&big.Rat{}).SetFloat64(real(z))
	zImag := (&big.Rat{}).SetFloat64(imag(z))

	real, imag := &big.Rat{}, &big.Rat{}
	for n := uint8(0); n < iterations; n++ {
		real2, imag2 := &big.Rat{}, &big.Rat{}
		real2.Mul(real, real).Sub(real2, (&big.Rat{}).Mul(imag, imag)).Add(real2, zReal)
		imag2.Mul(real, imag).Mul(imag2, big.NewRat(2, 1)).Add(imag2, zImag)
		real, imag = real2, imag2

		sum := &big.Rat{}
		sum.Mul(real, real).Add(sum, (&big.Rat{}).Mul(imag, imag))
		if sum.Cmp(big.NewRat(4, 1)) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
