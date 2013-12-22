package eculid

import "math"
import "github.com/mortdeus/mathgl"

type Point struct{ x, y float32 }

func (p *Point) SetPoint(x, y float32) { p.x, p.y = x, y }
func (p Point) Equals(p2 Point) bool {
	return mathgl.FloatEqual32(p.x, p2.x) && mathgl.FloatEqual32(p.y, p2.y)
}
func (p Point) FuzzyEquals(p2 Point, v float32) bool {
	if p.x-v <= p2.x && p2.x <= p.x+v {
		if p.y-v <= p2.y && p2.y <= p.y+v {
			return true
		}
	}
	return false
}
func (p Point) Angle(p2 Point) float32 {
	a, b := p.Normalize(), p2.Normalize()
	ang := float32(math.Atan2(float64(a.Cross(b)), float64(a.Dot(b))))
	if mathgl.FalmostEqual32(ang, 0) && mathgl.FloatEqual32(ang, 0) {
		return ang
	}
	return 0

}
func (p Point) RotateByAngle(pivot Point, angle float32) Point {
	piv := mathgl.Vec2f{pivot.x, pivot.y}
	v := [2]float32(piv.Add((mathgl.Vec2f{p.x, p.y}).Sub(piv)))
	return (Point{v[0], v[1]}).Rotate(ForAngle(angle))

}
func is1DimensionSegOverlap(a, b, c, d float32) (stat bool, s, e float32) {
	abMin, abMax := mathgl.Fmin32(a, b), mathgl.Fmax32(a, b)
	cdMin, cdMax := mathgl.Fmin32(c, d), mathgl.Fmax32(c, d)
	stat = true
	switch {
	case abMax < cdMin || cdMax < abMin:
		stat = false
	case abMin >= cdMin && abMin <= cdMax:
		if cdMax < abMax {
			s, e = abMax, cdMax
		} else {
			s, e = abMin, abMax
		}
	case abMax >= cdMin && abMax <= cdMax:
		s, e = cdMin, abMax
	default:
		s, e = cdMin, cdMax
	}
	return
}
func IsLineIntersect(a, b, c, d Point) (stat bool, s, t float32) {
	if (a.x == b.x && a.y == b.y) || (c.x == d.x && c.y == d.y) {
		stat = false
	}
	if denom := cross2Vect(a, b, c, d); denom != 0 {
		stat = true
		s = cross2Vect(c, d, c, a) / denom
		t = cross2Vect(a, b, c, a) / denom
	}
	return
}

func check(f float32) bool { return f == 0 }
func IsLineParallel(a, b, c, d Point) bool {

	switch {
	case (a.x == b.x && a.y == b.y) || c.x == d.x && c.y == d.y:
		return false
	case !check(cross2Vect(a, b, c, d)):
		return false
	case check(cross2Vect(c, d, c, a)) || check(cross2Vect(a, b, c, a)):
		return false
	default:
		return true
	}

}
func IsLineOverlap(a, b, c, d Point) bool {
	switch {
	case (a.x == b.x && a.y == b.y) || c.x == d.x && c.y == d.y:
		return false
	case !check(cross2Vect(a, b, c, d)):
		return false
	case check(cross2Vect(c, d, c, a)) || check(cross2Vect(a, b, c, a)):
		return true
	default:
		return false
	}
}
func IsSegmentOverlap(a, b, c, d Point) (stat bool, s, e Point) {
	var stata, statb bool
	if !IsLineOverlap(a, b, c, d) {
		return
	}
	stata, s.x, e.x = is1DimensionSegOverlap(a.x, b.x, c.x, d.x)
	statb, s.y, e.y = is1DimensionSegOverlap(a.y, b.y, c.y, d.y)
	stat = stata && statb
	return
}
func IsSegmentIntersect(a, b, c, d Point) bool {
	stat, s, t := IsLineIntersect(a, b, c, d)
	if stat && (s >= 0 && s <= 1 && s >= 0 && t <= 1) {
		return true
	}
	return false
}
func IntersectPoint(a, b, c, d Point) Point {
	if stat, s, _ := IsLineIntersect(a, b, c, d); stat {
		return Point{a.x + s*(b.x-a.x), a.y + s*(b.y-a.y)}
	}
	return Point{0, 0}

}

func (p Point) Len() float32 { return mathgl.Fsqrt32(mathgl.Fsqr32(p.x) + mathgl.Fsqr32(p.y)) }

