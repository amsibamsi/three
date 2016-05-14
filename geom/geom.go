// Package geom contains some geometric primitives and realted transformation
// and drawing.
package geom

import (
	"github.com/amsibamsi/three"
	"github.com/amsibamsi/three/window"
)

// A triangle in 2D space, contains 3 2D vectors.
type Tri2 [3]three.Vec2

// NewTri2 returns a new 2D triangle with points (x1,y1), (x2,y2) and (x3,y3).
func NewTri2(x1, y1, x2, y2, x3, y3 int) *Tri2 {
	return &Tri2{
		three.Vec2{x1, y1},
		three.Vec2{x2, y2},
		three.Vec2{x3, y3},
	}
}

// Draw draws the triangle on a window as wireframe.
func (t *Tri2) Draw(w *window.Window) {
	w.Dot(&t[0], 255, 0, 0)
	w.Dot(&t[1], 255, 0, 0)
	w.Dot(&t[2], 255, 0, 0)
	w.Line(&t[0], &t[1], 255, 0, 0)
	w.Line(&t[1], &t[2], 255, 0, 0)
	w.Line(&t[2], &t[0], 255, 0, 0)
}

// A triangle in 3D space with homogeneous coordinates
type Tri4 [3]three.Vec4

// NewTri4 returns a new triangle with points (x1,y1,z1), (x2,y2,z2) and
// (x3,y3,z3).
func NewTri4(x1, y1, z1, x2, y2, z2, x3, y3, z3 float64) *Tri4 {
	return &Tri4{
		*three.NewVec4(x1, y1, z1),
		*three.NewVec4(x2, y2, z2),
		*three.NewVec4(x3, y3, z3),
	}
}

// Transf transforms a 3D triangle with the given transformation matrix and
// returns a 2D triangle.
func (t *Tri4) Transf(m *three.Mat4) *Tri2 {
	p1 := m.Transf(&t[0])
	x1 := three.Round(p1[0] / p1[3])
	y1 := three.Round(p1[1] / p1[3])
	p2 := m.Transf(&t[1])
	x2 := three.Round(p2[0] / p2[3])
	y2 := three.Round(p2[1] / p2[3])
	p3 := m.Transf(&t[2])
	x3 := three.Round(p3[0] / p3[3])
	y3 := three.Round(p3[1] / p3[3])
	return NewTri2(x1, y1, x2, y2, x3, y3)
}
