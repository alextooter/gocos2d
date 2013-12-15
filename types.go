package gocos2d

//TODO: Make these methods thread safe using sync/atomic pkg.

type Rectangle interface {
	GetRect() (float64, float64)
	SetRect(float64, float64)
}
type rect struct{ w, h float64 }

func (r *rect) GetRect() (w, h float64) { return r.w, r.h }
func (r *rect) SetRect(w, h float64)    { r.w, r.h = w, h }

type Position interface {
	GetPos() (float64, float64)
	SetPos(float64, float64)
}
type pos struct{ x, y float64 }

func (p *pos) GetPos() (x, y float64) { return p.x, p.y }
func (p *pos) SetPos(x, y float64)    { p.x, p.y = x, y }

type Camera interface {
	GetCam() (float64, float64)
	SetCam(float64, float64)
}
type camera pos

func (c *camera) GetCam() (x, y float64) { return c.x, c.y }
func (c *camera) SetCam(x, y float64)    { c.x, c.y = x, y }

/*
type Grid interface{
	GetGrid()
	SetGrid
}
type grid struct{}

func (g *grid) GetGrid() {}
func (g *grid) SetGrid() {}
*/

type Anchor interface {
	GetAnchor() (float64, float64)
	SetAnchor(float64, float64)
}
type anchor pos

func (a *anchor) GetAnchor() (x, y float64) { return a.x, a.y }
func (a *anchor) SetAnchor(x, y float64)    { *a = anchor{x, y} }

type Rotation interface {
	GetRot() float64
	SetRot(float64)
}
type rot float64

func (r *rot) GetRot() float64  { return float64(*r) }
func (r *rot) SetRot(f float64) { *r = rot(f) }

type Scale interface {
	GetScale() (float64, float64, float64)
	SetScale(float64, float64, float64)
}
type scale struct{ w, h, f float64 }

func (s *scale) GetScale() (width, height, factor float64) { return s.w, s.h, s.f }
func (s *scale) SetScale(width, height, factor float64) {
	s.w, s.h, s.f = width, height, factor
}

type Skew interface {
	GetSkew() (float64, float64)
	SetSkew(float64, float64)
}
type skew pos

func (s *skew) GetSkew() (x, y float64) { return s.x, s.y }
func (s *skew) SetSkew(x, y float64)    { s.x, s.y = x, y }

type ZOrder interface {
	GetZ()
	SetZ()
}

type zOrder float64

func (z *zOrder) GetZ() float64      { return float64(*z) }
func (z *zOrder) SetZ(depth float64) { *z = zOrder(depth) }
