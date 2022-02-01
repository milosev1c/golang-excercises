// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//!-main
// Packages not needed by version in book.
import (
	"time"
)

//!+main

var palette = []color.Color{color.Black, color.RGBA{R: 255, A: 255}, color.RGBA{B: 255, A: 255}, color.RGBA{G: 255, A: 255}, color.RGBA{R: 0xA8, G: 0xE4, B: 0xA0, A: 0xFF}}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	//!+main
	args := make(map[string]string)
	for _, arg := range os.Args[1:] {
		argsPieces := strings.Split(arg, "=")
		args[strings.Replace(argsPieces[0], "-", "", 5)] = argsPieces[1]
	}
	fmt.Println(args["type"])
	bgColor, _ := strconv.Atoi(args["bgcolor"])
	drawColor, _ := strconv.Atoi(args["drawcolor"])
	if val, ok := args["type"]; val == "web" && ok {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			//form := make(map[string]string)
			err := r.ParseForm()
			if err != nil {
				fmt.Println("Леее брат ошибка с запуском сервера")
			}
			cycles, nframes := 5, 256
			if val, ok := r.Form["cycles"]; ok {
				cycles, _ = strconv.Atoi(val[0])
			}
			if val, ok := r.Form["nframes"]; ok {
				nframes, _ = strconv.Atoi(val[0])
			}
			if val, ok := r.Form["bgcolor"]; ok {
				bgColor, _ = strconv.Atoi(val[0])
			}
			if val, ok := r.Form["drawcolor"]; ok {
				drawColor, _ = strconv.Atoi(val[0])
			}
			lissajous(w, uint8(bgColor), uint8(drawColor), cycles, nframes)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
}

func lissajous(out io.Writer, bgIndex uint8, drawIndex uint8, cycles int, nframes int) {
	if nframes == 0 {
		nframes = 256
	}
	if cycles == 0 {
		cycles = 5
	}
	const (
		res   = 0.001 // angular resolution
		size  = 400   // image canvas covers [-size..+size]
		delay = 4     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		img.SetColorIndex(0, 0, bgIndex)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				drawIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
}

//!-main
