package gocos2d

type Node interface {
	OnEnter()
	OnExit()
	Draw()
	AddChild()
	Cleanup()
	ConvertToNodeSpace()
	ConvertToWorldSpace()
}
