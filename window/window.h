void setGlfwErrorCallback();

GLuint createTexture(GLFWwindow* window,
                     const GLvoid* data,
                     const GLint width,
                     const GLint height);
void drawTexture(GLFWwindow* window,
                 const GLvoid* data,
                 const GLint width,
                 const GLint height);
