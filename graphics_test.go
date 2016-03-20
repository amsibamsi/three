// Generate random matrix with Wolfram Alpha:
// RandomInteger[{-10,10},{4,4}]
package graphics

import (
	"math/rand"
	"testing"
)

func TestNorm(t *testing.T) {
	v := Vec3{10, 0, 0}
	r := Vec3{1, 0, 0}
	v.Norm()
	if v != r {
		t.Errorf("expected '%v' but got '%v'", r, v)
	}
}

func TestNeg(t *testing.T) {
	v := Vec3{1, -2, 0}
	r := Vec3{-1, 2, 0}
	v.Neg()
	if v != r {
		t.Errorf("expected '%v' but got '%v'", r, v)
	}
}

func TestSub(t *testing.T) {
	v := Vec3{50, -2, 7}
	w := Vec3{1, 1, -6}
	r := Vec3{49, -3, 13}
	v.Sub(&w)
	if v != r {
		t.Errorf("expected '%v' but got '%v'", r, v)
	}
}

func TestCross(t *testing.T) {
	v1 := Vec3{2, 3, 4}
	v2 := Vec3{5, 6, 7}
	v3 := *v1.Cross(&v2)
	r := Vec3{-3, 6, -3}
	if v3 != r {
		t.Errorf("expected '%v' but got '%v'", r, v3)
	}
}

func TestTranslTransf(t *testing.T) {
	m := *TranslTransf(&Vec3{1, 2, 3})
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
func TestMul(t *testing.T) {
	m := Mat4{0, 3, 0, 1, 6, 3, 5, 3, 7, 4, 8, 7, 3, 6, 0, 3}
	n := Mat4{9, 0, 4, 10, 4, 7, 0, 5, 6, 5, 8, 7, 9, 10, 7, 10}
	r := Mat4{21, 31, 7, 25, 123, 76, 85, 140, 190, 138, 141, 216, 78, 72, 33, 90}
	m.Mul(&n)
	if m != r {
		t.Errorf("expected '%v' but got '%v'", r, m)
	}
}

func TestTransf(t *testing.T) {
	m := Mat4{1, 3, 2, 2, 9, 10, 1, 9, 0, 4, 5, 1, 6, 8, 5, 8}
	v := Vec4{10, 7, 0, 8}
	r := Vec4{47, 232, 36, 180}
	p := *m.Transf(&v)
	if p != r {
		t.Errorf("expected '%v' but got '%v'", r, p)
	}
}

func TestTransl(t *testing.T) {
	v := Vec4{2, 1, 3, 1}
	m := TranslTransf(&Vec3{7, 1, -3})
	r := Vec4{9, 2, 0, 1}
	w := *m.Transf(&v)
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
