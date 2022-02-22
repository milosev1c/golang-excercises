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
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 2048 * 4, 2048 * 4
)

var bigImage [width][height]color.Color
var smallImage [width / 2][height / 2]color.Color

func main() {
	getBigImage(-2, -2, 2, 2)
	getSampledImage()
	render(os.Stdout)
}

func render(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
	for py := 0; py < height/2; py++ {
		for px := 0; px < width/2; px++ {
			img.Set(px, py, smallImage[px][py])
		}
	}
	png.Encode(w, img)
}
func getBigImage(xMin, yMin, xMax, yMax float64) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			bigImage[px][py] = mandelbrot(z)
		}
	}
}

func getSampledImage() {
	for py := 0; py < height/2; py++ {
		for px := 0; px < width/2; px++ {
			r1, g1, b1, _ := color.Color.RGBA(bigImage[2*px][2*py])
			r2, g2, b2, _ := color.Color.RGBA(bigImage[2*px+1][2*py])
			r3, g3, b3, _ := color.Color.RGBA(bigImage[2*px][2*py+1])
			r4, g4, b4, _ := color.Color.RGBA(bigImage[2*px+1][2*py+1])
			r1 >>= 8
			r2 >>= 8
			r3 >>= 8
			r4 >>= 8
			g1 >>= 8
			g2 >>= 8
			g3 >>= 8
			g4 >>= 8
			b1 >>= 8
			b2 >>= 8
			b3 >>= 8
			b4 >>= 8
			r := math.Sqrt(float64((r1*r1 + r2*r2 + r3*r3 + r4*r4) / 4))
			g := math.Sqrt(float64((g1*g1 + g2*g2 + g3*g3 + g4*g4) / 4))
			b := math.Sqrt(float64((b1*b1 + b2*b2 + b3*b3 + b4*b4) / 4))
			smallImage[px][py] = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return map2RGB(255 - 16*n)
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
