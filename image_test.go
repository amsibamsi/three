package third

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestNewImage(t *testing.T) {
	img := NewImage(&Screen{100, 100})
	rgba := img.Rgba
	bounds := rgba.Bounds()
	rect := image.Rect(0, 0, 100, 100)
	if bounds != rect {
		t.Errorf("expected '%v' but got '%v'", rect, bounds)
	}
}

func TestDrawDot(t *testing.T) {
	img := NewImage(&Screen{100, 100})
	rgba := img.Rgba
	col1 := color.RGBA{200, 111, 38, 1}
	img.DrawDot(50, 50, col1)
	col2 := rgba.At(50, 50)
	if col1 != col2 {
		t.Errorf("expected '%v' but got '%v'", col1, col2)
	}
}

func TestDrawLine1(t *testing.T) {
	img := NewImage(&Screen{100, 100})
	rgba := img.Rgba
	col := color.RGBA{1, 2, 3, 4}
	img.DrawLine(10, 10, 12, 12, col)
	should := [3]color.Color{col, col, col}
	is := [3]color.Color{
		rgba.At(10, 10),
		rgba.At(11, 11),
		rgba.At(12, 12),
	}
	if is != should {
		t.Errorf("expected '%v' but got '%v'", should, is)
	}
}

func TestDrawLine2(t *testing.T) {
	img := NewImage(&Screen{100, 100})
	rgba := img.Rgba
	col := color.RGBA{1, 2, 3, 4}
	img.DrawLine(10, 10, 13, 11, col)
	should := [4]color.Color{col, col, col, col}
	is := [4]color.Color{
		rgba.At(10, 10),
		rgba.At(11, 10),
		rgba.At(12, 11),
		rgba.At(13, 11),
	}
	if is != should {
		t.Errorf("expected '%v' but got '%v'", should, is)
	}
}

func TestDrawLine3(t *testing.T) {
	img := NewImage(&Screen{100, 100})
	rgba := img.Rgba
	col := color.RGBA{1, 2, 3, 4}
	img.DrawLine(10, 10, 9, 9, col)
	should := [2]color.Color{col, col}
	is := [2]color.Color{
		rgba.At(9, 9),
		rgba.At(10, 10),
	}
	if is != should {
		t.Errorf("expected '%v' but got '%v'", should, is)
	}
}

func TestWritePng(t *testing.T) {
	scr := Screen{100, 100}
	img1 := NewImage(&scr)
	col1 := color.RGBA{0, 11, 0, 255}
	img1.DrawDot(4, 5, col1)
	var buf bytes.Buffer
	img1.WritePng(&buf)
	img2, _ := png.Decode(&buf)
	col2 := img2.At(4, 5)
	_, r1, _, _ := col1.RGBA()
	_, r2, _, _ := col2.RGBA()
	if r2 != r1 {
		t.Errorf("expected '%v' but got '%v'", r1, r2)
	}
}
