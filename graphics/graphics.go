// Package graphics renders 3D graphics.
package graphics

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math"
	"math/rand"
)

// Abs returns the absolute value of an integer.
func Abs(i int) int {
	m := i >> 31
	return (m ^ i) - m
}

// Round returns the rounded integer for a float.
func Round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// Vec3 is a vector in 3D space with cartesian coordinates. Holds 3 components:
// x, y and z in this order.
type Vec3 [3]float64

// Norm normalizes a vector to length 1 keeping its direction.
func (v *Vec3) Norm() {
	abs := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	v[0] /= abs
	v[1] /= abs
	v[2] /= abs
}

// Neg negates the vector's components.
func (v *Vec3) Neg() {
	v[0] = -v[0]
	v[1] = -v[1]
	v[2] = -v[2]
}

// Sub subtracts another vector.
func (v *Vec3) Sub(w *Vec3) {
	v[0] -= w[0]
	v[1] -= w[1]
	v[2] -= w[2]
}

// Cross returns a new vector that is the cross product of the two vectors.
func Cross(v, w *Vec3) *Vec3 {
	return &Vec3{
		v[1]*w[2] - v[2]*w[1],
		v[2]*w[0] - v[0]*w[2],
		v[0]*w[1] - v[1]*w[0],
	}
}

// Vec4 is a vector in 3D space with homogeneous coordinates. Holds 4
// components: x, y, z and w in this order.
type Vec4 [4]float64

// Norm normalizes a homogeneous vector by dividing x, y and z by w so that w
// will be 1.
func (v *Vec4) Norm() {
	v[0] /= v[3]
	v[1] /= v[3]
	v[2] /= v[3]
	v[3] = 1.0
}

// NewVec4 returns a new vector with homogeneous coordinates corresponding to
// the given cartesian coordinates (w will be 1).
func NewVec4(x, y, z float64) *Vec4 {
	return &Vec4{x, y, z, 1}
}

// Mat4 is a matrix with homogeneous coordinates used to transform homogeneous
// vectors.  Holds 16 components, the 4 first elements make up the first row
// from left to right, and so on.
type Mat4 [16]float64

// ZeroMat returns a new matrix with all values set to zero.
func ZeroMat() *Mat4 {
	return &Mat4{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
}

// RandMat returns a new matrix random values.
func RandMat(r *rand.Rand) *Mat4 {
	m := Mat4{}
	for i := 0; i < len(m); i++ {
		m[i] = r.Float64()
	}
	return &m
}

// TranslTransf returns a new translation matrix that translates vectors by the
// argument vector.
func TranslTransf(v *Vec3) *Mat4 {
	return &Mat4{
		1, 0, 0, v[0],
		0, 1, 0, v[1],
		0, 0, 1, v[2],
		0, 0, 0, 1,
	}
}

// Mul multiplies the matrix with another one, modifying the former one.
func (m *Mat4) Mul(n *Mat4) {
	t := ZeroMat()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				t[i*4+j] += m[i*4+k] * n[j+k*4]
			}
		}
	}
	*m = *t
}

