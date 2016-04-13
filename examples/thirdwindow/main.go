// Package main contains an example program that opens a window and waits for
// an event to close it again.
package main

import (
	"fmt"
	"github.com/amsibamsi/third/window"
	"os"
)

// main initializes windowing, creates a new, waits for close event, destroys
// the window and terminates windowing.
func main() {
	w, err := window.NewWindow(1024, 768, "Third Example")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer window.Terminate()
	for close := false; !close; close = w.ShouldClose() {
		w.Update()
	}
	w.Destroy()
}
