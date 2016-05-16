// Package main contains an example program that draws some animated graphics
// in a window.
package main

import (
	"fmt"
	tmath "github.com/amsibamsi/three/math"
	"github.com/amsibamsi/three/window"
	"math/rand"
	"os"
)

// main initializes windowing, creates a new, continuously draws some pixels,
// waits for close event, destroys the window and terminates windowing.
func main() {
	width := 1024
	height := 768
	w, err := window.NewWindow(width, height, "Three Animate", true)
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
		x = tmath.Mini(width, tmath.Maxi(0, x))
		y = y + r.Intn(3) - 1
		y = tmath.Mini(height, tmath.Maxi(0, y))
		w.Setxy(x, y, 255, 0, 0)
		w.Update()
	}
	w.Destroy()
}
