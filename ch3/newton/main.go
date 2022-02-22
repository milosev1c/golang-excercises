package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/cmplx"
	"os"
)

const (
	width                  = 2048
	height                 = 2048
	xMin, yMin, xMax, yMax = -2, -2, 2, 2
)

var Image [width][height]color.Color

func newton(z complex128) color.Color {
	const (
		iterations = 255
		contrast   = 128
		tr         = 0.00001
	)

	for n := uint8(0); n < iterations; n++ {
		z = z - (z-1/(z*z*z))/4
		if cmplx.Abs(z*z*z*z-1) < tr {
			return map2RGB(255 - contrast*n)
		}
	}

	return color.Black
}

func map2RGB(n uint8) color.Color {

	h := 360 * float64(n) / 256
	x := uint8(255 * (1 - math.Abs(math.Mod(h/60, 2)-1)))

	switch {
	case h < 60:
		return color.RGBA{255, x, 0, 255}
	case h < 120:
		return color.RGBA{x, 255, 0, 255}
	case h < 180:
		return color.RGBA{0, 255, x, 255}
	case h < 240:
		return color.RGBA{0, x, 255, 255}
	case h < 300:
		return color.RGBA{x, 0, 255, 255}
	default:
		return color.RGBA{255, 0, x, 255}
	}
}

func createImage() {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			Image[px][py] = newton(z)
		}
	}
}

func render(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py, Image[px][py])
		}
	}
	png.Encode(w, img)
}

func main() {
	createImage()
	render(os.Stdout)
}
