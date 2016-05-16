// Package main contains an example program that renders a simple triangle and
// updates the camera according to player movement.
package main

import (
	"github.com/amsibamsi/three/geom"
	mgeom "github.com/amsibamsi/three/math/geom"
	"github.com/amsibamsi/three/render"
	"github.com/amsibamsi/three/window"
	"time"
)

// main creates a new scene with a camera and a triangle, renders the scene,
// draws the result to a window and displays it. The middle point of the
// triangle continuously changes position relative to the current time.
func main() {
	win, err := window.NewWindow(1024, 768, "Three Move", true)
	if err != nil {
		panic(err)
	}
	defer window.Terminate()
	cam := render.NewDefCam()
	p := geom.NewTri4(-1, 0, -3, 0, 1, -3, 1, 0, -3)
	then := time.Now()
	now := time.Now()
	for close := false; !close; close = win.ShouldClose() || win.KeyDown(window.KeyQ) {
		then = now
		now = time.Now()
		dt := now.Sub(then)
		d := mgeom.Vec3{0, 0, 0}
		if win.KeyDown(window.KeyW) {
			d.Add(&cam.At)
		}
		if win.KeyDown(window.KeyS) {
			d.Sub(&cam.At)
		}
		if win.KeyDown(window.KeyA) {
			d.Add(mgeom.Cross(&cam.At, &cam.Up))
		}
		if win.KeyDown(window.KeyD) {
			d.Add(mgeom.Cross(&cam.Up, &cam.At))
		}
		d.Norm()
		d.Scale(dt.Seconds())
		cam.Eye.Add(&d)
		cam.Ar = float64(win.Width()) / float64(win.Height())
		t := cam.PerspTransf(win.Width(), win.Height())
		q := p.Transf(t)
		win.Clear()
		q.Draw(win)
		win.Update()
	}
}
