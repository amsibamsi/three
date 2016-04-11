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
// OpenGL in version 2.1. Should only be called once before creating a window.
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
// prints the error to stderr.
//export GlfwError
func GlfwError(error C.int, description *C.char) {
	fmt.Fprintf(
		os.Stderr,
		"GLFW error %d: %s\n",
		int(error), C.GoString(description),
	)
}

//
type Window struct {

	//
	Width C.GLsizei

	//
	Height C.GLsizei

	//
	GlfwWin *C.GLFWwindow

	//
	TexId C.GLuint

	//
	Tex []byte
}

//
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
	window := Window{C.GLsizei(w), C.GLsizei(h), glfwWin, texId, tex}
	C.glfwMakeContextCurrent(glfwWin)
	C.glewExperimental = C.GL_TRUE
	err := C.glewInit()
	if err != C.GLEW_OK {
		errstr := (*C.char)(unsafe.Pointer(C.glewGetErrorString(err)))
		msg := fmt.Sprintf("Failed to initialize GLEW: %s", C.GoString(errstr))
		return nil, errors.New(msg)
	}
	return &window, nil
}

//
func (w *Window) Draw() {
	C.glfwMakeContextCurrent(w.GlfwWin)
	C.glPixelStorei(C.GL_UNPACK_ALIGNMENT, 1)
	C.glPixelStorei(C.GL_PACK_ALIGNMENT, 1)
	C.glBindTexture(C.GL_TEXTURE_2D, w.TexId)
	C.glTexImage2D(
		C.GL_TEXTURE_2D,
		0,
		C.GL_RGB8,
		w.Width,
		w.Height,
		0,
		C.GL_RGB,
		C.GL_UNSIGNED_BYTE,
		unsafe.Pointer(&w.Tex[0]),
	)
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MIN_FILTER, C.GL_NEAREST)
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MAG_FILTER, C.GL_NEAREST)
	C.glClearColor(0.0, 0.0, 0.0, 0.0)
	C.glClear(C.GL_COLOR_BUFFER_BIT)
	C.glMatrixMode(C.GL_PROJECTION)
	C.glLoadIdentity()
	C.glOrtho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)
	C.glMatrixMode(C.GL_MODELVIEW)
	C.glLoadIdentity()
	C.glEnable(C.GL_TEXTURE_2D)
	C.glBegin(C.GL_QUADS)
	C.glTexCoord2f(0.0, 0.0)
	C.glVertex2f(-1.0, 1.0)
	C.glTexCoord2f(1.0, 0.0)
	C.glVertex2f(1.0, 1.0)
	C.glTexCoord2f(0.0, 1.0)
	C.glVertex2f(1.0, -1.0)
	C.glTexCoord2f(1.0, 1.0)
	C.glVertex2f(-1.0, -1.0)
	C.glEnd()
	C.glfwSwapBuffers(w.GlfwWin)
}

//
func (w *Window) Destroy() {
	C.glfwDestroyWindow(w.GlfwWin)
}

//
func (w *Window) ShouldClose() bool {
	should := C.glfwWindowShouldClose(w.GlfwWin)
	return should != 0
}

//
func PollEvents() {
	C.glfwPollEvents()
}

//
func Terminate() {
	C.glfwTerminate()
}
