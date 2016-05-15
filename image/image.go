package image

import (
	tmath "github.com/amsibamsi/three/math"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
)

// Image is a simple 2D image. It wraps an RGBA image from the standard image
// package.
type Image struct {
	Rgba image.RGBA
}

// NewImage returns a new image with the given screen width and height and
// black background.
func NewImage(w, h int) *Image {
	rect := image.Rect(0, 0, w, h)
	rgba := image.NewRGBA(rect)
	bg := image.Uniform{color.Black}
	draw.Draw(rgba, rgba.Bounds(), &bg, image.Point{}, draw.Src)
	return &Image{*rgba}
}

// DrawDot draws a clearly visible dot (more than 1 pixel) at (x,y) with the
// given color.
func (img *Image) DrawDot(x, y int, c color.Color) {
	r := img.Rgba
	ind := []int{
		-2, 0,
		-1, -1,
		-1, 0,
		-1, 1,
		0, -1,
		0, 0,
		0, 1,
		1, -1,
		1, 0,
		1, 1,
		2, 0,
	}
	for i := 0; i < len(ind); i += 2 {
		r.Set(x+ind[i], y+ind[i+1], c)
	}
}

// DrawLine draws a 1 pixel thick line between the (x1,y1) and (x2,y2) with the
// given color.
func (img *Image) DrawLine(x1, y1, x2, y2 int, c color.Color) {
	r := img.Rgba
	// Always draw from left to right (x1 <= x2)
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}
	dx := x2 - x1
	dy := y2 - y1
	var steps int
	if tmath.Absi(dx) > tmath.Absi(dy) {
		steps = tmath.Absi(dx)
	} else {
		steps = tmath.Absi(dy)
	}
	xinc := float64(dx) / float64(steps)
	yinc := float64(dy) / float64(steps)
	x := float64(x1)
	y := float64(y1)
	for s := 0; s <= steps; s++ {
		r.Set(tmath.Round(x), tmath.Round(y), c)
		x += xinc
		y += yinc
	}
}

// WritePng stores the image in PNG format to the given writer and returns the
// error from png.Encode() if any.
func (img *Image) WritePng(w io.Writer) error {
	return png.Encode(w, &img.Rgba)
}
