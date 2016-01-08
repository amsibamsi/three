package main

import (
	"math/rand"
	"testing"
)

func TestMulVec(t *testing.T) {
	m := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	v := Vector{0, 1, 2}
	v.MulVec(&m)
	if v[0] != 5 || v[1] != 14 || v[2] != 23 {
		t.Fail()
	}
}

func TestMulVecNew(t *testing.T) {
	m := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	v := Vector{0, 1, 2}
	w := v.MulVecNew(&m)
	if w[0] != 5 || w[1] != 14 || w[2] != 23 {
		t.Fail()
	}
}

func TestMulMat(t *testing.T) {
	m := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	n := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m.MulMat(&n)
	if m[0] != 18 ||
		m[1] != 21 ||
		m[2] != 24 ||
		m[3] != 54 ||
		m[4] != 66 ||
		m[5] != 78 ||
		m[6] != 90 ||
		m[7] != 111 ||
		m[8] != 132 {
		t.Fail()
	}
}

func TestMulMatNew(t *testing.T) {
	m := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	n := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	o := m.MulMatNew(&n)
	if o[0] != 18 ||
		o[1] != 21 ||
		o[2] != 24 ||
		o[3] != 54 ||
		o[4] != 66 ||
		o[5] != 78 ||
		o[6] != 90 ||
		o[7] != 111 ||
		o[8] != 132 {
		t.Fail()
	}
}

func BenchmarkMulVec(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	v := RandVec(r)
	m := RandMat(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MulVec(m)
	}
}

func BenchmarkMulVecNew(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	v := RandVec(r)
	m := RandMat(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MulVecNew(m)
	}
}

func BenchmarkMulMat(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	m := RandMat(r)
	n := RandMat(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.MulMat(n)
	}
}

func BenchmarkMulMatNew(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	m := RandMat(r)
	n := RandMat(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.MulMatNew(n)
	}
}
