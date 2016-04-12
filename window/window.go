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
	//"math/rand"
	"os"
	//"time"
	"unsafe"
)

// Init initializes GLFW. Returns a generic error if initialization failed.
// The error will be printed to stderr via a callback function. Initializes
// OpenGL in version 2.1. Should only be called once before creating the first
// window.
func Init() error {
	C.setGlfwErrorCallback()
	if C.glfwInit() != C.GL_TRUE {
		return errors.New("Failed to initialize GLFW")
	}
	C.glfwWindowHint(C.GLFW_CLIENT_API, C.GLFW_OPENGL_API)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 2)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 1)
	return nil
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

// Window represents a graphical window.
type Window struct {

	// Width of the window content
	Width int

	// Height of the window content
	Height int

	// The actual window, a GLFW window
	GlfwWin *C.GLFWwindow

	// Texture ID from OpenGL to draw the content to
	TexId int

	// Texture data. The format is based on OpenGL: 3 consecutive bytes build the
	// color for 1 pixel, red/green/blue values. Pixels are mapped to the screen
	// from left to right and top to bottom. So the texture starts at the top
	// left, first continues to the right and then breaks lines towards the
	// bottom.
	Tex []byte
}

// NewWindow returns a new window. It creates a new GLFW window, initializes
// the texture data and initializes GLEW with the window's context.
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
	var texId C.GLuint
	C.glGenTextures(1, &texId)
	tex := make([]byte, 3*w*h)
	window := Window{w, h, glfwWin, int(texId), tex}
	C.glfwMakeContextCurrent(glfwWin)
	C.glfwSwapInterval(0)
	C.glewExperimental = C.GL_TRUE
	err := C.glewInit()
	if err != C.GLEW_OK {
		errstr := (*C.char)(unsafe.Pointer(C.glewGetErrorString(err)))
		msg := fmt.Sprintf("Failed to initialize GLEW: %s", C.GoString(errstr))
		return nil, errors.New(msg)
	}
	C.glBindTexture(C.GL_TEXTURE_2D, C.GLuint(texId))
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MIN_FILTER, C.GL_NEAREST)
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MAG_FILTER, C.GL_NEAREST)
	C.glPixelStorei(C.GL_UNPACK_ALIGNMENT, 1)
	C.glPixelStorei(C.GL_PACK_ALIGNMENT, 1)
	C.glTexImage2D(
		C.GL_TEXTURE_2D,
		0,
		C.GL_RGB8,
		C.GLsizei(w),
		C.GLsizei(h),
		0,
		C.GL_RGB,
		C.GL_UNSIGNED_BYTE,
		unsafe.Pointer(&tex[0]),
	)
	return &window, nil
}

// Draw draws the current texture data of the window.
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
func (w *Window) Draw() {
	width := C.GLint(w.Width)
	height := C.GLint(w.Height)
	C.glfwMakeContextCurrent(w.GlfwWin)
	C.glPixelStorei(C.GL_UNPACK_ALIGNMENT, 1)
	C.glPixelStorei(C.GL_PACK_ALIGNMENT, 1)
	C.glTexSubImage2D(
		C.GL_TEXTURE_2D,
		0,
		0,
		0,
		C.GLsizei(width),
		C.GLsizei(height),
		C.GL_RGB,
		C.GL_UNSIGNED_BYTE,
		unsafe.Pointer(&w.Tex[0]),
	)
	C.glClearColor(0.0, 0.0, 0.0, 0.0)
	C.glClear(C.GL_COLOR_BUFFER_BIT)
	C.glMatrixMode(C.GL_PROJECTION)
	C.glLoadIdentity()
	C.glOrtho(0.0, C.GLdouble(width-1), 0.0, C.GLdouble(height-1), -1.0, 1.0)
	C.glMatrixMode(C.GL_MODELVIEW)
	C.glLoadIdentity()
	C.glEnable(C.GL_TEXTURE_2D)
	C.glBegin(C.GL_QUADS)
	C.glTexCoord2i(0, 0)
	C.glVertex2i(0, height-1)
	C.glTexCoord2i(1, 0)
	C.glVertex2i(width-1, height-1)
	C.glTexCoord2i(1, 1)
	C.glVertex2i(width-1, 0)
	C.glTexCoord2i(0, 1)
	C.glVertex2i(0, 0)
	C.glEnd()
	C.glfwSwapBuffers(w.GlfwWin)
}

// Destroy destroys the GLFW window.
func (w *Window) Destroy() {
	C.glfwDestroyWindow(w.GlfwWin)
}

// ShouldClose returns true if the window was requested to close by a GUI
// event.
func (w *Window) ShouldClose() bool {
	should := C.glfwWindowShouldClose(w.GlfwWin)
	return should != 0
}

// PollEvents registers pending event input and makes it ready to be queried.
func PollEvents() {
	C.glfwPollEvents()
}

// Terminate should be called once on program termination to signal GLFW to
// clean up.
func Terminate() {
	C.glfwTerminate()
}
