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

To install:


Gocos2D currently depends on libglfw for window handling. 

Important: libglfw builds/installs itself as a static library by default. This does not work well with go. Eventhough the
building of this package may succeed without problems, any application you use it in will likely throw up a range of symbol
lookup errors. It is therefore strongly recommended to build libglfw as a SHARED library.

Linux:

First make sure all your basic GNU tools are installed like gcc and make. Download the source tarball for glfw and extract it
to a directory of your choice. Open up the extracted directory. Type "sudo make x11-dist-install" if you are on ubuntu,
debian or fedora. Then you must cd into /usr/local/lib and delete libglfw.a.
                         
Then just run "go get code.google.com/p/gocos2d" or "go get github.com/gocos2d/gocos2d-lib". Both repos are being maintained
concurrenly. 
