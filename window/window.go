package window

/*
#cgo pkg-config: glew glfw3

#include <GL/glew.h>
#include <GLFW/glfw3.h>
#include "window.h"
*/
import "C"

import (
	"errors"
	"github.com/amsibamsi/third"
	"runtime"
	"unsafe"
)

var (
	// To track GLFW initialization that only needs to be done once.
	glfwInitDone = false
)

// initGlfw initializes windowing by initializing GLFW. The current goroutine
// will be locked to the OS thread since most GLFW functions are not
// thread-safe.
func initGlfw() error {
	runtime.LockOSThread()
	err := C.initGlfw()
	if err != 1 {
		return errors.New("GLFW init failed")
	}
	return nil
}

// ensGlfwInit ensures GLFW has been initialized. It calls initGlfw() if not
// yet done so.
func ensGlfwInit() error {
	if glfwInitDone {
		return nil
	} else {
		return initGlfw()
	}
}

// newTex creates a new byte slice that holds the texture data.
func newTex(w, h int) []byte {
	return make([]byte, 3*w*h)
}

// Terminate destroys and cleans up all remaining windows and terminates
// windowing. Should be called at the end of a program or when no more
// windowing is needed.
func Terminate() {
	C.glfwTerminate()
}

// pollEvents registers pending event input and makes it ready to be queried.
func pollEvents() {
	C.glfwPollEvents()
}

// Window represents a graphical window.
type Window struct {

	// Width of the window content
	width int

	// Height of the window content
	height int

	// The actual window, a GLFW window
	glfwWin *C.GLFWwindow

	// Texture ID from OpenGL to draw the content to
	texId C.GLuint

	// Texture data. The format is based on OpenGL: 3 consecutive bytes build the
	// color for 1 pixel with red/green/blue values. Pixels are mapped to the
	// screen from left to right and top to bottom. So the texture starts at the
	// top left, first continues to the right and then breaks lines towards the
	// bottom.
	tex []byte
}

// NewWindow returns a new window. It initializes GLFW and GLEW, creates a new
// GLFW window with given width, height and title, initializes the texture
// data.
func NewWindow(width, height int, title string) (*Window, error) {
	if width < 0 {
		return nil, errors.New("Width must not be < 0")
	}
	if height < 0 {
		return nil, errors.New("Height must not be < 0")
	}
	err := ensGlfwInit()
	if err != nil {
		return nil, err
	}
	glfwWin := C.createWin(C.int(width), C.int(height), C.CString(title))
	if glfwWin == nil {
		Terminate()
		return nil, errors.New("Failed to create window")
	}
	errno := int(C.initGlew(glfwWin))
	if errno != 1 {
		return nil, errors.New("Failed to init GLEW")
	}
	C.initWin(glfwWin, C.int(width), C.int(height))
	tex := newTex(width, height)
	texId := C.createTex(
		glfwWin,
		unsafe.Pointer(&tex),
		C.int(width),
		C.int(height),
	)
	return &Window{width, height, glfwWin, texId, tex}, nil
}

// redraw draws the current texture data to the window. The content will first
// be shown on screen when the window is updated.
func (w *Window) redraw() {
	C.drawTex(
		w.glfwWin,
		unsafe.Pointer(&w.tex[0]),
		C.int(w.width),
		C.int(w.height),
	)
}

// refreshWait refreshes the window content on screen with the currently drawn
// data on the window. The call will block until buffers have been swapped.
func (w *Window) refreshWait() {
	C.glfwSwapBuffers(w.glfwWin)
}

// resize adapts the window and texture content to the current size of the
// window. It should be called periodically to adapt to GUI changes to the
// window. It checks the new window dimensions and if necessary creates a new
// texture with new size. Previously drawn content will be lost.
func (w *Window) resize() {
	var width, height int
	C.glfwGetWindowSize(
		w.glfwWin,
		(*C.int)(unsafe.Pointer(&width)),
		(*C.int)(unsafe.Pointer(&height)),
	)
	if width != w.width || height != w.height {
		w.width = width
		w.height = height
		w.tex = newTex(width, height)
		w.texId = C.resizeTex(
			w.glfwWin,
			w.texId,
			unsafe.Pointer(&w.tex[0]),
			C.int(width),
			C.int(height),
		)
		C.winResized(w.glfwWin, C.int(width), C.int(height))
	}
}

// Set works like Setxy, but for vectors.
func (w *Window) Set(v *third.Vec2, r, g, b byte) {
	if v[0] < 0 || v[0] > w.width || v[1] < 0 || v[1] > w.height {
		return
	}
	i := v[1]*3*w.width + v[0]*3
	w.tex[i] = r
	w.tex[i+1] = g
	w.tex[i+2] = b
}

// Setxy sets the texture color at the given position. If (x,y) lies not within
// the window nothing is drawn.
func (w *Window) Setxy(x, y int, r, g, b byte) {
	if x < 0 || x > w.width || y < 0 || y > w.height {
		return
	}
	i := y*3*w.width + x*3
	w.tex[i] = r
	w.tex[i+1] = g
	w.tex[i+2] = b
}

// Dot draws visible dot at the given coordinates. It's bigger than just one
// pixel.
func (w *Window) Dot(v *third.Vec2, r, g, b byte) {
	d := *v
	w.Set(&d, r, g, b)
	d[0] -= 1
	w.Set(&d, r, g, b)
	d[0] += 2
	w.Set(&d, r, g, b)
	d[0] -= 1
	d[1] -= 1
	w.Set(&d, r, g, b)
	d[1] += 2
	w.Set(&d, r, g, b)
}

// Line draws a line between two points. It's 1 pixel thick.
func (w *Window) Line(v1, v2 *third.Vec2, r, g, b byte) {
	// Always draw from left to right (x1 <= x2)
	if v1[0] > v2[0] {
		v1, v2 = v2, v1
	}
	dx := v2[0] - v1[0]
	dy := v2[1] - v1[1]
	var steps int
	if third.Abs(dx) > third.Abs(dy) {
		steps = dx
	} else {
		steps = third.Abs(dy)
	}
	xinc := float64(dx) / float64(steps)
	yinc := float64(dy) / float64(steps)
	x := float64(v1[0])
	y := float64(v1[1])
	for s := 0; s <= steps; s++ {
		w.Setxy(third.Round(x), third.Round(y), r, g, b)
		x += xinc
		y += yinc
	}
}

// Update updates the window. It does the following in order listed:
//
//   1. Draws the current content to the framebuffer
//   2. Waits for the content to be displayed by swapping buffers (V-Sync)
//   3. Adapts the window for any resizing
//   4. Polls events and makes them ready for processing
func (w *Window) Update() {
	w.redraw()
	w.refreshWait()
	w.resize()
	pollEvents()
}

// Clear clears the window content by setting all pixels to black.
func (w *Window) Clear() {
	for i, _ := range w.tex {
		w.tex[i] = 0
	}
}

// ShouldClose returns true if the window was requested to close by a GUI
// operation.
func (w *Window) ShouldClose() bool {
	should := C.glfwWindowShouldClose(w.glfwWin)
	return should != 0
}

// Width returns the currently set width of the window. This may not be up to
// date with the current GUI width of the window.
func (w *Window) Width() int {
	return w.width
}

// Height returns the currently set height of the window. This may not be up to
// date with the current GUI height of the window.
func (w *Window) Height() int {
	return w.height
}

// Destroy destroys the GLFW window.
func (w *Window) Destroy() {
	C.glfwDestroyWindow(w.glfwWin)
}
