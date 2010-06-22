package main

/*

 OpenGL app

*/

import "sdl"
import "gl"
import "game/sched"
import "fmt"
//import "os"
import "flag"

// need to cast pointer from C. This really belongs in sdl package
import "unsafe"
type cast unsafe.Pointer

type Cmd interface {
	Name() string
	SleepTime() int
	OnSetup()
	OnUpdate()
	OnDraw()
}

// Graphics init flags for SDL
const SDL_FLAGS uint32 = 
	sdl.OPENGLBLIT | sdl.OPENGL | sdl.DOUBLEBUF | 
	sdl.HWPALETTE | sdl.HWSURFACE | sdl.HWACCEL | sdl.RESIZABLE

// List of possible commands to present to stdout
var cmds []Cmd = make([]Cmd, 0, 200)

// Scheduler to run functions at specified times
var sch *sched.Scheduler

var MouseX int
var MouseY int
var MouseState uint8
var WinW int = 1080 // 640
var WinH int = 1024 // 480
var KeySym uint32 = sdl.K_UNKNOWN

const NearHeight = gl.GLdouble(0.05)

// Show the commands available to stdout and prompt the
// user to select one
func selectCommand(def_cmd int) *Cmd {
	fmt.Printf("Select a job to run:\n")
	for i := 0; i < len(cmds); i++ {
		fmt.Printf("%v : %v\n", i, cmds[i].Name())
	}
	var idx int
	fmt.Printf("> ")
	if def_cmd == -1 {
		fmt.Scan(&idx)
	} else {
		idx = def_cmd
	}
	fmt.Printf("Using %v %v\n",idx, cmds[idx].Name())
	if idx >= 0 && idx < len(cmds) {
		return &cmds[idx]
	}
	return nil
}

// global ticks variable. 
// Ticks can only be read in sdl.GetTicks() by main thread
// therefore we set this variable
var Ticks uint32 

func main() {
	sch = sched.NewScheduler(0)


	flag.Parse() 

	cmd := selectCommand(CmdNum)

	SetupVideo()

	// Let the selected command initialize
	cmd.OnSetup()

	// Extract sleep time between frames for this command
	sleeptime := cmd.SleepTime()

	var frames uint32 = 0
	// var t0 uint32 = sdl.GetTicks()
	// 	gt := sdl.GetTicks

	if ShowFPS {
		sch.AddInterval(1.0, 2, func(tc sched.TaskChan,sc *sched.Scheduler) {
			
			val := <-tc 
			if val != sched.RUNNING { 
				sc.C<-sched.TaskResult{val,tc}
				return
			} 
			sc.C<-sched.TaskResult{sched.COMPLETED,tc}

			// seconds := (t - t0) / 1000.0
			// fps := float(frames) / float(1.0)
			// os.Stdout.WriteString("Fdafsda")
			//sch.LOG <- "FPS"
			// fmt.Println("FPS")
			return
			// t0 = t
			frames = 0

		});
	}

	var running = true

	for running {

		e := &sdl.Event{}

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false
				break
			case sdl.KEYDOWN:
				kb := e.Keyboard()
				KeySym = kb.Keysym.Sym
				if KeySym == sdl.K_ESCAPE {
					running = false
				} 
				break
			case sdl.MOUSEBUTTONUP, sdl.MOUSEBUTTONDOWN:	
				me := e.MouseButton()
				MouseState = me.State
				break
			case sdl.MOUSEMOTION:	
				me := e.MouseMotion()
				MouseX = int(me.X)
				MouseY = int(me.Y)
				MouseState = me.State
				break
			case sdl.VIDEORESIZE:
				me := (*sdl.ResizeEvent)(cast(e))
				sdl.SetVideoMode(int(me.W), int(me.H), 32, SDL_FLAGS)
				ResizeWindow(me.W, me.H)
			}
		}

		// sch.Update()
		cmd.OnUpdate()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT )
		cmd.OnDraw()
		sdl.GL_SwapBuffers()

		if sleeptime != 0 {
			sdl.Delay(uint32(sleeptime))
		}
		frames++
	}

	sdl.Quit()

}
