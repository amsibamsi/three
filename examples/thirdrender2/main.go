// Package main contains an example program that renders a simple triangle that
// continuously changes coordinates.
package main

import (
	"github.com/amsibamsi/third"
	"github.com/amsibamsi/third/geom"
	"github.com/amsibamsi/third/window"
	"math"
	"time"
)

// main creates a new scene with a camera and a triangle, renders the scene,
// draws the result to a window and displays it. The middle point of the
// triangle continuously changes position relative to the current time.
func main() {
	win, err := window.NewWindow(1024, 768, "Third Render 2")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()
	cam := third.NewDefCam()
	p := geom.NewTri4(-1, 0, -3, 0, 1, -3, 1, 0, -3)
	for close := false; !close; close = win.ShouldClose() {
		now := time.Now()
		m := &p[1][1]
		*m = math.Sin(float64(now.UnixNano()) / 1e9)
		c := &cam.At[0]
		*c = math.Cos(float64(now.UnixNano()) / 1e9)
		w := win.Width()
		h := win.Height()
		cam.Ar = float64(w) / float64(h)
		t := cam.PerspTransf(w, h)
		q := p.Transf(t)
		win.Clear()
		q.Draw(win)
		win.Update()
	}
}
