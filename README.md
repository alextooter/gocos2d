Gocos2D

Introduction:
	
Gocos2D is a Google Go(lang) game development library based on the popular 
Cocos2D video game sdk. For those who are unfamiliar with Cocos2D, the design philosophy behind
the project is that it intends to simplify game development. One way it does this is by providing
objects called singletons that handle openGL commands and initialization for you; effectively 
simplifying the game development experience drastically.

Here is an example of what the Cocos2D/Gocos2D experience is like in regards to the singletons.

The main singleton is called the director. The director's job is to handle your games scenes for
you. You can pass the director a scene that contains all your sprites and scene related logic;
then the director pushes it onto the scene stack. The director will then draw your scene to the 
openGL context while updating your game logic for you. If the director is commanded to swap
scenes with another, for instance when you beat level 1 and need to start level 2, the director 
pops level 1 off the stack. Which the garbage collector then collects the popped scene and the director starts the
entire process all over again but this time with the current scene pointer pointing at level 2's
scene.