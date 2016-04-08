#include "_cgo_export.h"

void glfwErrorCallback(int error, const char* description) {
  GlfwError(error, (char*)description);
}

void setGlfwErrorCallback() {
  glfwSetErrorCallback(glfwErrorCallback);
}
