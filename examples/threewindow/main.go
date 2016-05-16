package main

import (
	"fmt"
	"github.com/amsibamsi/three/window"
	"os"
)

// main just runs the example with a visible window.
func main() {
	run(false)
}

// run initializes windowing, creates a new, waits for close event, destroys
// the window and terminates windowing.
func run(testing bool) {
	w, err := window.NewWindow(1024, 768, "Three Example", !testing)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer window.Terminate()
	for close := false; !close; close = w.ShouldClose() {
		w.Update()
		if testing {
			w.SetClose()
		}
	}
	w.Destroy()
}
