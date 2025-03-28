package main

import "testing"

var test_args []complex128

func init() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 4, 4
	)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			test_args = append(test_args, z)
		}
	}
}

func BenchmarkMand64(b *testing.B) {
	for b.Loop() {
		mandelbrot64(test_args[0])
	}
}

func BenchmarkMand128(b *testing.B) {
	for b.Loop() {
		mandelbrot128(test_args[1])
	}
}

func BenchmarkMandFlt(b *testing.B) {
	for b.Loop() {
		mandelbrotFlt(test_args[2])
	}
}

func BenchmarkMandRat(b *testing.B) {
	for b.Loop() {
		mandelbrotRat(test_args[3])
	}
}
