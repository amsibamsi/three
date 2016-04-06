// Package main renders a simple triangle and stores it in a file.
package main

import (
	"flag"
	"github.com/amsibamsi/3d/graphics"
	"image/color"
	"os"
)

// main creates a new scene with a camera and 3 vectors, renders the scene,
// draws the result to an image and encodes the image as PNG into a file.
// Optionally specify the filename.
func main() {
	var filename = flag.String("file", "triangle.png", "Filename to store image")
	flag.Parse()
	scr := graphics.Screen{500, 500}
	cam := graphics.NewDefCam()
	v1 := graphics.NewVec4(-1, -1, -3)
	v2 := graphics.NewVec4(0, 1, -5)
	v3 := graphics.NewVec4(1, -1, -3)
	t := cam.PerspTransf(&scr)
	w1 := t.Transf(v1)
	w1.Norm()
	w2 := t.Transf(v2)
	w2.Norm()
	w3 := t.Transf(v3)
	w3.Norm()
	x1 := graphics.Round(w1[0])
	y1 := graphics.Round(w1[1])
	x2 := graphics.Round(w2[0])
	y2 := graphics.Round(w2[1])
	x3 := graphics.Round(w3[0])
	y3 := graphics.Round(w3[1])
	img := graphics.NewImage(&scr)
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
