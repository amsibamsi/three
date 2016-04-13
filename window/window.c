#include "_cgo_export.h"

void glfwErrorCallback(int error, const char* description) {
  GlfwError(error, (char*)description);
}

void setGlfwErrorCallback() {
  glfwSetErrorCallback(glfwErrorCallback);
}

// TODO: Can texture be created without uploading data?
GLuint createTexture(GLFWwindow* window,
                     const GLvoid* data,
                     const GLint width,
                     const GLint height) {
  GLuint tex;
  glGenTextures(1, &tex);
  glfwMakeContextCurrent(window);
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

void drawTexture(GLFWwindow* window,
                 const GLvoid* data,
                 const GLint width,
                 const GLint height) {
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
