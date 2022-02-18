package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

type Polygon struct {
	ax, ay, bx, by, cx, cy, dx, dy, zavg float64
}

func (p *Polygon) invalid() bool {
	return math.IsNaN(p.ax) || math.IsNaN(p.ay) ||
		math.IsNaN(p.bx) || math.IsNaN(p.by) || math.IsNaN(p.cx) ||
		math.IsNaN(p.cy) || math.IsNaN(p.dx) || math.IsNaN(p.dy)
}

var out io.Writer = os.Stdout
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var polygons [cells][cells]Polygon

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy)
}

func GeneratePoints(cells int) (float64, float64) {
	var z1, z2, z3, z4, zMin, zMax float64
	var err bool
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if polygons[i][j].ax, polygons[i][j].ay, z1, err = corner(i+1, j); err {
				continue
			}
			if polygons[i][j].bx, polygons[i][j].by, z2, err = corner(i, j); err {
				continue
			}
			if polygons[i][j].cx, polygons[i][j].cy, z3, err = corner(i, j+1); err {
				continue
			}
			if polygons[i][j].dx, polygons[i][j].dy, z4, err = corner(i+1, j+1); err {
				continue
			}
			zAvg := (z1 + z2 + z3 + z4) / 4.0
			polygons[i][j].zavg = zAvg
			if !math.IsNaN(zAvg) {
				zMin = math.Min(zMin, zAvg)
				zMax = math.Max(zMax, zAvg)
			}
		}
	}
	return zMin, zMax
}

func GenerateSVG(writer io.Writer) {
	fmt.Fprintf(writer, "<svg "+
		"version='1.1' "+
		"baseProfile='full' "+
		"xmlns='http://www.w3.org/2000/svg' "+
		"xmlns:xlink='http://www.w3.org/1999/xlink' "+
		"xmlns:ev='http://www.w3.org/2001/xml-events' "+
		"stroke='grey' "+
		"stroke-width='0.7px' "+
		"fill='white' "+
		"width='%dpx' "+
		"height='%dpx'>\n",
		width, height)
	zMin, zMax := GeneratePoints(cells)

	// Add SVG body skipping polygons with invalid corners
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if polygons[i][j].invalid() {
				continue
			}
			red := int(math.Floor(255 * ((polygons[i][j].zavg - zMin) / (zMax - zMin))))
			green := 0
			blue := 255 - red
			fmt.Fprintf(writer, "<polygon fill='rgb(%d,%d,%d)' stroke='rgb(%d,%d,%d)' points='%g,%g %g,%g %g,%g %g,%g' />\n",
				red, green, blue,
				red, green, blue,
				polygons[i][j].ax, polygons[i][j].ay,
				polygons[i][j].bx, polygons[i][j].by,
				polygons[i][j].cx, polygons[i][j].cy,
				polygons[i][j].dx, polygons[i][j].dy)
		}
	}

	fmt.Fprintf(writer, "</svg>\n")
}

func Server() {
	handler := func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		GenerateSVG(writer)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func main() {

	// If called with "web" as a parameter, launch web server
	// otherwise output an SVG
	Server()
}
