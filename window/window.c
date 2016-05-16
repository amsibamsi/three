#include <stdio.h>
#include <GL/glew.h>
#include <GLFW/glfw3.h>

// Callback function for GLFW errors.
// Just print errors to stderr.
void glfwError(int err, const char* desc) {
  fprintf(stderr, "GLFW error: %d: %s\n", err, desc);
}

// Initialize GLFW.
// Returns 0 for success and otherwise 1. Errors will be handled by the error
// callback.
int initGlfw() {
  int err;
  glfwSetErrorCallback(glfwError);
  err = glfwInit();
  if (err != GL_TRUE) {
    return 0;
  }
  return 1;
}

// Create a GLFW window.
// Initializes the window with OpenGL API 2.1 and sets the swap interval to 1.
// Returns the window in case of success and otherwise NULL. Errors will be
// handled by the error callback.
GLFWwindow* createWin(int width,
                      int height,
                      char* title,
                      int visible) {
  GLFWwindow* win;
  glfwWindowHint(GLFW_CLIENT_API, GLFW_OPENGL_API);
  glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 2);
  glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 1);
  if (!visible) {
    glfwWindowHint(GLFW_VISIBLE, GL_FALSE);
  }
  win = glfwCreateWindow(width, height, title, NULL, NULL);
  if (win != NULL) {
    glfwMakeContextCurrent(win);
    glfwSwapInterval(1);
  }
  return win;
}

// Initializes GLEW on the given window.
// Makes the window and it's context ready for OpenGL calls. Returns 0 for
// success and otherwise 1. Errors will be printed to stderr.
int initGlew(GLFWwindow* win) {
  int err;
  glfwMakeContextCurrent(win);
  glewExperimental = GL_TRUE;
  err = glewInit();
  if (err != GLEW_OK) {
    fprintf(stderr, "GLEW init failed: %s\n", glewGetString(err));
    return 0;
  }
  return 1;
}

// Initializes a window after GLEW has been initialized.
// Sets the viewport to the current width and height.
void initWin(GLFWwindow* win,
             int width,
             int height) {
  glfwMakeContextCurrent(win);
  glViewport(0, 0, (GLsizei)width, (GLsizei)height);
}

// Adapts the window after it has been resized.
// Resets the viewport.
void winResized(GLFWwindow* win,
                int width,
                int height) {
  glfwMakeContextCurrent(win);
  glViewport(0, 0, (GLsizei)width, (GLsizei)height);
}

// Creates the texture to be drawn as content of the window.
// Generates a new texture and uploads the texture data. The format of the
// texture data must match what is hard coded here. Returns the ID for the
// newly generated texture object.
//
// TODO: Can texture be created without uploading data?
GLuint createTex(GLFWwindow* window,
                 GLvoid* data,
                 int width,
                 int height) {
  GLuint tex;
  glfwMakeContextCurrent(window);
  glGenTextures(1, &tex);
  glBindTexture(GL_TEXTURE_2D, tex);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST);
  glPixelStorei(GL_UNPACK_ALIGNMENT, 1);
  glPixelStorei(GL_PACK_ALIGNMENT, 1);
  glTexImage2D(GL_TEXTURE_2D,
               0,
               GL_RGB8,
               (GLsizei)width,
               (GLsizei)height,
               0,
               GL_RGB,
               GL_UNSIGNED_BYTE,
               data);
  return tex;
}

// Deletes a texture object.
void delTex(GLFWwindow* window,
            GLuint tex) {
  glDeleteTextures(1, &tex);
}

// Draws the texture with the given data.
// Clears the framebuffer, sets an orthogonal projection to cover the whole
// window content and draws a rectangle with the texture. The format of the
// texture data must match what is hard coded here. The result won't be
// shown on screen until buffers are swapped.
void drawTex(GLFWwindow* window,
             GLvoid* data,
             int width,
             int height) {
  glfwMakeContextCurrent(window);
  glPixelStorei(GL_UNPACK_ALIGNMENT, 1);
  glPixelStorei(GL_PACK_ALIGNMENT, 1);
  glTexSubImage2D(GL_TEXTURE_2D,
                  0,
                  0,
                  0,
                  (GLsizei)width,
                  (GLsizei)height,
                  GL_RGB,
                  GL_UNSIGNED_BYTE,
                  data);
  glClearColor(0.0, 0.0, 0.0, 0.0);
  glClear(GL_COLOR_BUFFER_BIT);
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  glOrtho(0.0, (GLdouble)(width-1), 0.0, (GLdouble)(height-1), -1.0, 1.0);
  glMatrixMode(GL_MODELVIEW);
  glLoadIdentity();
  glEnable(GL_TEXTURE_2D);
  glBegin(GL_QUADS);
  glTexCoord2i(0, 0);
  glVertex2i(0, height-1);
  glTexCoord2i(1, 0);
  glVertex2i(width-1, height-1);
  glTexCoord2i(1, 1);
  glVertex2i(width-1, 0);
  glTexCoord2i(0, 1);
  glVertex2i(0, 0);
  glEnd();
}

// Resize the texture.
// Deletes the old texture and creates a new one with the given dimensions.
// Returns the ID of the new texture object.
GLuint resizeTex(GLFWwindow* win,
                 GLuint tex,
                 GLvoid* data,
                 int width,
                 int height) {
  GLuint newTex;
  newTex = createTex(win, data, width, height);
  delTex(win, tex);
  return newTex;
}
