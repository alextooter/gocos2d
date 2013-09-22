package gocos2d

var AppID = "gocos2d"

//Node
type (
	INode interface {
		Update()
		Draw()
		OnEnter()
		OnExit()
		Cleanup()

		GetTag() string
		SetTag(string)

		GetRect() (float32, float32)
		SetRect(float32, float32)

		GetPosition() (float32, float32)
		SetPosition(float32, float32)

		GetCamera() (float32, float32)
		SetCamera(float32, float32)

		GetGrid()
		SetGrid()

		GetAnchor() (float32, float32)
		SetAnchor(float32, float32)

		GetRotation() float32
		SetRotation(float32)

		GetScale() (float32, float32, float32)
		SetScale(float32, float32, float32)

		GetZ() float32
		SetZ(float32)
	}
	IScene interface {
		INode
	}
	ISprite interface {
		INode
	}
	ILayer interface {
		INode
	}
	IParticleSystem interface {
		INode
	}
)

//Action

type (
	IAction interface {
	}
)
