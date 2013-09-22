package gocos2d

import "fmt"

//TODO: Make these methods thread safe using sync/atomic pkg.

type rect struct{ w, h float32 }

func (r *rect) GetRect() (w, h float32) { return r.w, r.h }
func (r *rect) SetRect(w, h float32)    { r.w, r.h = w, h }

type pos struct{ x, y float32 }

func (p *pos) GetPosition() (x, y float32) { return p.x, p.y }
func (p *pos) SetPosition(x, y float32)    { p.x, p.y = x, y }

type camera pos

func (c *camera) GetCamera() (x, y float32) { return c.x, c.y }
func (c *camera) SetCamera(x, y float32)    { c.x, c.y = x, y }

type grid struct{}

func (g *grid) GetGrid() {}
func (g *grid) SetGrid() {}

type anchor pos

func (a *anchor) GetAnchor() (x, y float32) { return a.x, a.y }
func (a *anchor) SetAnchor(x, y float32)    { *a = anchor{x, y} }

type rot float32

func (r *rot) GetRotation() float32  { return float32(*r) }
func (r *rot) SetRotation(f float32) { *r = rot(f) }

type scale struct{ w, h, f float32 }

func (s *scale) GetScale() (width, height, factor float32) { return s.w, s.h, s.f }
func (s *scale) SetScale(width, height, factor float32) {
	s.w, s.h, s.f = width, height, factor
}

type skew pos

func (s *skew) GetSkew() (x, y float32) { return s.x, s.y }
func (s *skew) SetSkew(x, y float32)    { s.x, s.y = x, y }

type tag string

func (i *tag) GetTag() (id string) { return string(*i) }
func (i *tag) SetTag(id string)    { *i = tag(id) }

type zOrder float32

func (z *zOrder) GetZ() float32      { return float32(*z) }
func (z *zOrder) SetZ(depth float32) { *z = zOrder(depth) }

type children struct {
	q    []INode
	dict map[string]int
}

func (c *children) init() {
	c.q = make([]INode, 0)
	c.dict = make(map[string]int)
}

func (c *children) AddChild(child INode) error {
	if _, exists := c.dict[child.GetTag()]; !exists {
		c.dict[child.GetTag()] = len(c.q)
		c.q = append(c.q, child)
		return nil
	}
	return fmt.Errorf("%s already exists.", child.GetTag())
}
func (c *children) GetChild(id string) INode {
	if n, exists := c.dict[id]; exists {
		return c.q[n]
	}
	return nil
}
func (c *children) RemoveChild(id string) {
	if n, exists := c.dict[id]; exists {
		c.q[n] = nil
		delete(c.dict, id)
	}
}