// Transf returns a new transformed vector by multiplying the matrix with the
// given vector.
func (m *Mat4) Transf(v *Vec4) *Vec4 {
	p := Vec4{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			p[i] += m[i*4+j] * v[j]
		}
	}
	return &p
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
func CoordTransf(x, y, z *Vec3) *Mat4 {
	return &Mat4{
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
type Camera struct {

	// Eye is the position of the eye to look from. For each object to project a
	// virtual line to the eye is drawn.
	Eye Vec3

	// At is the direction to look at from the eye.
	At Vec3

	// Up determines the orientation of the view. Up not being perpendicular to
	// At results in the same orientation as if Up was first projected to the
	// normal plane of At through Eye.
	Up Vec3

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
		Eye:  Vec3{0, 0, 0},
		At:   Vec3{0, 0, -1},
		Up:   Vec3{0, 1, 0},
		Near: 1.0,
		Far:  100.0,
		Fov:  math.Pi / 2,
		Ar:   1.0,
	}
}

// CamAxes returns the 3 axes that make up the orthonormal basis of the
// camera's right-handed coordinate system.
func (c *Camera) CamAxes() (*Vec3, *Vec3, *Vec3) {
	z := c.Eye
	z.Sub(&c.At)
	z.Norm()
	x := Cross(&c.Up, &z)
	x.Norm()
	// Recompute up, c.Up might not be perpendicular or normalized
	y := Cross(&z, x)
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
func (c *Camera) CamTransf() *Mat4 {
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
func (c *Camera) ProjTransf() *Mat4 {
	n := c.Near
	return &Mat4{
		n, 0, 0, 0,
		0, n, 0, 0,
		0, 0, n, 0,
		0, 0, -1, 0,
	}
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

// Screen is the rectangular area to render an image on. The origin is in the
// upper left (0,0) and it extends to the right bottom.
type Screen struct {

	// Width of the screen in pixels
	Width int

	// Height of the screen in pixels
	Height int
}

// ScreenTransf returns a new matrix that transforms vectors after projection
// to screen coordinates. The upper left corner of the near rectangle will be
// (0,0) and the bottom right will be (width,height). If the aspect ratio does
// not match the camera the image will be distorted. It needs the frustum of
// the camera and a screen as arguments.
func ScreenTransf(f *Frustum, s *Screen) *Mat4 {
	w := float64(s.Width)
	h := float64(s.Height)
	t := TranslTransf(&Vec3{f.Nwidth / 2, -f.Nheight / 2, 0})
	m := Mat4{
		w / f.Nwidth, 0, 0, 0,
		0, -h / f.Nheight, 0, 0,
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
func (c *Camera) PerspTransf(s *Screen) *Mat4 {
	f := c.Frustum()
	m := ScreenTransf(f, s)
	m.Mul(c.ProjTransf())
	m.Mul(c.CamTransf())
	return m
}

// Image is a simple 2D image. It wraps an RGBA image from the standard image
// package.
type Image struct {
	Rgba image.RGBA
}

// NewImage returns a new image with the given screen width and height and
// black background.
func NewImage(s *Screen) *Image {
	rect := image.Rect(0, 0, s.Width, s.Height)
	rgba := image.NewRGBA(rect)
	bg := image.Uniform{color.Black}
	draw.Draw(rgba, rgba.Bounds(), &bg, image.Point{}, draw.Src)
	return &Image{*rgba}
}

// DrawDot draws a clearly visible dot (more than 1 pixel) at (x,y) with the
// given color.
func (img *Image) DrawDot(x, y int, c color.Color) {
	r := img.Rgba
	ind := []int{
		-2, 0,
		-1, -1,
		-1, 0,
		-1, 1,
		0, -1,
		0, 0,
		0, 1,
		1, -1,
		1, 0,
		1, 1,
		2, 0,
	}
	for i := 0; i < len(ind); i += 2 {
		r.Set(x+ind[i], y+ind[i+1], c)
	}
}

// DrawLine draws a 1 pixel thick line between the (x1,y1) and (x2,y2) with the
// given color.
func (img *Image) DrawLine(x1, y1, x2, y2 int, c color.Color) {
	r := img.Rgba
	// Always draw from left to right (x1 <= x2)
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}
	dx := x2 - x1
	dy := y2 - y1
	var steps int
	if Abs(dx) > Abs(dy) {
		steps = Abs(dx)
	} else {
		steps = Abs(dy)
	}
	xinc := float64(dx) / float64(steps)
	yinc := float64(dy) / float64(steps)
	x := float64(x1)
	y := float64(y1)
	for s := 0; s <= steps; s++ {
		r.Set(Round(x), Round(y), c)
		x += xinc
		y += yinc
	}
}

// WritePng stores the image in PNG format to the given writer and returns the
// error from png.Encode() if any.
func (img *Image) WritePng(w io.Writer) error {
	return png.Encode(w, &img.Rgba)
}
