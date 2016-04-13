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
	"fmt"
	"os"
	"runtime"
	"unsafe"
)

// init initializes GLFW.  Any error will be handled by GlfwError().
// Initializes OpenGL in version 2.1. The calling thread will be locked since
// most GLFW functions are not thread-safe.
func init() {
	runtime.LockOSThread()
	C.setGlfwErrorCallback()
	err := C.glfwInit()
	if err != C.GL_TRUE {
		GlfwError(err, C.CString("Failed to initialize GLFW"))
	}
	C.glfwWindowHint(C.GLFW_CLIENT_API, C.GLFW_OPENGL_API)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 2)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 1)
}

// GlfwError is used as callback from GLFW to handle errors. Currently just
// prints the error to stderr. This function must be exported because Cgo can
// not register a Go function as callback in C, only C functions.
//export GlfwError
func GlfwError(error C.int, description *C.char) {
	fmt.Fprintf(
		os.Stderr,
		"GLFW error %d: %s\n",
		int(error), C.GoString(description),
	)
}

// Terminate should be called once on program termination to signal GLFW to
// clean up. Destroys all remaining windows.
func Terminate() {
	C.glfwTerminate()
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
	glfwWin := C.glfwCreateWindow(
		C.int(w),
		C.int(h),
		C.CString(t),
		nil,
		nil,
	)
	if glfwWin == nil {
		C.glfwTerminate()
		return nil, errors.New("Failed to open GLFW window")
	}
	C.glfwMakeContextCurrent(glfwWin)
	C.glfwSwapInterval(1)
	C.glewExperimental = C.GL_TRUE
	err := C.glewInit()
	if err != C.GLEW_OK {
		errstr := (*C.char)(unsafe.Pointer(&err))
		msg := fmt.Sprintf("Failed to initialize GLEW: %s", C.GoString(errstr))
		return nil, errors.New(msg)
	}
	C.glViewport(0, 0, C.GLsizei(w), C.GLsizei(h))
	tex := make([]byte, 3*w*h)
	texId := C.createTexture(
		glfwWin,
		unsafe.Pointer(&tex),
		C.GLint(w),
		C.GLint(h),
	)
	window := Window{w, h, glfwWin, texId, tex}
	return &window, nil
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
func (w *Window) refreshContent() {
	C.drawTexture(
		w.glfwWin,
		unsafe.Pointer(&w.tex[0]),
		C.GLint(w.width),
		C.GLint(w.height),
	)
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
		C.glViewport(0, 0, C.GLsizei(width), C.GLsizei(height))
		C.glDeleteTextures(1, &w.texId)
		tex := make([]byte, 3*width*height)
		w.tex = tex
		w.texId = C.createTexture(
			w.glfwWin,
			unsafe.Pointer(&w.tex[0]),
			C.GLint(width),
			C.GLint(height),
		)
		w.width = width
		w.height = height
	}
}

// pollEvents registers pending event input and makes it ready to be queried.
func pollEvents() {
	C.glfwPollEvents()
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
	w.refreshContent()
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
