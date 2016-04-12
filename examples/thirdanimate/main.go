// Package main contains an example program that draws some animated graphics
// in a window.
package main

import (
	"fmt"
	"github.com/amsibamsi/third/window"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// main initializes windowing, creates a new, continuously draws some pixels,
// waits for close event, destroys the window and terminates windowing.
func main() {
	width := 1024
	height := 768
	runtime.LockOSThread()
	err := window.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer window.Terminate()
	w, err := window.NewWindow(width, height, "Third Animate")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	x := 0
	y := 0
	r := rand.New(rand.NewSource(0))
	frames := 0
	elapsed := int64(0)
	then := time.Now()
	now := time.Now()
	for close := false; !close; close = w.ShouldClose() {
		w.Tex[x*3+y*3*width] = 0
		n := r.Intn(4)
		switch n {
		case 0:
			x++
		case 1:
			y++
		case 2:
			x--
		case 3:
			y--
		}
		if x < 0 {
			x = width - x
		}
		if y < 0 {
			y = height - y
		}
		x %= width
		y %= height
		w.Tex[x*3+y*3*width] = 255
		w.Draw()
		//time.Sleep(10 * time.Millisecond)
		window.PollEvents()
		frames++
		now = time.Now()
		elapsed += now.Sub(then).Nanoseconds()
		then = now
		if elapsed > 1e9 {
			fmt.Printf("fps: %f\n", float64(frames)/(float64(elapsed)/1e9))
			elapsed = 0
			frames = 0
		}
	}
	w.Destroy()
}
