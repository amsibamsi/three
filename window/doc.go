// Package window covers creating windows, drawing a texture as windows content
// and processing input events.
//
// It needs GLFW3 and GLEW libraries installed on the system and calls them
// with the help of Cgo. For compilation header files are also required.
//
// Quickstart
//
// Windowing is currently not dead simple, the following procedure should
// always be satisfied:
//
//   - Call Init() before creating any window
//   - Call Terminate() at and of program (may use defer)
//   - When querying user input on a window first bring input data up to date
//     with PollEvents()
//   - Lock the main thread with runtime.LockOSThread() in the main program
//     (reason unknown, program will crash often otherwise)
package window
