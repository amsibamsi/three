// Package main renders a simple triangle and stores it as JPEG file.
package main

import (
	"flag"
	"github.com/amsibamsi/3d/graphics"
	"image"
	"image/color"
	"os"
)

// main creates a new scene with a camera and 3 vectors, renders the scene and
// encodes the image as JPEG into a file. Optionally specify the filename.
func main() {
	var filename = flag.String("file", "triangle.jpg", "Filename to store image")
	flag.Parse()
	file, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	cam := graphics.NewDefCam()
	v1 := graphics.NewVec4(-1, -1, -3)
	v2 := graphics.NewVec4(0, 1, -5)
	v3 := graphics.NewVec4(1, -1, -3)
	t := cam.PerspTransf(&graphics.Screen{500, 500})
	w1 := t.Transf(v1)
	w1.Norm()
	w2 := t.Transf(v2)
	w2.Norm()
	w3 := t.Transf(v3)
	w3.Norm()
	img := graphics.NewImage(500, 500)
	c := color.RGBA{255, 255, 0, 255}
	p := image.Point{graphics.Round(w1[0]), graphics.Round(w1[1])}
	q := image.Point{graphics.Round(w2[0]), graphics.Round(w2[1])}
	r := image.Point{graphics.Round(w3[0]), graphics.Round(w3[1])}
	img.DrawDot(p, c)
	img.DrawDot(q, c)
	img.DrawDot(r, c)
	img.DrawLine(p, q, c)
	img.DrawLine(q, r, c)
	img.DrawLine(r, p, c)
	img.WriteJpeg(file)
}
