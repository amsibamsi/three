// Package main contains an example program that draws some animated graphics
// in a window.
package main

import (
	"fmt"
	"github.com/amsibamsi/third/window"
	"os"
	"runtime"
	"time"
)

// main initializes windowing, creates a new, continuously draws some pixels,
// waits for close event, destroys the window and terminates windowing.
func main() {
	runtime.LockOSThread()
	err := window.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer window.Terminate()
	w, err := window.NewWindow(1024, 768, "Third Animate")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	i := 0
	for close := false; !close; close = w.ShouldClose() {
		w.Tex[i+0] = 0
		w.Tex[i+3] = 0
		w.Tex[i+6] = 255
		w.Tex[i+9] = 255
		i = (i + 6) % 300
		w.Draw()
		time.Sleep(16 * time.Millisecond)
		window.PollEvents()
	}
	w.Destroy()
}
