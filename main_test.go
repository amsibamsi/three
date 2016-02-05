package main

import (
	"math/rand"
	"testing"
)

func TestEqualVec(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{1, 2, 3}
	v3 := Vector{1, 2, 4}
	if !v1.EqualVec(&v2) {
		t.Fail()
	}
	if v2.EqualVec(&v3) {
		t.Fail()
	}
}

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

func TestTranslate(t *testing.T) {
	v := Vector{2, 7, 6}
	w := Vector{-1, 8, 1.1}
	v.Translate(&w)
	if v[0] != 1 || v[1] != 15 || v[2] != 7.1 {
		t.Fail()
	}
}

func TestMag(t *testing.T) {
	if (&Vector{0, -3, 4}).Mag() != 5 {
		t.Fail()
	}
}

func TestDot(t *testing.T) {
	if (&Vector{2, -1, 3}).Dot(&Vector{4, 1, 1}) != 10 {
		t.Fail()
	}
}

func TestNorm(t *testing.T) {
	if !(&Vector{1, 0, 0}).EqualVec(&Vector{1, 0, 0}) {
		t.Fail()
	}
}

func TestCross(t *testing.T) {
	v1 := Vector{2, 3, 4}
	v2 := Vector{5, 6, 7}
	v3 := v1.Cross(&v2)
	v4 := Vector{-3, 6, -3}
	if !v3.EqualVec(&v4) {
		t.Fail()
	}
}

func TestEqualMat(t *testing.T) {
	if !(&Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}).EqualMat(&Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fail()
	}
}

func TestMulMat(t *testing.T) {
	m1 := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	m2 := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m1.MulMat(&m2)
	m3 := Matrix{18, 21, 24, 54, 66, 78, 90, 111, 132}
	if !m1.EqualMat(&m3) {
		t.Fail()
	}
}

func TestMulMatNew(t *testing.T) {
	m1 := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	m2 := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m3 := m1.MulMatNew(&m2)
	m4 := Matrix{18, 21, 24, 54, 66, 78, 90, 111, 132}
	if !m3.EqualMat(&m4) {
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

func BenchmarkTranslate(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	v := RandVec(r)
	w := RandVec(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Translate(w)
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
