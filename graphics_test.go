// Generate random matrix with Wolfram Alpha:
// RandomInteger[{-10,10},{4,4}]
package graphics

import (
	"math"
	"math/rand"
	"reflect"
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
	v3 := *Cross(&v1, &v2)
	r := Vec3{-3, 6, -3}
	if v3 != r {
		t.Errorf("expected '%v' but got '%v'", r, v3)
	}
}

func TestNewVec4(t *testing.T) {
	v := *NewVec4(1, 2, 3)
	r := Vec4{1, 2, 3, 1}
	if v != r {
		t.Errorf("expected '%v' but got '%v'", r, v)
	}
}

func TestRandMat(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	m := *RandMat(r)
	tm := reflect.TypeOf(m)
	n := Mat4{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	tn := reflect.TypeOf(n)
	if tm != tn {
		t.Errorf("expected '%v' but got '%v'", tn, tm)
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

func TestNewDefCam(t *testing.T) {
	c := *NewDefCam()
	tc := reflect.TypeOf(c)
	r := Camera{
		Eye:  Vec3{0, 0, 0},
		At:   Vec3{0, 0, -1},
		Up:   Vec3{0, 1, 0},
		Near: 1.0,
		Far:  100.0,
		Fov:  math.Pi / 2,
		Ar:   1.0,
	}
	tr := reflect.TypeOf(r)
	if tc != tr {
		t.Errorf("expected '%v' but got '%v'", tr, tc)
	}
}

func TestCamAxes(t *testing.T) {
	c := Camera{
		Eye: Vec3{0, 0, 0},
		At:  Vec3{-3, 0, 0},
		Up:  Vec3{0, 0, -2},
	}
	x, y, z := c.CamAxes()
	xr := Vec3{0, -1, 0}
	yr := Vec3{0, 0, -1}
	zr := Vec3{1, 0, 0}
	if *x != xr {
		t.Errorf("expected x to be '%v' but got '%v'", xr, *x)
	}
	if *y != yr {
		t.Errorf("expected y to be '%v' but got '%v'", yr, *y)
	}
	if *z != zr {
		t.Errorf("expected z to be '%v' but got '%v'", zr, *z)
	}
}

func TestCoordTransf(t *testing.T) {
	x := Vec3{1, 0, 0}
	y := Vec3{0, 1, 0}
	z := Vec3{0, 0, 1}
	m := CoordTransf(&x, &y, &z)
	r := Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	if *m != r {
		t.Errorf("expected '%v' but got '%v'", r, *m)
	}
}

func TestCamTransf(t *testing.T) {
	c := Camera{
		Eye: Vec3{1, 1, 1},
		At:  Vec3{1, 1, 0},
		Up:  Vec3{0, 1, 0},
	}
	m := c.CamTransf()
	r := Mat4{
		1, 0, 0, -1,
		0, 1, 0, -1,
		0, 0, 1, -1,
		0, 0, 0, 1,
	}
	if *m != r {
		t.Errorf("expected '%v' but got '%v'", r, *m)
	}
}

func TestPerspTransf(t *testing.T) {
	c := Camera{
		Near: 3,
	}
	m := c.PerspTransf()
	r := Mat4{
		3, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, 3, 0,
		0, 0, -1, 0,
	}
	if *m != r {
		t.Errorf("expected '%v' but got '%v'", r, *m)
	}
}

func TestFrustum(t *testing.T) {
	c := Camera{
		Near: 2,
		Far:  12,
		Fov:  math.Pi / 2,
		Ar:   2,
	}
	f := c.Frustum()
	r := Frustum{
		Nwidth:  4,
		Nheight: 2,
		Fwidth:  24,
		Fheight: 12,
	}
	if *f != r {
		t.Errorf("expected '%v' but got '%v'", r, *f)
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
