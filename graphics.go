// Render 3D graphics
package graphics

import (
	"math"
	"math/rand"
)

// Vector with homogeneous coordinates
// Holds 4 components: x, y, z and w in this order.
type Vec4 [4]float64

// NewVec returns a new vector with given x/y/z and w set to 1.
func NewVec(x, y, z float64) *Vec4 {
	return &Vec4{x, y, z, 1}
}

// Matrix with homogeneous coordinates
// Holds 16 components, the 4 first elements make up the first row from left to right, and so on.
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

// TranslMat returns a new translation matrix that translates vectors by the given values for x/y/z axes.
func TranslMat(x, y, z float64) *Mat4 {
	return &Mat4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}

// RxMat returns a new rotation matrix that rotates about the x axis by a radians.
func RxMat(a float64) *Mat4 {
	s := math.Sin(a)
	c := math.Cos(a)
	return &Mat4{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

// RyMat returns a new rotation matrix that rotates about the y axis by a radians.
func RyMat(a float64) *Mat4 {
	s := math.Sin(a)
	c := math.Cos(a)
	return &Mat4{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
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

// Transf projects a given vector with the projection matrix.
// Returns a new vector.
func (m *Mat4) Transf(v *Vec4) *Vec4 {
	p := Vec4{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			p[i] += m[i*4+j] * v[j]
		}
	}
	return &p
}
