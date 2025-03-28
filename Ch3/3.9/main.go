// This file emits a PNG mage of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

const (
	width, height = 1024, 1024
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	x, y, z := parseQuery(r.URL.Query())

	s := math.Abs(z)
	v := float64(2) / s
	if math.IsNaN(v) {
		v = 2
	}

	dx, dy := x/(s*width), y/(s*height)
	if math.IsNaN(dx) {
		dx = 0
	}
	if math.IsNaN(dy) {
		dy = 0
	}

	var xmin, ymin, xmax, ymax = -v, -v, +v, +v

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x+dx, y+dy)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func parseQuery(val url.Values) (float64, float64, float64) {
	x, err := strconv.ParseFloat(val.Get("x"), 64)
	if err != nil {
		x = 0
	}

	y, err := strconv.ParseFloat(val.Get("y"), 64)
	if err != nil {
		y = 0
	}

	z, err := strconv.ParseFloat(val.Get("zoom"), 64)
	if err != nil {
		z = 1
	}

	return x, y, z
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
