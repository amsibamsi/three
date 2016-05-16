void glfwError(int err, const char* desc);
int initGlfw();
GLFWwindow* createWin(int width, int height, char* title, int visible);
int initGlew(GLFWwindow* win);
void initWin(GLFWwindow* win, int width, int height);
void winResized(GLFWwindow* win, int width, int height);

GLuint createTex(GLFWwindow* window, GLvoid* data, int width, int height);
void delTex(GLFWwindow* window, GLuint tex);
void drawTex(GLFWwindow* window, GLvoid* data, int width, int height);
GLuint resizeTex(GLFWwindow* win, GLuint tex, GLvoid* data, int width, int height);
