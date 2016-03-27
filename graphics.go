// Package graphics renders 3D graphics.
package graphics

import (
	"math"
	"math/rand"
)

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

// CoordTransf returns a new matrix that transforms vectors to a new basis with
// axes specified as x, y, z arguments which must be orthonormal.
//
// Reference is the standard basis with the axes (1,0,0), (0,1,0) and (0,0,1).
// A new basis is formed from the argument axes given in standard coordinates.
// Any vector in the standard basis is interpreted as seen by the new basis
// and its coordinates are transformed to this new basis. This corresponds to
// the inverse transformation of the standard basis into the new basis. We
// see a vector through the view of the new basis, transform both back to the
// standard basis and get the standard coordinates of the vector as seen by
// the new basis.
//
// Transforming from the standard basis into the new basis is just a matter of
// writing the standard coordinates of a vector as linear combination of the
// new axes. E.g. a vector (v0,v1,v2) can be written as v0 * (1,0,0) + v1 *
// (0,1,0) + v2 * (0,0,1) in the standard basis. And in the new basis with
// axes (a0,a1,a2), (b0,b1,b2) and (c0,c1,c2) as v0 * (a0,a1,a2) + v1 *
// (b0,b1,b2) + v2 * (c0,c1,c2). So the x coordinate of the new vector is
// v0*a0 + v1*b0 + v2*c0 and similar for y and z coordinates. Thus the 3 new
// basis vectors make up the columns of the transformation matrix. The final
// matrix is then the inverted transformation, so the transposed matrix where
// the axes of the new bases are the rows, since we want to transform vectors
// as seen by the new basis back to the standard basis to get the standard
// coordinates. Finally we fill up the matrix with 0 an 1 to make it
// homogeneous.
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

// CamAxes returns the 3 axes that make up the orthonormal basis of the cameras
// right-handed coordinate system.
func (c *Camera) CamAxes() (*Vec3, *Vec3, *Vec3) {
	z := c.Eye
	z.Sub(&c.At)
	z.Norm()
	x := Cross(&c.Up, &z)
	x.Norm()
	// Recompute up, c.Up might not be perpendicular or normalized.
	y := Cross(&z, x)
	return x, y, &z
}

// CamTransf returns a new matrix that transforms from world coordinates into
// view coordinates of the camera. It is the inverse transformation of getting
// from the world view (defined as x to the right, y up and looking down -z)
// to the camera view. That means first translating the eye of the camera to
// the origin and then rotating the camera to match the world view.
func (c *Camera) CamTransf() *Mat4 {
	x, y, z := c.CamAxes()
	m := CoordTransf(x, y, z)
	e := c.Eye
	e.Neg()
	t := TranslTransf(&e)
	m.Mul(t)
	return m
}
