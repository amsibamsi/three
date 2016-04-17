// Package main contains an example program that draws some animated graphics
// in a window.
package main

import (
	"fmt"
	"github.com/amsibamsi/third"
	"github.com/amsibamsi/third/window"
	"math/rand"
	"os"
)

// main initializes windowing, creates a new, continuously draws some pixels,
// waits for close event, destroys the window and terminates windowing.
func main() {
	width := 1024
	height := 768
	w, err := window.NewWindow(width, height, "Third Animate")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer window.Terminate()
	x := width / 2
	y := height / 2
	r := rand.New(rand.NewSource(0))
	for close := false; !close; close = w.ShouldClose() {
		w.Setxy(x, y, 0, 0, 0)
		width = w.Width()
		height = w.Height()
		x = x + r.Intn(3) - 1
		x = third.Min(width, third.Max(0, x))
		y = y + r.Intn(3) - 1
		y = third.Min(height, third.Max(0, y))
		w.Setxy(x, y, 255, 0, 0)
		w.Update()
	}
	w.Destroy()
}
