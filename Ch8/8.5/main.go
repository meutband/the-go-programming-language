// This file refactors the Mandelbrot program with goroutines
package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"time"
)

// N represent the number of goroutines
const N = 128

func main() {

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	start := time.Now()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	size := height / N
	var done = make(chan struct{})
	for i := 0; i < N; i++ {
		go func(j int) {
			for py := size * j; py < size*(j+1); py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			done <- struct{}{}

		}(i)
	}

	for i := 0; i < N; i++ {
		<-done
	}

	// png.Encode(os.Stdout, img) // NOTE: ignoring errors
	fmt.Printf("time to run %v\n", time.Since(start))
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
