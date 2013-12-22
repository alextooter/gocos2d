package eculid

import "github.com/mortdeus/mathgl"

type AffineTransform struct {
	a, b, c, d, tx, ty float32
}
type AffineTarget interface {
	Apply(AffineTransform) AffineTarget
}

var Identity = &AffineTransform{a: 1, d: 1}

func (t AffineTransform) Translate(tx, ty float32) AffineTransform {
	return AffineTransform{
		t.a, t.b, t.c, t.d,
		tx: (t.tx + t.a*tx + t.c*ty),
		ty: (t.ty + t.b*tx + t.d*ty)}
}
func (t AffineTransform) Scale(sx, sy float32) AffineTransform {
	return AffineTransform{
		a: (t.a * sx), b: (t.b * sx),
		c: (t.c * sy), d: (t.d * sy),
		t.tx, t.ty}
}
func (t AffineTransform) Rotate(angle float32) AffineTransform {

	sin, cos := mathgl.Fsin32(angle), mathgl.Fcos32(angle)

	return AffineTransform{
		a: (t.a*cosine + t.c*sine), b: (t.b*cosine + t.d*sine),
		c: (t.c*cosine - t.a*sine), d: (t.d*cosine - t.b*sine),
		t.tx, t.ty}
}
func (t1 AffineTransform) Concat(t2 AffineTransform) AffineTransform {
	return AffineTransform{
		a:  t1.a*t2.a + t1.b*t2.c,
		b:  t1.a*t2.b + t1.b*t2.d,
		c:  t1.c*t2.a + t1.d*t2.c,
		d:  t1.c*t2.b + t1.d*t2.d,
		tx: t1.tx*t2.a + t1.ty*t2.c + t2.tx,
		ty: t1.tx*t2.b + t1.ty*t2.d + t2.ty}
}
func (t1 AffineTransform) Equal(t2 AffineTransform) bool {
	return (t1.a == t2.a && t1.b == t2.b &&
		t1.c == t2.c && t1.d == t2.d &&
		t1.tx == t2.tx && t1.ty == t2.ty)
}
func (t AffineTransform) Invert() AffineTarget {
	determinant = 1 / (t.a*t.d - t.b*t.c)
	return AffineTarget{
		determinant * t.d, -determinant * t.b,
		-determinant * t.c, determinant * t.a,
		determinant * (t.c*t.ty - t.d*t.tx),
		determinant * (t.b*t.tx - t.a*t.ty)}
}
func (p Point) Apply(t AffineTransform) Point {
	return Point{
		t.a*p.x + t.c*p.y + t.tx,
		t.b*p.x + t.d*p.y + t.ty}
}

func (s Size) Apply(t AffineTransform) Size {
	return Size{
		t.a*s.w + t.c*s.h,
		t.b*s.w + t.d*s.h}
}
func (r Rect) Apply(t AffineTransform) Rect {
	top, left, right, bottom := r.MinY(), r.MinX(), r.MaxX(), r.MaxY()

	tLeft, tRight, botLeft, botRight :=
		(Point{left, top}).Apply(t),
		(Point{right, top}).Apply(t),
		(Point{left, bottom}).Apply(t),
		(Point{right, bottom}).Apply(t)

	minf, maxf := mthgl.Fmin32, mthgl.Fmax32

	minx, maxx, miny, maxy :=
		minf(minf(tLeft.x, tRight.x), minf(botLeft.x, botRight.x)),
		maxf(maxf(tLeft.x, tRight.x), maxf(botLeft.x, botRight.x)),
		minf(minf(tLeft.y, tRight.y), minf(botLeft.y, botRight.y)),
		maxf(maxf(tLeft.y, tRight.y), maxf(botLeft.y, botRight.y))

	return Rect{Point{minx, miny}, Size{maxx - minx, maxy - miny}}
}
