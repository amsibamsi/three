// Package main contains an example program that renders a simple triangle and
// stores it in a file.
package main

import (
	"flag"
	"github.com/amsibamsi/three"
	"image/color"
	"os"
)

// main creates a new scene with a camera and 3 vectors, renders the scene,
// draws the result to an image and encodes the image as PNG into a file.
// Optionally specify the filename.
func main() {
	var filename = flag.String("file", "triangle.png", "Filename to store image")
	flag.Parse()
	cam := three.NewDefCam()
	v1 := three.NewVec4(-1, -1, -3)
	v2 := three.NewVec4(0, 1, -5)
	v3 := three.NewVec4(1, -1, -3)
	t := cam.PerspTransf(500, 500)
	w1 := t.Transf(v1)
	w1.Norm()
	w2 := t.Transf(v2)
	w2.Norm()
	w3 := t.Transf(v3)
	w3.Norm()
	x1 := three.Round(w1[0])
	y1 := three.Round(w1[1])
	x2 := three.Round(w2[0])
	y2 := three.Round(w2[1])
	x3 := three.Round(w3[0])
	y3 := three.Round(w3[1])
	img := three.NewImage(500, 500)
	col := color.RGBA{255, 255, 0, 255}
	img.DrawDot(x1, y1, col)
	img.DrawDot(x2, y2, col)
	img.DrawDot(x3, y3, col)
	img.DrawLine(x1, y1, x2, y2, col)
	img.DrawLine(x2, y2, x3, y3, col)
	img.DrawLine(x3, y3, x1, y1, col)
	file, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img.WritePng(file)
}
