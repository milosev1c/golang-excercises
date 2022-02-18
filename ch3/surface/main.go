// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
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

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var polygons [cells][cells]Polygon

func main() {
	f, err := os.Create("out/ch3/surface.svg")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	zMin, zMax := GeneratePoints(cells)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if polygons[i][j].invalid() {
				continue
			}
			red := int(math.Floor(255 * ((polygons[i][j].zavg - zMin) / (zMax - zMin))))
			green := 0
			blue := 255 - red
			_, err = f.WriteString(fmt.Sprintf("<polygon fill='rgb(%d,%d,%d)' stroke='rgb(%d,%d,%d)' points='%g,%g %g,%g %g,%g %g,%g' />\n",
				red, green, blue,
				red, green, blue,
				polygons[i][j].ax, polygons[i][j].ay,
				polygons[i][j].bx, polygons[i][j].by,
				polygons[i][j].cx, polygons[i][j].cy,
				polygons[i][j].dx, polygons[i][j].dy))
		}
	}
	_, err = f.WriteString(fmt.Sprint("</svg>"))
	defer f.Close()
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
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

// пишешь свои типы……… люби и свои методы писать………
func (p *Polygon) invalid() bool {
	return math.IsNaN(p.ax) || math.IsNaN(p.ay) ||
		math.IsNaN(p.bx) || math.IsNaN(p.by) || math.IsNaN(p.cx) ||
		math.IsNaN(p.cy) || math.IsNaN(p.dx) || math.IsNaN(p.dy)
}
