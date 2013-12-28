package gocos2d

import "github.com/mortdeus/mathgl"

type AffineTransform interface {
	Translate(tx, ty float32) AffineTransform
	Scale(sx, sy float32) AffineTransform
	Rotate(angle float32) AffineTransform
	Concat(t2 struct{ A, B, C, D, Tx, Ty float32 }) AffineTransform
	Equal(t2 struct{ A, B, C, D, Tx, Ty float32 }) bool
	Invert() AffineTransform
}
type affineTransform struct {
	A, B, C, D, Tx, Ty float32
}
type AffineTarget interface {
	Apply(affineTransform) AffineTarget
}

var Identity = &affineTransform{A: 1, D: 1}

func (t affineTransform) Translate(tx, ty float32) AffineTransform {
	return affineTransform{
		A: t.A, B: t.B,
		C: t.C, D: t.D,
		Tx: (t.Tx + t.A*tx + t.C*ty),
		Ty: (t.Ty + t.B*tx + t.D*ty)}
}
func (t affineTransform) Scale(sx, sy float32) AffineTransform {
	return affineTransform{
		A: (t.A * sx), B: (t.B * sx),
		C: (t.C * sy), D: (t.D * sy),
		Tx: t.Tx,
		Ty: t.Ty}
}
func (t affineTransform) Rotate(angle float32) AffineTransform {

	sin, cos := mathgl.Fsin32(angle), mathgl.Fcos32(angle)

	return affineTransform{
		A: (t.A*cos + t.C*sin), B: (t.B*cos + t.D*sin),
		C: (t.C*cos - t.A*sin), D: (t.D*cos - t.B*sin),
		Tx: t.Tx,
		Ty: t.Ty}
}
func (t1 affineTransform) Concat(t2 struct{ A, B, C, D, Tx, Ty float32 }) AffineTransform {
	return affineTransform{
		A:  t1.A*t2.A + t1.B*t2.C,
		B:  t1.A*t2.B + t1.B*t2.D,
		C:  t1.C*t2.A + t1.D*t2.C,
		D:  t1.C*t2.B + t1.D*t2.D,
		Tx: t1.Tx*t2.A + t1.Ty*t2.C + t2.Tx,
		Ty: t1.Tx*t2.B + t1.Ty*t2.D + t2.Ty}
}
func (t1 affineTransform) Equal(t2 struct{ A, B, C, D, Tx, Ty float32 }) bool {
	return (t1.A == t2.A && t1.B == t2.B &&
		t1.C == t2.C && t1.D == t2.D &&
		t1.Tx == t2.Tx && t1.Ty == t2.Ty)
}
func (t affineTransform) Invert() AffineTransform {
	determinant := 1 / (t.A*t.D - t.B*t.C)
	return affineTransform{
		determinant * t.D, -determinant * t.B,
		-determinant * t.C, determinant * t.A,
		determinant * (t.C*t.Ty - t.D*t.Tx),
		determinant * (t.B*t.Tx - t.A*t.Ty)}
}
func (p Point) Apply(t struct{ A, B, C, D, Tx, Ty float32 }) Point {
	return Point{
		t.A*p.X + t.C*p.Y + t.Tx,
		t.B*p.X + t.D*p.Y + t.Ty}
}

func (s Size) Apply(t struct{ A, B, C, D, Tx, Ty float32 }) Size {
	return Size{
		t.A*s.W + t.C*s.H,
		t.B*s.W + t.D*s.H}
}
func (r Rect) Apply(t struct{ A, B, C, D, Tx, Ty float32 }) Rect {
	top, left, right, bottom := r.MinY(), r.MinX(), r.MaxX(), r.MaxY()

	tLeft, tRight, botLeft, botRight :=
		(Point{left, top}).Apply(t),
		(Point{right, top}).Apply(t),
		(Point{left, bottom}).Apply(t),
		(Point{right, bottom}).Apply(t)

	minf, maxf := mathgl.Fmin32, mathgl.Fmax32

	minx, maxx, miny, maxy :=
		minf(minf(tLeft.X, tRight.X), minf(botLeft.X, botRight.X)),
		maxf(maxf(tLeft.X, tRight.X), maxf(botLeft.X, botRight.X)),
		minf(minf(tLeft.Y, tRight.Y), minf(botLeft.Y, botRight.Y)),
		maxf(maxf(tLeft.Y, tRight.Y), maxf(botLeft.Y, botRight.Y))

	return Rect{Point{minx, miny}, Size{maxx - minx, maxy - miny}}
}
