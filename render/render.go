package render

import (
	"github.com/amsibamsi/three/math/geom"
	"github.com/amsibamsi/three/window"
	"math"
)

func TranslTransf(v *geom.Vec3) *geom.Mat4 {
	return &geom.Mat4{
		1, 0, 0, v[0],
		0, 1, 0, v[1],
		0, 0, 1, v[2],
		0, 0, 0, 1,
	}
}

// CoordTransf returns a new matrix that transforms from the orthonormal basis
// given by the 3 argument axes to the standard basis. It is used to to
// transform vectors from the world to the camera view.
//
// Reference is the standard basis with the axes (1,0,0), (0,1,0) and (0,0,1).
// A new basis is formed from the argument axes given in standard coordinates.
// A vector is also given in standard coordinates but interpreted in the view
// of the argument basis. The origin is the same for both bases. Instead of
// rotating we can also project any vector onto the new basis: E.g. the x
// coordinate of a vector (vx,vy,vz) in the new basis with x axis (ax,ay,az) is
// the dot product vx*ax + vy*ay + vz*az. The same is true for y and z
// coordinates. So the transformation is a matrix where the axes of the
// argument basis are the rows of the matrix. We fill up with 0 and 1 to make
// the matrix homogeneous.
func CoordTransf(x, y, z *geom.Vec3) *geom.Mat4 {
	return &geom.Mat4{
		x[0], x[1], x[2], 0,
		y[0], y[1], y[2], 0,
		z[0], z[1], z[2], 0,
		0, 0, 0, 1,
	}
}

// Camera describes a view in space. It is used to create a 2D image from the
// scene.
//
// The coordinate system of a camera is as follows:
//   - Center at Eye
//   - Positive y along Up
//   - Negative z along At
//   - Positive x along the cross product of y and z (to the right)
//
// Windowing
//
// If using a camera to render an image for a window it is useful to use
// UpdateAr() to update the aspect ratio to the current window size and
// PerspTransfWin() to get the projection matrix for the window.
type Camera struct {

	// Eye is the position of the eye to look from. For each object to project a
	// virtual line to the eye is drawn.
	Eye geom.Vec3

	// At is the direction to look at from the eye.
	At geom.Vec3

	// Up determines the orientation of the view. Up not being perpendicular to
	// At results in the same orientation as if Up was first projected to the
	// normal plane of At through Eye.
	Up geom.Vec3

	// Near is the distance from the eye in the looking direction where the
	// orthogonal plane is set to project onto.
	Near float64

	// Far is the distance from the eye in the looking direction where an
	// orthogonal plane is drawn. Everything beyond that plane is not projected.
	Far float64

	// Fov is the horizontal field of view in radian degrees.
	Fov float64

	// Ar is the aspect ratio of width to height.
	Ar float64
}

// NewDefCam returns a new camera with default settings.
func NewDefCam() *Camera {
	return &Camera{
		Eye:  geom.Vec3{0, 0, 0},
		At:   geom.Vec3{0, 0, -1},
		Up:   geom.Vec3{0, 1, 0},
		Near: 1.0,
		Far:  100.0,
		Fov:  math.Pi / 2,
		Ar:   1.0,
	}
}

// CamAxes returns the 3 axes that make up the orthonormal basis of the
// camera's right-handed coordinate system.
func (c *Camera) CamAxes() (*geom.Vec3, *geom.Vec3, *geom.Vec3) {
	z := c.Eye
	z.Sub(&c.At)
	z.Norm()
	x := geom.Cross(&c.Up, &z)
	x.Norm()
	// Recompute up, c.Up might not be perpendicular or normalized
	y := geom.Cross(&z, x)
	return x, y, &z
}

// CamTransf returns a new matrix that transforms from world coordinates into
// view coordinates of the camera. Any object in world coordinates is viewed
// through the camera that is also given in world coordinates. To get the
// objects's camera coordinates we transform the camera coordinate system to
// match the world coordinate system. Any object we transform the same way and
// get standard coordinates relative to the camera view. It is the inverse
// transformation of getting from the world view to the camera view. That means
// we first translate the eye of the camera to the origin and then rotate the
// camera until it matches the world view. Instead of rotating we can use the
// coordinate transformation function CoordTransf.
func (c *Camera) CamTransf() *geom.Mat4 {
	x, y, z := c.CamAxes()
	m := CoordTransf(x, y, z)
	e := c.Eye
	e.Neg()
	t := TranslTransf(&e)
	m.Mul(t)
	return m
}

// ProjTransf returns a new matrix that does the perspective transformation by
// projecting on the near plane of the camera. The z coordinate becomes -c.Near
// and x, y are multiplied by -c.Near/z. -z is factored out as the homogeneous
// part.
func (c *Camera) ProjTransf() *geom.Mat4 {
	n := c.Near
	return &geom.Mat4{
		n, 0, 0, 0,
		0, n, 0, 0,
		0, 0, n, 0,
		0, 0, -1, 0,
	}
}

// UpdateAr update the aspect ratio from the current window size.
func (c *Camera) UpdateAr(w *window.Window) {
	c.Ar = float64(w.Width()) / float64(w.Height())
}

// Frustum is the shape formed by the camera that determines what objects are
// visible and how they are perspectively projected. It is formed by two
// perpendicular rectangles with centers on a line. The near rectangle is on
// the camera's near plane and corresponds to the projection screen. The far
// rectangle is on the far plane and determines how far the camera can see.
type Frustum struct {

	// Nwidth is the width of the near rectangle
	Nwidth float64

	// Nheight is the height of the near rectangle
	Nheight float64

	// Fwidth is the width of the far retangle
	Fwidth float64

	// Fheight ist the height of the far rectangle
	Fheight float64
}

// Frustum returns the camera's frustum.
func (c *Camera) Frustum() *Frustum {
	s := math.Sin(c.Fov)
	n := c.Near
	f := c.Far
	ar := c.Ar
	nw := 2 * s * n
	nh := nw / ar
	fw := 2 * s * f
	fh := fw / ar
	return &Frustum{nw, nh, fw, fh}
}

// ScreenTransf returns a new matrix that transforms vectors after projection
// to screen coordinates. The upper left corner of the near rectangle will be
// (0,0) and the bottom right will be (w,h). If the aspect ratio does not match
// the camera the image will be distorted.
func ScreenTransf(f *Frustum, w, h int) *geom.Mat4 {
	wf := float64(w)
	hf := float64(h)
	t := TranslTransf(&geom.Vec3{f.Nwidth / 2, -f.Nheight / 2, 0})
	m := geom.Mat4{
		wf / f.Nwidth, 0, 0, 0,
		0, -hf / f.Nheight, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	m.Mul(t)
	return &m
}

// PerspTransf returns a new matrix that transforms vectors from world
// coordinates to screen coordinates.
//
// The matrix is constructed by multiplying the following transformation
// matrices (last is applied first in transformation):
//   - screen transformation
//   - perspective transformation
//   - camera transformation
func (c *Camera) PerspTransf(w, h int) *geom.Mat4 {
	f := c.Frustum()
	m := ScreenTransf(f, w, h)
	m.Mul(c.ProjTransf())
	m.Mul(c.CamTransf())
	return m
}

// PerspTransfWin is like PerspTransf() but takes the width/height from the
// current window state.
func (c *Camera) PerspTransfWin(w *window.Window) *geom.Mat4 {
	return c.PerspTransf(w.Width(), w.Height())
}
