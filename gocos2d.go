package gocos2d

type (
	INode interface {
		Init(Tag) INode
		Cleanup()
		Update()
		Draw()
		OnEnter()
		OnExit()
		Visit()

		Transform(uint)
		ConvertTo(uint)

		Anchor() *Anchor
		Position() *Position
		Rotation(bool) *Rotation
		Scale() *Scale
		Skew() *Position
		Tag() *Tag
		ZOrder() *ZOrder
		Children() *Children
		Parent() INode
		Grid() *Grid
		Camera() *Camera
		BoundingBox() *BoundingBox
	}
	IScene interface {
		INode
		IsCurrent() *bool
	}
	ISprite interface {
		INode
		IsBatchNode() *bool
	}
	ILayer interface {
		INode
		IsParallax() *bool
	}
	IParticleSystem interface {
		INode
	}
)
