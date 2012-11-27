package gocos2d

type (
	INode interface {
		Init(Tag)
		Cleanup()
		Update()
		Draw()

		AddChild(INode)
		RemoveChild(Tag)
		GetChild(Tag) INode
		OnEnter()
		OnExit()
		Visit()

		Transform(uint)
		ConvertTo(uint)

		Anchor() *Anchor
		Position() *Position
		Rotation(bool) *Rotation
		Scale() *Scale
		Skew() *Skew
		Tag() *Tag
		ZOrder() *ZOrder
		Parent() INode
		Grid() *Grid
		Camera() *Camera
		BoundingBox() *BoundingBox
	}
	IScene interface {
		INode
	}
	ISprite interface {
		IsBatchNode() *bool
	}
	ILayer interface {
		IsParallax() *bool
	}
	IParticleSystem interface {
	}
)
