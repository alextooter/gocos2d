package gocos2d

//Node
type (
	INode interface {
		Node_() *Node
		Update()
		Draw()
		OnEnter()
		OnExit()
		Cleanup()
	}
	IScene interface {
		INode
	}
	ISprite interface {
		INode
		IsBatchNode() *bool
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
		Action_() *Action
	}
)
