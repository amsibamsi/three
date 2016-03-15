// Generate random matrix with Wolfram Alpha:
// RandomInteger[{-10,10},{4,4}]
package graphics

import (
	"math"
	"math/rand"
	"testing"
)

func TestNewVec(t *testing.T) {
	v := *NewVec(1, 2, 3)
	r := Vec4{1, 2, 3, 1}
	if v != r {
		t.Errorf("expected '%v' but got '%v'", r, v)
	}
}

func TestTranslMat(t *testing.T) {
	m := *TranslMat(1, 2, 3)
	r := Mat4{
		1, 0, 0, 1,
		0, 1, 0, 2,
		0, 0, 1, 3,
		0, 0, 0, 1,
	}
	if m != r {
		t.Errorf("expected '%v' but got '%v'", r, m)
	}
}

func TestRxMat(t *testing.T) {
	m := *RxMat(0)
	r := Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	if m != r {
		t.Errorf("expected '%v' but got '%v'", r, m)
	}
}

func TestRyMat(t *testing.T) {
	m := *RyMat(0)
	r := Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	if m != r {
		t.Errorf("expected '%v' but got '%v'", r, m)
	}
}

func TestMul(t *testing.T) {
	m := Mat4{0, 3, 0, 1, 6, 3, 5, 3, 7, 4, 8, 7, 3, 6, 0, 3}
	n := Mat4{9, 0, 4, 10, 4, 7, 0, 5, 6, 5, 8, 7, 9, 10, 7, 10}
	r := Mat4{21, 31, 7, 25, 123, 76, 85, 140, 190, 138, 141, 216, 78, 72, 33, 90}
	m.Mul(&n)
	if m != r {
		t.Errorf("expected '%v' but got '%v'", r, m)
	}
}

func TestProj(t *testing.T) {
	m := Mat4{1, 3, 2, 2, 9, 10, 1, 9, 0, 4, 5, 1, 6, 8, 5, 8}
	v := Vec4{10, 7, 0, 8}
	r := Vec4{47, 232, 36, 180}
	p := *m.Proj(&v)
	if p != r {
		t.Errorf("expected '%v' but got '%v'", r, p)
	}
}

func TestTransl(t *testing.T) {
	v := Vec4{2, 1, 3, 1}
	m := TranslMat(7, 1, -3)
	r := Vec4{9, 2, 0, 1}
	w := *m.Proj(&v)
	if w != r {
		t.Errorf("expected '%v' but got '%v'", r, w)
	}
}

func TestRx(t *testing.T) {
	v := Vec4{1, 1, 0, 1}
	m := RxMat(math.Pi)
	r := Vec4{1, -1, 0, 1}
	w := *m.Proj(&v)
	if w != r {
		t.Errorf("expected '%v' but got '%v'", r, w)
	}
}

func TestRy(t *testing.T) {
	v := Vec4{1, 1, 0, 1}
	m := RyMat(math.Pi / 2)
	r := Vec4{0, 1, -1, 1}
	w := *m.Proj(&v)
	if w != r {
		t.Errorf("expected '%v' but got '%v'", r, w)
	}
}

func BenchmarkMul(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	m := RandMat(r)
	n := RandMat(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Mul(n)
	}
}
