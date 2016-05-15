// Package window provides windowing, drawing and processing input events.
//
// It needs GLFW3 and GLEW libraries installed on the system and calls them
// with the help of Cgo. For compilation header files are also required.
//
// Quickstart
//
//   1. Create new window with NewWindow()
//   2. Periodically:
//      - Draw to the window with Set()
//      - Call Update()
//      - Stop if ShouldClose() returns true
//   3. Call Destroy() on the window
//   4. Call Terminate() at end of program (defer after window was created)
//
// Implementation
//
// A window holds a GLFW window and some texture data. The texture matches
// exactly the number pixels of the window content. When drawing with OpenGL a
// single rectangle with this texture is created and drawn with an orthographic
// projection to fill the whole window. The texture is uploaded and drawn every
// frame. The only reason for choosing OpenGL was that GLFW and GLEW present
// platform independent and realtively easy to use C APIs that can be used from
// Go. The performance benefit from offloading graphics to the GPU is not
// really used.
package window
