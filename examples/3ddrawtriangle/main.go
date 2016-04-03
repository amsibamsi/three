//
package main

import (
	"github.com/amsibamsi/3d/graphics"
	"image"
	"image/color"
)

//
func main() {
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
	i.WriteJpeg("triangle.jpeg")
}
