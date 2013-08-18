package doc

/*
Define node for the director to call on the interface without type assertion.
This will be embedded in all other interfaces, like sprite and scene, to ensure
the developer that extends a gocos2d type will create a consistent type the 
gocos2d engine will know how to work with.*/
type node interface {
	Node_() *gocos2d.Node
	Init()
	Update()
	Draw()
	OnEnter()
	OnExit()
	Cleanup()
}

/*
Here we define a base Node type that will be the base foundation
of every gocos2d type.  
*/
type (
	Node struct {
		anchor   Anchor
		position Position
		rotation Rotation
		camera   Camera
		skew     Skew
		scale    Scale
		tag      Tag
		z        ZOrder
		bbox     BoundingBox
		grid     Grid
		parent   INode
		children *children
	}
	children []node
)

//This Node_() function is defined so the director can
//access the underlying node's fields without having to use
//reflection. Type assertion will not work for custom types like
//the groundhog sprite listed below. 
func (this *Node) Node_() *Node {
	return this
}
func (this *Node) Init(id Tag) {
	this.tag = id
	tmp := make(children, 0)
	this.children = &tmp
}

//func (this *Node).Update(), etc... go here so Node implements the node interface.

//gocos2d defines an interface for all base nodes in the gocos2d library. 
type sprite interface {
	node
	Sprite_() *Sprite
	IsBatchNode() *bool
}

/*
Sprite embeds an anonymous *Node which means that Sprite implements
the node interface now because Node's exported methods are visible at Sprite's
namespace level. This includes Node_(), which is important because it ensures all gocos2d 
types are able to return *Node's for gocos2d to use.  
*/
type Sprite struct {
	*Node
	IsBatchNode bool
}

func (this *Sprite) Sprite_() *Sprite {
	return this
}

/*
This method overloads the underlying Node's Init(). Which means the director only 
has to do the following to initialze an entire scene's set of nodes.  
		for _, node := range director.scene.children {
			node.Init()
		}

No need to assert types or anything like that. Its all rather straight forward 
and completely capable of being launched concurrently using goroutines simply by changing
`node.Init()` to `go node.Init()`. This includes all goroutine calls in recursive Init() 
calls as well.    
*/
func (this *Sprite) Init(t Tag) {
	this.Node = new(Node)
	this.Node.Init(t)
}

/*
a node will not see this, however a type that implements 
the sprite interface will. Calls like this allow for a more conditionally programmable
engine that abstracts the nit and grit away from the developer. 

For example in cocos2d-iphone, you have a CCSprite and then you have
a CCSpriteBatchNode. The difference is that if a CCSpriteBatchNode contains children 
sprites they all have to draw from the same texture2d. This means that all your sprites 
can be drawn in one opengl call. Which is highly valuable for a game, especially when using
sprite sheets.

My problem with cocos2d's node heirarchy is that this kind of feature should be a 
no-brainer. If a sprite contains children, then I assume it is likely because the 
developer wants to make a sprite batch. So gocos2d will assume this functionality is desired
unless it is explicitly told not to. 

If the developer wants to do something like dynamically change the sprites appearance on
the image at runtime, then they can set the child's flag and the texture2d will be made
into a copy of the parent sprite instead. 

For users who want to create some insane parent->child node scene graph where nodes of 
layers can point to children of sprites that each point to nodes of particle effects.
A sprite child, is not necessarily a child of the underlying node's children. It is 
only assumed they are. The only requirement is that tag names must stay consistent.     
*/
func (this *Sprite) IsBatchNode() bool {
	return this.isBatchNode
}

//Users custom type, implemented in their games "$GAME/groundhog.go" to 
//use with all their groundhog sprites.
type Groundhog struct {
	gocos2d.Sprite
}

/*These methods will overload the underlying types. The user can either
directly modify the underlying node directly (not recommend), 
or they can use actions and scheduled events that the director may
handle concurrently. Abstracting concurrency away from the player at this
point is ideal, however not restrictive because actions and events are 
much like node's in that they can easily be extended and customized.  
*/
func (this *Groundhog) Update() {

}
func (this *Groundhog) Draw() {

}
func (this *Groundhog) OnEnter() {

}
func (this *Groundhog) OnExit() {

}
