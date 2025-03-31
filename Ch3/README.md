# Problem Set

## 3.1
If the function ```f``` returns a non-finite ```float64``` value, the SVG file will contain invalid ```<polygon>``` elements (although many SVG renders handle this gracefully). Modify the program to skip invalid polygons.

## 3.2
Experiment with visualizations of other functions from the ```math``` package. Can you produce an egg box, moguls, or a saddle?

## 3.3
Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valeys blue (#0000ff).

## 3.4
Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the ```Content-Type``` header like this:

```
w.Header().Set("Content-Type", "image/svg+xml")
```

## 3.5
Implement a full-color Mandlebrot set usig the function ```image.NewRGBA``` and the type ```color.RGBA``` or ```color.YCbCr```.

## 3.6
Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average. The simplest method is to divide each pixel into four "subpixels." Implement it.

## 3.7
Another simple fractal used Newton's method to find complex solutions to a function such as z<sup>4</sup> - 1 = 0. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root if applicable.

## 3.8
Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: ```complex64```, ```complex128```, ```big.Float```, and ```big.Rat```. (The latter two types are found in the math/big package. ```Float``` uses arbitraty but bounded-precision floating-point; ```Rat``` uses unbounded-precision rational numbers.) How do they compare in performance and memory usage? At what zoom levels do rendering artifacts become visual. 

## 3.9
Write a web server that renders fractals and writes the image data to the client. Allow the client to pecify the x, y, amd zoom values as paramters to the HTTP request. 

## 3.10
Write a non-recursive version of ```comma```, using ```bytes.Buffer``` instead of string concatenation.

## 3.11
Enchance ```comma``` so that it deals correctly with floating-point numbers and an optional sign.

## 3.12
Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters in a different order.

## 3.13
Write ```const``` declarations of KB, MB, up through YB as compactly as you can.
