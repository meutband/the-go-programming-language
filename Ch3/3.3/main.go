// This file computes SVG rendering of a 3-D surface function

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(i+1, j)
			bx, by, berr := corner(i, j)
			cx, cy, cerr := corner(i, j+1)
			dx, dy, derr := corner(i+1, j+1)
			r, b := color(i, j)
			if aerr == nil && berr == nil && cerr == nil && derr == nil {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"rgb(%d,0,%d)\"/>\n", ax, ay, bx, by, cx, cy, dx, dy, r, b)
			}
		}
	}
	fmt.Println("</svg>")
}

func color(i, j int) (int, int) {
	_, _, z := xyz(i, j)
	var r, b float64
	if !math.IsNaN(z) || !math.IsInf(z, 0) {
		if z >= 0 {
			r = 255 * z
		} else {
			b = 255 * math.Abs(z)
		}
	}
	return int(r), int(b)
}

func corner(i, j int) (float64, float64, error) {

	x, y, z := xyz(i, j)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, fmt.Errorf("f(%g,%g) is invalid for corner(%d,%d)", x, y, i, j)
	}

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func xyz(i, j int) (float64, float64, float64) {
	// Find point (x, y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Computer service height z.
	z := f(x, y)
	return x, y, z
}