func (p Point) Dot(p2 Point) float32    { return p.x*p2.x + p.y*p2.y }
func (p Point) Cross(p2 Point) float32  { return p.x*p2.x - p.y*p2.y }
func (p Point) Perp() Point             { return Point{-p.y, p.x} }
func (p Point) RPerp() Point            { return Point{p.y, -p.x} }
func (p Point) Rotate(p2 Point) Point   { return Point{p.x*p2.x - p.y*p2.y, p.x*p2.y + p.y*p2.x} }
func (p Point) Unrotate(p2 Point) Point { return Point{p.x*p2.x + p.y*p2.y, p.y*p2.x - p.x*p2.y} }
func (p Point) Midpoint(p2 Point) Point { return Point{(p.x + p2.x) / 2, (p.y + p2.y) / 2} }

func (p Point) CompOp(f func(float32) float32) Point { return Point{f(p.x), f(p.y)} }

func (p Point) ClampPoint(min, max Point) Point {
	return Point{
		mathgl.Clampf(p.x, min.x, max.x),
		mathgl.Clampf(p.y, min.y, max.y)}
}
func (p Point) Project(p2 Point) Point {
	return func(f float32) Point { return Point{p2.x * f, p2.y * f} }(p.Dot(p2) / p2.Dot(p2))
}

func (p Point) Normalize() Point {
	if l := p.Len(); l == 0 {
		return Point{1, 0}
	} else {
		return Point{p.x / l, p.y / l}
	}
}
func (p *Point) Lerp(p2 Point, alpha float32) Point {
	return func(p3 Point) Point {
		return Point{
			(p3.x + p2.x*alpha),
			(p3.y + p2.y*alpha),
		}
	}(p.CompOp(func(i float32) float32 { return (i * (1 - alpha)) }))
}

func ForAngle(a float32) Point {
	return Point{mathgl.Fcos32(a), mathgl.Fsin32(a)}
}

func cross2Vect(a, b, c, d Point) float32 {
	return (d.y-c.y)*(b.x-a.x) - (d.x-c.x)*(b.y-a.y)
}

type Size struct{ w, h float32 }

func (sz *Size) SetSize(w, h float32) { sz.w, sz.h = w, h }
func (sz Size) Equals(sz2 Size) bool {
	return mathgl.FloatEqual32(sz.w, sz2.w) && mathgl.FloatEqual32(sz.h, sz2.h)
}

type Rect struct {
	Point
	Size
}

func (r *Rect) SetRect(x, y, w, h float32) { r.x, r.y, r.w, r.h = x, y, w, h }

func (r Rect) Equals(r2 Rect) bool { return (r.Point.Equals(r2.Point) && r.Size.Equals(r2.Size)) }

func (r Rect) MinX() float32 { return r.x }
func (r Rect) MidX() float32 { return r.x + r.w/2 }
func (r Rect) MaxX() float32 { return r.x + r.w }
func (r Rect) MinY() float32 { return r.y }
func (r Rect) MidY() float32 { return r.y + r.h/2 }
func (r Rect) MaxY() float32 { return r.y + r.h }

func (r Rect) ContainsPoint(p Point) bool {
	return p.x >= r.MinX() && p.x <= r.MaxX() && p.y >= r.MinY() && p.y <= r.MaxY()
}
func (r Rect) IntersectsRect(r2 Rect) bool {
	return !(r.MaxX() < r2.MinX() ||
		r2.MaxX() < r.MinX() ||
		r.MaxY() < r2.MinY() ||
		r2.MaxY() < r.MinY())
}
func swapf32(x, y *float32) {
	if *x < *y {
		tmp := x
		x = y
		y = tmp
	}
}
func (r Rect) UnionWithRect(r2 Rect) Rect {
	L, R, T, B := r.x, (r.x + r.w), (r.y + r.h), (r.y)
	L2, R2, T2, B2 := r2.x, (r2.x + r2.w), (r2.y + r2.h), (r2.y)

	swapf32(&R, &L)
	swapf32(&T, &B)
	swapf32(&R2, &L2)
	swapf32(&T2, &B2)

	L3, R3 := mathgl.Fmin32(L, L2), mathgl.Fmax32(R, R2)
	T3, B3 := mathgl.Fmax32(T, T2), mathgl.Fmin32(B, B2)
	return Rect{Point{L3, B3}, Size{R3 - L3, T3 - B3}}
}
