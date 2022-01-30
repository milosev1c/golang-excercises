// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"time"
)

//!+main

var palette []color.Color

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	var bg_color, paint_color color.Color
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "black":
			bg_color = color.Black
		case "red":
			bg_color = color.RGBA{R: 255, A: 255}
		case "green":
			bg_color = color.RGBA{G: 255, A: 255}
		case "blue":
			bg_color = color.RGBA{B: 255, A: 255}
		default:
			bg_color = color.Black
		}
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "black":
				paint_color = color.Black
			case "red":
				paint_color = color.RGBA{R: 255, A: 255}
			case "green":
				paint_color = color.RGBA{G: 255, A: 255}
			case "blue":
				paint_color = color.RGBA{B: 255, A: 255}
			default:
				paint_color = color.Black
			}
		} else {
			paint_color = color.RGBA{G: 255, A: 255}
		}
	} else {
		bg_color = color.Black
		paint_color = color.RGBA{G: 255, A: 255}
	}

	//!+main
	lissajous(os.Stdout, bg_color, paint_color)
}

func lissajous(out io.Writer, bg_color color.Color, paint_color color.Color) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 400   // image canvas covers [-size..+size]
		nframes = 256   // number of animation frames
		delay   = 4     // delay between frames in 10ms units
	)
	//colors := os.Args[1:2]
	//for _, color := range colors {
	//
	//}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	palette = []color.Color{bg_color, paint_color}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
