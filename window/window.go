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
	"runtime"
	"unsafe"
)

var (
	glfwInitDone = false
)

// initGlfw initializes windowing by initializing GLFW. Any error will be
// printed to stderr. The current goroutine will be locked to the OS thread
// since most GLFW functions are not thread-safe.
func initGlfw() error {
	runtime.LockOSThread()
	err := C.initGlfw()
	if err != 1 {
		return errors.New("GLFW init failed")
	}
	return nil
}

//
func ensGlfwInit() error {
	if glfwInitDone {
		return nil
	} else {
		return initGlfw()
	}
}

//
func newTex(w, h int) []byte {
	return make([]byte, 3*w*h)
}

// Terminate should be called once on program termination to signal GLFW to
// clean up. Destroys all remaining windows.
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
	// color for 1 pixel, red/green/blue values. Pixels are mapped to the screen
	// from left to right and top to bottom. So the texture starts at the top
	// left, first continues to the right and then breaks lines towards the
	// bottom.
	tex []byte
}

// NewWindow returns a new window. It creates a new GLFW window with given
// width, height and title, initializes the texture data and initializes
// GLEW with the window's context.
func NewWindow(w, h int, t string) (*Window, error) {
	if w < 0 {
		return nil, errors.New("Width must not be < 0")
	}
	if h < 0 {
		return nil, errors.New("Height must not be < 0")
	}
	err := ensGlfwInit()
	if err != nil {
		return nil, err
	}
	glfwWin := C.createWin(
		C.int(w),
		C.int(h),
		C.CString(t),
	)
	if glfwWin == nil {
		Terminate()
		return nil, errors.New("Failed to create window")
	}
	errno := int(C.initGlew(glfwWin))
	if errno != 1 {
		return nil, errors.New("Failed to init GLEW")
	}
	C.initWin(glfwWin, C.int(w), C.int(h))
	tex := newTex(w, h)
	texId := C.createTex(
		glfwWin,
		unsafe.Pointer(&tex),
		C.int(w),
		C.int(h),
	)
	return &Window{w, h, glfwWin, texId, tex}, nil
}

// Show draws the current texture data of the window.
//
// This is implemented with basic OpenGL 2:
//
//   - Upload the texture data to the GPU
//   - Clear the framebuffer
//   - Set up an orthographic projection
//   - Create a rectangle to cover the whole screen
//   - Map the texture to the rectangle with the right orientation
//   - Swap buffers to draw the new content
//
// Important:
//
//   - GLFW window must first be made current
//   - Order of glVertex2f and glTexCoord2f determines orientation
func (w *Window) redraw() {
	C.drawTex(
		w.glfwWin,
		unsafe.Pointer(&w.tex[0]),
		C.int(w.width),
		C.int(w.height),
	)
}

//
func (w *Window) refreshWait() {
	C.glfwSwapBuffers(w.glfwWin)
}

//
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

//
func (w *Window) Set(x, y int, r, g, b byte) {
	xc := ((x % w.width) + w.width) % w.width
	yc := ((y % w.height) + w.height) % w.height
	i := yc*3*w.width + xc*3
	w.tex[i] = r
	w.tex[i+1] = g
	w.tex[i+2] = b
}

//
func (w *Window) Update() {
	w.redraw()
	w.refreshWait()
	w.resize()
	pollEvents()
}

// ShouldClose returns true if the window was requested to close by a GUI
// event.
func (w *Window) ShouldClose() bool {
	should := C.glfwWindowShouldClose(w.glfwWin)
	return should != 0
}

//
func (w *Window) Width() int {
	return w.width
}

//
func (w *Window) Height() int {
	return w.height
}

// Destroy destroys the GLFW window.
func (w *Window) Destroy() {
	C.glfwDestroyWindow(w.glfwWin)
}
