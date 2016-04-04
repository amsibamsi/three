// Package main draws a simple triangle and stores it as JPEG file.
package main

import (
	"flag"
	"github.com/amsibamsi/3d/graphics"
	"image"
	"image/color"
	"os"
)

// main creates a new image, draws 3 dots and 3 lines onto the image, and then
// encodes the image as JPEG into a file. Optionally specify the filename.
func main() {
	var filename = flag.String("file", "triangle.jpg", "Filename to store image")
	flag.Parse()
	file, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	i := graphics.NewImage(500, 500)
	c := color.RGBA{255, 255, 0, 255}
	p := image.Point{250, 10}
	q := image.Point{490, 490}
	r := image.Point{10, 490}
	i.DrawDot(p, c)
	i.DrawDot(q, c)
	i.DrawDot(r, c)
	i.DrawLine(p, q, c)
	i.DrawLine(q, r, c)
	i.DrawLine(r, p, c)
	i.WriteJpeg(file)
}
