package main

import (
	"testing"
)

func TestMulv(t *testing.T) {
	m := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	v := Vector{0, 1, 2}
	w := m.Mulv(&v)
	if w[0] != 5 || w[1] != 14 || w[2] != 23 {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	m := Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}
	n := Matrix{1, 2, 3, 4, 5, 6, 7, 8, 9}
	o := m.Mul(&n)
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
