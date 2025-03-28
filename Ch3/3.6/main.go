// This file emits a PNG mage of the Mandelbrot fractal.
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
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	var colors [width * 2][height * 2]color.Color

	for py := 0; py < height*2; py++ {
		y := float64(py)/height*2*(ymax-ymin) + ymin
		for px := 0; px < width*2; px++ {
			x := float64(px)/width*2*(xmax-xmin) + xmin
			z := complex(x, y)

			colors[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {

			si, sj := 2*i, 2*j
			r1, g1, b1, a1 := colors[si][sj].RGBA()
			r2, g2, b2, a2 := colors[si+1][sj].RGBA()
			r3, g3, b3, a3 := colors[si+1][sj+1].RGBA()
			r4, g4, b4, a4 := colors[si][sj+1].RGBA()

			avgColor := color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 1028),
				uint8((g1 + g2 + g3 + g4) / 1028),
				uint8((b1 + b2 + b3 + b4) / 1028),
				uint8((a1 + a2 + a3 + a4) / 1028),
			}

			img.Set(i, j, avgColor)
		}
	}
	png.Encode(os.Stdout, img)
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
