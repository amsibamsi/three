package threed

import (
	"math"
	"math/rand"
)

type Vector [3]float64

// Dup duplicates the vector and returns a new instance.
func (v *Vector) Dup() *Vector {
	return &Vector{v[0], v[1], v[2]}
}

// RandVec returns a new vector with random values.
func RandVec(r *rand.Rand) *Vector {
	v := Vector{}
	for i := 0; i < 3; i++ {
		v[i] = r.Float64()
	}
	return &v
}

// EqualVec returns whether the vector equals another one.
func (v *Vector) EqualVec(w *Vector) bool {
	return v[0] == w[0] && v[1] == w[1] && v[2] == w[2]
}

// MulVec multiplies the given matrix with the vector and modifies it.
func (v *Vector) MulVec(m *Matrix) {
	x := m[0]*v[0] + m[1]*v[1] + m[2]*v[2]
	y := m[3]*v[0] + m[4]*v[1] + m[5]*v[2]
	z := m[6]*v[0] + m[7]*v[1] + m[8]*v[2]
	v[0] = x
	v[1] = y
	v[2] = z
}

// MulVecNew multiplies a matrix with the vector and returns a new vector.
// MulVecNew is slower than MulVec.
func (v *Vector) MulVecNew(m *Matrix) *Vector {
	return &Vector{
		m[0]*v[0] + m[1]*v[1] + m[2]*v[2],
		m[3]*v[0] + m[4]*v[1] + m[5]*v[2],
		m[6]*v[0] + m[7]*v[1] + m[8]*v[2],
	}
}

// Translate adds another vector.
func (v *Vector) Translate(w *Vector) {
	v[0] += w[0]
	v[1] += w[1]
	v[2] += w[2]
}

// Invert inverts the vector.
func (v *Vector) Invert() {
	v[0] = -v[0]
	v[1] = -v[1]
	v[2] = -v[2]
}

// Magnitude of a vector.
func (v *Vector) Mag() float64 {
	return math.Sqrt(math.Pow(v[0], 2) + math.Pow(v[1], 2) + math.Pow(v[2], 2))
}

// Dot returns the product (scalar) with another vector.
func (v *Vector) Dot(w *Vector) float64 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2]
}

// Norm normalizes the vector, making its length 1.
func (v *Vector) Norm() {
	m := v.Mag()
	v[0] /= m
	v[1] /= m
	v[2] /= m
}

// Cross returns the cross product of two vectors as new vector.
func (v *Vector) Cross(w *Vector) *Vector {
	return &Vector{
		v[1]*w[2] - v[2]*w[1],
		v[2]*w[0] - v[0]*w[2],
		v[0]*w[1] - v[1]*w[0],
	}
}

// Angle returns the angle between two vectors in radians.
func (v *Vector) Angle(w *Vector) float64 {
	return math.Acos(v.Dot(w) / (v.Mag() * w.Mag()))
}

type Matrix [9]float64

// EqualMat returns whether the matrix equals another one.
func (m *Matrix) EqualMat(n *Matrix) bool {
	e := true
	for i := 0; i < 9; i++ {
		if m[i] != n[i] {
			e = false
			break
		}
	}
	return e
}

// RandMat return a new matrix with random values.
func RandMat(r *rand.Rand) *Matrix {
	m := Matrix{}
	for i := 0; i < 9; i++ {
		m[i] = r.Float64()
	}
	return &m
}

// MulMat multiples the matrix with another matrix and modifies the former one.
func (m *Matrix) MulMat(n *Matrix) {
	m0 := m[0]*n[0] + m[1]*n[3] + m[2]*n[6]
	m1 := m[0]*n[1] + m[1]*n[4] + m[2]*n[7]
	m2 := m[0]*n[2] + m[1]*n[5] + m[2]*n[8]
	m3 := m[3]*n[0] + m[4]*n[3] + m[5]*n[6]
	m4 := m[3]*n[1] + m[4]*n[4] + m[5]*n[7]
	m5 := m[3]*n[2] + m[4]*n[5] + m[5]*n[8]
	m6 := m[6]*n[0] + m[7]*n[3] + m[8]*n[6]
	m7 := m[6]*n[1] + m[7]*n[4] + m[8]*n[7]
	m8 := m[6]*n[2] + m[7]*n[5] + m[8]*n[8]
	m[0] = m0
	m[1] = m1
	m[2] = m2
	m[3] = m3
	m[4] = m4
	m[5] = m5
	m[6] = m6
	m[7] = m7
	m[8] = m8
}

// MulMatNew multiples the matrix with another matrix and returns a new matrix.
// MulMatNew is slower than MulMat.
func (m *Matrix) MulMatNew(n *Matrix) *Matrix {
	return &Matrix{
		m[0]*n[0] + m[1]*n[3] + m[2]*n[6],
		m[0]*n[1] + m[1]*n[4] + m[2]*n[7],
		m[0]*n[2] + m[1]*n[5] + m[2]*n[8],
		m[3]*n[0] + m[4]*n[3] + m[5]*n[6],
		m[3]*n[1] + m[4]*n[4] + m[5]*n[7],
		m[3]*n[2] + m[4]*n[5] + m[5]*n[8],
		m[6]*n[0] + m[7]*n[3] + m[8]*n[6],
		m[6]*n[1] + m[7]*n[4] + m[8]*n[7],
		m[6]*n[2] + m[7]*n[5] + m[8]*n[8],
	}
}

type View struct {
	eye  Vector
	look Vector
	up   Vector
	near float64
	far  float64
	fovx float64
	fovy float64
}
