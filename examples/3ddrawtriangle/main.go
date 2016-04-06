// Package main draws a simple triangle and stores it in a file.
package main

import (
	"flag"
	"github.com/amsibamsi/3d/graphics"
	"image/color"
	"os"
)

// main creates a new image, draws 3 dots and 3 lines onto the image, and then
// encodes the image as PNG into a file. Optionally specify the filename.
func main() {
	var filename = flag.String("file", "triangle.png", "Filename to store image")
	flag.Parse()
	file, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s := graphics.Screen{500, 500}
	i := graphics.NewImage(&s)
	c := color.RGBA{255, 255, 0, 255}
	x1 := 250
	y1 := 10
	x2 := 490
	y2 := 490
	x3 := 10
	y3 := 490
	i.DrawDot(x1, y1, c)
	i.DrawDot(x2, y2, c)
	i.DrawDot(x3, y3, c)
	i.DrawLine(x1, y1, x2, y2, c)
	i.DrawLine(x2, y2, x3, y3, c)
	i.DrawLine(x3, y3, x1, y1, c)
	i.WritePng(file)
}
