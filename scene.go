package gocos2d

type Scene struct {
	next, prev *Scene
}
type sceneStack struct {
	currentScene *Scene
	length       int
}

func (ss *sceneStack) init() {
	ss.currentScene = nil
	ss.length = 0
}
func (ss *sceneStack) pop() {
	ss.remove()
	println("Popped scene from stack")
}
func (ss *sceneStack) push(s *Scene) {
	if ss.length == 0 {
		ss.currentScene = s
		ss.length++
		println("Pushed scene to empty stack")
		return
	}
	s.prev = ss.currentScene
	ss.currentScene.next = s
	ss.currentScene = s
	ss.length++
	println("Pushed scene to stack")
}
func (ss *sceneStack) replace(s *Scene) {
	ss.remove()
	ss.push(s)
	println("Replaced Scene")
}
func (ss *sceneStack) remove() {
	if ss.length > 0 {
		s := ss.currentScene
		ss.currentScene = ss.currentScene.prev
		ss.currentScene.next = nil
		s.next = nil
		s.prev = nil
		ss.length--
		return
	}
	println("There is nothing in the stack to remove.")
	Running = false
	return
}
