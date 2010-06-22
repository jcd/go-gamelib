package main

import "sdl"
import "gl"
//import "glu"
// import "fmt"
import "sdl/ttf"

func ResizeWindow(w int32, h int32) {
	// resizeWindow
	gl.MatrixMode(gl.PROJECTION)
	gl.Viewport(0, 0, gl.GLsizei(w), gl.GLsizei(h))
	gl.LoadIdentity()

	//	glu.Perspective( 45.0, ratio, 0.1, 100.0)

	aspect := gl.GLdouble(w) / gl.GLdouble(h)
	zNear := gl.GLdouble(0.1)
	zFar := gl.GLdouble(100.0)

	gl.Frustum(-NearHeight * aspect, 
		NearHeight * aspect, 
		-NearHeight,
		NearHeight, zNear, zFar );
	
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	WinW = int(w)
	WinH = int(h)

}

func SetupVideo() {

	if sdl.Init(sdl.INIT_VIDEO) < 0 {
		panic("Couldn't initialize sdl")
	}

	w := WinW
	h := WinH

	var screen = sdl.SetVideoMode(w, h, 32, SDL_FLAGS ) 

//	var screen = sdl.SetVideoMode(w, h, 32, sdl.OPENGLBLIT | sdl.DOUBLEBUF | sdl.HWSURFACE)
//	var screen = sdl.SetVideoMode(w, h, 32, sdl.OPENGL)

	if screen == nil {
		panic("sdl error")
	}

	if ttf.Init() != 0 {
		panic("ttf init error")
	}

	if (gl.Init() != 0) {
		panic("Couldn't init gl")
	}

	ResizeWindow(screen.W, screen.H)

	gl.ClearColor(0, 0, 0, 0)
//	gl.ClearColor(1, 1, 1, 0)
	gl.ClearDepth(1.0)
	gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.DEPTH_TEST)
	gl.ShadeModel(gl.SMOOTH)
	gl.Hint( gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST );


/*
	if gl.Init() != 0 {
		panic("glew error")
	}
*/
//	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT )

	// gl.Clear(gl.COLOR_BUFFER_BIT)

	// initGL

	// gl.Ortho(0, gl.GLdouble(screen.W), gl.GLdouble(screen.H), 0, -1.0, 1.0)


}

