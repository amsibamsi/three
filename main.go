package main

import (
	"math/rand"
)

type Vector [3]float64

func RandVec(r *rand.Rand) *Vector {
	v := Vector{}
	for i := 0; i < 3; i++ {
		v[i] = r.Float64()
	}
	return &v
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

type Matrix [9]float64

func RandMat(r *rand.Rand) *Matrix {
	m := Matrix{}
	for i := 0; i < 9; i++ {
		m[i] = r.Float64()
	}
	return &m
}

// MulVecNew multiplies the matrix with a vector and returns a new vector.
// MulVecNew Is slower than MulVec.
func (m *Matrix) MulVecNew(v *Vector) *Vector {
	return &Vector{
		m[0]*v[0] + m[1]*v[1] + m[2]*v[2],
		m[3]*v[0] + m[4]*v[1] + m[5]*v[2],
		m[6]*v[0] + m[7]*v[1] + m[8]*v[2],
	}
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
