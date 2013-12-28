package gocos2d

import "math"
import "github.com/mortdeus/mathgl"

//TODO(mortdeus): Write tests for eculid objects.

type Point struct{ X, Y float32 }

func (p *Point) SetPoint(x, y float32) { p.X, p.Y = x, y }
func (p Point) Equals(p2 Point) bool {
	return mathgl.FloatEqual32(p.X, p2.X) && mathgl.FloatEqual32(p.Y, p2.Y)
}
func (p Point) FuzzyEquals(p2 Point, v float32) bool {
	if p.X-v <= p2.X && p2.X <= p.X+v && p.Y-v <= p2.Y && p2.Y <= p.Y+v {
		return true
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
	piv := mathgl.Vec2f{pivot.X, pivot.Y}
	v := [2]float32(piv.Add((mathgl.Vec2f{p.X, p.Y}).Sub(piv)))
	return (Point{v[0], v[1]}).Rotate(ForAngle(angle))
}

func is1DimensionSegOverlap(a, b, c, d float32, s_e *struct{ S, E float32 }) bool {
	abMin, abMax := mathgl.Fmin32(a, b), mathgl.Fmax32(a, b)
	cdMin, cdMax := mathgl.Fmin32(c, d), mathgl.Fmax32(c, d)
	ltab := [12]struct{ S, E float32 }{
		{abMin, abMax}, {abMin, cdMin}, {abMin, cdMax}, {abMax, abMin},
		{abMax, cdMin}, {abMax, cdMax}, {cdMin, cdMax}, {cdMin, abMin},
		{cdMin, abMax}, {cdMax, cdMin}, {cdMax, abMin}, {cdMax, abMax},
	}
	switch {
	case abMax < cdMin || cdMax < abMin:
		return false
	case abMin >= cdMin && abMin <= cdMax:
		switch cdMax < abMax {
		case true:
			*s_e = ltab[5]
		default:
			*s_e = ltab[0]
		}
	case abMax >= cdMin && abMax >= cdMax:
		*s_e = ltab[8]
	default:
		*s_e = ltab[6]
	}
	return true

}
func IsLineIntersect(a, b, c, d Point, s_t *struct{ S, T float32 }) bool {
	if (a.X == b.X && a.Y == b.Y) || (c.X == d.X && c.Y == d.Y) {
		return false
	}
	denom := cross2Vect(a, b, c, d)
	if denom == 0 {
		return false
	}
	*s_t = struct{ S, T float32 }{cross2Vect(c, d, c, a) / denom, cross2Vect(a, b, c, a) / denom}
	return true
}

func check(f float32) bool { return f == 0 }
func IsLineParallel(a, b, c, d Point) bool {
	switch {
	case (a.X == b.X && a.Y == b.Y) || c.X == d.X && c.Y == d.Y:
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
	case (a.X == b.X && a.Y == b.Y) || c.X == d.X && c.Y == d.Y:
		return false
	case !check(cross2Vect(a, b, c, d)):
		return false
	case check(cross2Vect(c, d, c, a)) || check(cross2Vect(a, b, c, a)):
		return true
	default:
		return false
	}
}
func IsSegmentOverlap(a, b, c, d Point, s_e []Point) bool {
	if !IsLineOverlap(a, b, c, d) {
		return false
	}
	vs := struct{ S, E float32 }{s_e[0].X, s_e[1].X}
	ve := struct{ S, E float32 }{s_e[0].Y, s_e[1].Y}
	defer func() {
		s_e[0].X, s_e[0].Y = vs.S, vs.E
		s_e[1].X, s_e[1].Y = ve.S, vs.E

	}()
	return is1DimensionSegOverlap(a.X, b.X, c.X, d.X, &vs) &&
		is1DimensionSegOverlap(a.Y, b.Y, c.Y, d.Y, &ve)
}
func IsSegmentIntersect(a, b, c, d Point) bool {
	s_t := struct{ S, T float32 }{}
	if IsLineIntersect(a, b, c, d, &s_t) &&
		(s_t.S >= 0 && s_t.S <= 1 && s_t.S >= 0 && s_t.T <= 1) {
		return true
	}
	return false
}
func IntersectPoint(a, b, c, d Point) Point {
	s_t := struct{ S, T float32 }{}
	if IsLineIntersect(a, b, c, d, &s_t) {
		return Point{a.X + s_t.S*(b.X-a.X), a.Y + s_t.S*(b.Y-a.Y)}
	}
	return Point{0, 0}

}

func (p Point) Len() float32 { return mathgl.Fsqrt32(mathgl.Fsqr32(p.X) + mathgl.Fsqr32(p.Y)) }

func (p Point) Dot(p2 Point) float32    { return p.X*p2.X + p.Y*p2.Y }
func (p Point) Cross(p2 Point) float32  { return p.X*p2.X - p.Y*p2.Y }
func (p Point) Perp() Point             { return Point{-p.Y, p.X} }
func (p Point) RPerp() Point            { return Point{p.Y, -p.X} }
func (p Point) Rotate(p2 Point) Point   { return Point{p.X*p2.X - p.Y*p2.Y, p.X*p2.Y + p.Y*p2.X} }
func (p Point) Unrotate(p2 Point) Point { return Point{p.X*p2.X + p.Y*p2.Y, p.Y*p2.X - p.X*p2.Y} }
func (p Point) Midpoint(p2 Point) Point { return Point{(p.X + p2.X) / 2, (p.Y + p2.Y) / 2} }

func (p Point) CompOp(f func(float32) float32) Point { return Point{f(p.X), f(p.Y)} }

func (p Point) ClampPoint(min, max Point) Point {
	return Point{
		mathgl.Clampf(p.X, min.X, max.X), mathgl.Clampf(p.Y, min.Y, max.Y)}
}
func (p Point) Project(p2 Point) Point {
	return func(f float32) Point { return Point{p2.X * f, p2.Y * f} }(p.Dot(p2) / p2.Dot(p2))
}

func (p Point) Normalize() Point {
	if l := p.Len(); l == 0 {
		return Point{1, 0}
	} else {
		return Point{p.X / l, p.Y / l}
	}
}
func (p *Point) Lerp(p2 Point, alpha float32) Point {
	return func(p3 Point) Point {
		return Point{
			(p3.X + p2.X*alpha),
			(p3.Y + p2.Y*alpha),
		}
	}(p.CompOp(func(i float32) float32 { return (i * (1 - alpha)) }))
}

func ForAngle(a float32) Point {
	return Point{mathgl.Fcos32(a), mathgl.Fsin32(a)}
}

func cross2Vect(a, b, c, d Point) float32 {
	return (d.Y-c.Y)*(b.X-a.X) - (d.X-c.X)*(b.Y-a.Y)
}

type Size struct{ W, H float32 }

func (sz *Size) SetSize(w, h float32) { sz.W, sz.H = w, h }
func (sz Size) Equals(sz2 Size) bool {
	return mathgl.FloatEqual32(sz.W, sz2.W) && mathgl.FloatEqual32(sz.H, sz2.H)
}

type Rect struct {
	Point
	Size
}

func (r *Rect) SetRect(x, y, w, h float32) { r.X, r.Y, r.W, r.H = x, y, w, h }

func (r Rect) Equals(r2 Rect) bool { return (r.Point.Equals(r2.Point) && r.Size.Equals(r2.Size)) }

func (r Rect) MinX() float32 { return r.X }
func (r Rect) MidX() float32 { return r.X + r.W/2 }
func (r Rect) MaxX() float32 { return r.X + r.W }
func (r Rect) MinY() float32 { return r.Y }
func (r Rect) MidY() float32 { return r.Y + r.H/2 }
func (r Rect) MaxY() float32 { return r.Y + r.H }

func (r Rect) ContainsPoint(p Point) bool {
	return p.X >= r.MinX() && p.X <= r.MaxX() && p.Y >= r.MinY() && p.Y <= r.MaxY()
}
func (r Rect) IntersectsRect(r2 Rect) bool {
	return !(r.MaxX() < r2.MinX() ||
		r2.MaxY() < r.MinX() ||
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
	L, R, T, B := r.X, (r.X + r.W), (r.Y + r.H), (r.Y)
	L2, R2, T2, B2 := r2.X, (r2.X + r2.W), (r2.Y + r2.H), (r2.Y)

	swapf32(&R, &L)
	swapf32(&T, &B)
	swapf32(&R2, &L2)
	swapf32(&T2, &B2)

	L3, R3 := mathgl.Fmin32(L, L2), mathgl.Fmax32(R, R2)
	T3, B3 := mathgl.Fmax32(T, T2), mathgl.Fmin32(B, B2)
	return Rect{Point{L3, B3}, Size{R3 - L3, T3 - B3}}
}
