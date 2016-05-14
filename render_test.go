package three

import (
	"math"
	"reflect"
	"testing"
)

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

func TestProjTransf(t *testing.T) {
	c := Camera{
		Near: 3,
	}
	m := c.ProjTransf()
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

func TestScreenTransf1(t *testing.T) {
	f := Frustum{
		Nwidth:  10,
		Nheight: 10,
	}
	m := *ScreenTransf(&f, 100, 100)
	v := Vec4{-10, 10, 2, 2}
	w := *m.Transf(&v)
	r := Vec4{0, 0, 2, 2}
	if w != r {
		t.Errorf("expected '%v' but got '%v'", r, w)
	}
}

func TestScreenTransf2(t *testing.T) {
	f := Frustum{
		Nwidth:  10,
		Nheight: 10,
	}
	m := *ScreenTransf(&f, 100, 100)
	v := Vec4{-5, 5, 2, 2}
	w := *m.Transf(&v)
	r := Vec4{50, 50, 2, 2}
	if w != r {
		t.Errorf("expected '%v' but got '%v'", r, w)
	}
}

func TestPerspTransf(t *testing.T) {
	c := NewDefCam()
	m := c.PerspTransf(100, 100)
	v := &Vec4{2, 1, -2, 1}
	w := m.Transf(v)
	w.Norm()
	r := &Vec4{100, 25, -1, 1}
	if *w != *r {
		t.Errorf("expected '%v' but got '%v'", r, w)
	}
}
