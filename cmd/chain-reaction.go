package main

import "rand"
import "gl"
import "math"
import "time"
import "fmt"
import "sdl"

import . "game/common"

const PlayArea float = 2.0
const ExplodeSize gl.GLfloat = 0.50
const ExplodeIncSize gl.GLfloat = 0.05
const ExplodeDecSize gl.GLfloat = 0.03
const ExplodeHoldTimeout int64 = 2000000000 // 2 seconds
const WinTimeout = 1000000000 
const alpha float = 0.8

type Color [4]float

var WinColor = Color{0.3,0.2,0.2,1.0}
var ExplodeColor = Color{0.8,0.1,0.0,alpha}
var TriggerColors = [5]Color{
	Color{0.1,0.1,1.0,alpha}, 
	Color{1.0,0.1,0.1,alpha}, 
	Color{0.1,1.0,0.1,alpha}, 
	Color{0.1,0.7,0.7,alpha}, 
	Color{0.7,0.7,0.1,alpha} }
var FrameColor = Color{0.9, 0.2, 0.2, 1.0}

const (
	TS_Trig = iota
	TS_ExplodeInflate
	TS_ExplodeHold
	TS_ExplodeDeflate
	TS_Dead
)

const (
	GS_Start = iota
	GS_Running 
	GS_LevelFullfilled 
	GS_LevelWon 
	GS_LevelLost 
	GS_GameCompleted 
)

type trigger struct {
	State int
	Angle gl.GLfloat
	Size gl.GLfloat
	Pos Vector3
	Vel Vector3
	Col Color
	HoldTimeout int64
	Points int
}

type level struct {
	TotalTriggers int   // Number of triggers in this level at start
	RequiredTriggers int // Triggers that needs to explode in order to win this level
}

type ChainReaction struct {
	Triggers []trigger // vector.Vector
	Levels []level
	ActiveLevel int
	NumExploded int
	NumExploding int
	GameState int
	StateChangeTime int64
	Score int
	UIFont *Font
}

func (cmd *ChainReaction) Name() string {
	return "ChainReaction"
}

func (cmd *ChainReaction) SleepTime() int {
	return 25
}

func (cmd *ChainReaction) OnSetup() {

	gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	

	ls := make([]level, 20)

	i := 0
	ls[i] = level{5, 1}
	i++
	ls[i] = level{10, 5}
	i++
	ls[i] = level{18, 10}
	i++
	ls[i] = level{25, 21}
	i++
	ls[i] = level{30, 27}
	i++
	ls[i] = level{37, 32}
	i++
	ls[i] = level{61, 61}
	i++

	cmd.Levels = ls[0:i]
	cmd.ActiveLevel = 0
	cmd.Triggers = make([]trigger, 0, 100)
	cmd.UIFont = NewFont("./Test.ttf", 32)

	cmd.SpawnLevel()
}

func (cmd *ChainReaction) SpawnLevel() {
	
	// First resize the num active trigger for the level
	l := cmd.Levels[cmd.ActiveLevel]
	cmd.Triggers = cmd.Triggers[0:l.TotalTriggers]

	cmd.NumExploded = 0
	cmd.NumExploding = 0
	cmd.GameState = GS_Start
	cmd.StateChangeTime = time.Nanoseconds()

	// Spawn triggers in this level
	for i := 0 ; i < l.TotalTriggers; i++ {
		rx := ((rand.Float() - 0.01) * 2.0 - 1.0) 
		ry := ((rand.Float() - 0.01) * 2.0 - 1.0) 
		p := &Vector3{ Real(rx * PlayArea), Real(ry * PlayArea), -5.0}
		
		rx = rand.Float() - 0.5
		ry = rand.Float() - 0.5 
		v := &Vector3{ Real(rx), Real(ry), 0.0}
		v.Normalize().ImulScale(0.02)
		c := TriggerColors[i % 5]
		cmd.Triggers[i] = trigger{TS_Trig, 0.0, 0.05, *p, *v, c, 0.0, 0}
	}
}

func (cmd *ChainReaction) OnUpdate() {
	
	// load levels on demand
	if KeySym >= sdl.K_1 && KeySym <= sdl.K_8  {
		cmd.ActiveLevel = int(KeySym - sdl.K_1)
		cmd.SpawnLevel()
		KeySym = sdl.K_UNKNOWN
	}

	if cmd.GameState == GS_Start {
		// No action when user has not clicked ready
		if MouseState != 0 {
			cmd.GameState = GS_Running
			MouseState = 0
		}
		return
	}

	if MouseState != 0 && (cmd.GameState == GS_LevelLost || 
		cmd.GameState == GS_LevelWon || 
		cmd.GameState == GS_GameCompleted) {

		if cmd.GameState == GS_GameCompleted {
			cmd.ActiveLevel = 0
			cmd.Score = 0
		}
		if cmd.GameState == GS_LevelWon {
			cmd.ActiveLevel += 1
		}
		cmd.SpawnLevel()
		MouseState = 0
		return
	}

	// If mouse is clicked the spawn an explosion
	if MouseState != 0 && cmd.NumExploded == 0 {
		cmd.SpawnExplosion()
	}

	m := PlayArea

	t := Real(-1.0 * m)
	b := Real(1.0 * m)
	r := Real(1.0 * m)
	l := Real(-1.0 * m)

	// Change positions
	for i := 0; i < len(cmd.Triggers); i++ {
		k := &cmd.Triggers[i]
		if k.State != TS_Trig {
			continue
		}
		k.Pos.Iadd(&k.Vel)
		if k.Pos.X <= l {
			k.Vel.X = -k.Vel.X 
			k.Pos.X = l + l - k.Pos.X
		} else if k.Pos.X >= r { 
			k.Vel.X = -k.Vel.X 
			k.Pos.X = r + r - k.Pos.X
		}

		if k.Pos.Y <= t {
			k.Vel.Y = -k.Vel.Y 
			k.Pos.Y = t + t - k.Pos.Y
		} else if k.Pos.Y >= b { 
			k.Vel.Y = -k.Vel.Y 
			k.Pos.Y = b + b - k.Pos.Y
		}

		k.Angle += 2.0
	}

	// Check for collisions
	for i := 0; i < len(cmd.Triggers); i++ {

		exploder := &cmd.Triggers[i]

		if exploder.State == TS_Trig || exploder.State == TS_Dead {
			continue
		}
		
		for j := 0; j < len(cmd.Triggers); j++ {
			l := &cmd.Triggers[j]

			if l.State != TS_Trig {
				continue
			}
			
			d := l.Pos.Sub(&exploder.Pos).Len() - Real(exploder.Size)
			if d <= 0.0 {
				l.State = TS_ExplodeInflate
				cmd.NumExploded += 1
				cmd.NumExploding += 1
				l.Points = (cmd.NumExploded - 1) * 100
				cmd.Score += l.Points
			}
		}
	}

	level := cmd.Levels[cmd.ActiveLevel]
	if cmd.GameState == GS_Running && 
		cmd.NumExploded >= cmd.Levels[cmd.ActiveLevel].RequiredTriggers+1 {

		// Reached the goal for this level
		cmd.GameState = GS_LevelFullfilled
		cmd.StateChangeTime = time.Nanoseconds()
		
	} else if cmd.GameState == GS_LevelFullfilled && 
		(cmd.NumExploding == 0 || cmd.NumExploded == level.TotalTriggers + 1){

		// Reached the goal for this level and no more explosions
		if cmd.ActiveLevel == len(cmd.Levels) - 1 {
			cmd.GameState = GS_GameCompleted
		} else {
			cmd.GameState = GS_LevelWon
		}
		cmd.StateChangeTime = time.Nanoseconds()
		
	} else if cmd.GameState == GS_Running && cmd.NumExploded != 0 && cmd.NumExploding == 0 {

		// No enough have exploded and no more explosions are
		// present
		cmd.GameState = GS_LevelLost
		cmd.StateChangeTime = time.Nanoseconds()
	}
}

func (cmd *ChainReaction) OnDraw() {

	now := time.Nanoseconds()

	if cmd.GameState == GS_LevelFullfilled {
		dt := now - cmd.StateChangeTime
		d := float(dt) / float(WinTimeout)
		drawWin( d )
	}

	cmd.drawStats()

	drawFrame()

	for i := 0; i < len(cmd.Triggers); i++ {
		k := &cmd.Triggers[i]
		if k.State == TS_Trig {
			drawCircle(k)
		} else if k.State == TS_ExplodeInflate {
			if k.Size >= ExplodeSize {
				k.State = TS_ExplodeHold
				k.HoldTimeout = now + ExplodeHoldTimeout
			} else {
				k.Size += ExplodeIncSize
			}
			drawCircle(k)
		} else if k.State == TS_ExplodeHold {
			if k.HoldTimeout <= now {
				k.State = TS_ExplodeDeflate
			}
			drawCircle(k)
		} else if k.State == TS_ExplodeDeflate {
			if k.Size <= 0.0 {
				k.State = TS_Dead
				cmd.NumExploding -= 1
			} else {
				k.Size -= ExplodeDecSize
				k.Col[3] = float(k.Size/ExplodeSize) * alpha
			}
			drawCircle(k)
		}
	}
	
	cmd.drawPoints()

	cmd.drawMessage()

}

func (cmd *ChainReaction) SpawnExplosion() {
	
	// Build pick ray vector
	window_y := (WinH - MouseY) - WinH / 2;
	norm_y := gl.GLdouble(window_y)/gl.GLdouble(WinH / 2);
	window_x := MouseX - WinW / 2;
	norm_x := gl.GLdouble(window_x)/gl.GLdouble(WinW / 2);

	// find pos on znear plane
	aspect := gl.GLdouble(WinW) / gl.GLdouble(WinH)
	ny := NearHeight * norm_y;
	nx := NearHeight * aspect * norm_x;

	// Now your pick ray vector is (x, y, -zNear).
	nx = nx * (gl.GLdouble(PlayArea) / gl.GLdouble(0.04))
	ny = ny * (gl.GLdouble(PlayArea) / gl.GLdouble(0.04))

	// Add a new exploding trigger at this point
	p := &Vector3{ Real(nx), Real(ny), -5.0}
	v := &Vector3{ 0.0, 0.0, 0.0}
	c := ExplodeColor
	cmd.Triggers = cmd.Triggers[0:len(cmd.Triggers)+1]
	cmd.Triggers[len(cmd.Triggers)-1] = trigger{TS_ExplodeInflate, 0.0, 0.1, *p, *v, c, 0.0, 0}
	fmt.Println(nx, ny)
	cmd.NumExploding = 1
	cmd.NumExploded = 1
}


func init() {
	cmds = cmds[0:len(cmds)+1]
	t := &ChainReaction{}
	cmds[len(cmds)-1] = t
}

func drawFrame() {	

	gl.LoadIdentity()
	gl.LineWidth(5.0)
	gl.Enable(gl.POINT_SMOOTH)
	gl.Enable(gl.LINE_SMOOTH)

	gl.Color4f(gl.GLfloat(FrameColor[0]), gl.GLfloat(FrameColor[1]), gl.GLfloat(FrameColor[2]), 0.1)

	m := PlayArea

	t := 1.0 * m
	b := -1.0 * m
	r := 1.0 * m
	l := -1.0 * m

	gl.Begin(gl.LINE_STRIP)

	zo := -5.0

	gl.Vertex3f(gl.GLfloat(l), gl.GLfloat(t), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(r), gl.GLfloat(t), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(r), gl.GLfloat(b), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(l), gl.GLfloat(b), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(l), gl.GLfloat(t), gl.GLfloat(zo))

	gl.End()

}

func (cmd *ChainReaction) drawStats() {
	if cmd.ActiveLevel >= len(cmd.Levels) {
		// done
		cmd.UIFont.DrawText2D(WinW/2, WinH/2, [3]byte{0,255,0}, 1.0, "Game completed!!!")
	} else {
		l := cmd.Levels[cmd.ActiveLevel]
		str1 := fmt.Sprintf("Level %v - need %v of %v", cmd.ActiveLevel + 1, l.RequiredTriggers, l.TotalTriggers)
		cmd.UIFont.DrawText2D(0, WinH-30, [3]byte{0,150,150}, 0.5, str1)
		ne := cmd.NumExploded
		if ne == 0 {
			ne = 1
		}
		str1 = fmt.Sprintf("Exploded : %v", ne - 1)
		cmd.UIFont.DrawText2D(0, WinH-60, [3]byte{0,150,150}, 0.5, str1)
		str1 = fmt.Sprintf("Score : %v", cmd.Score)
		cmd.UIFont.DrawText2D(0, WinH-90, [3]byte{0,150,150}, 0.5, str1)
	}
}

func (cmd *ChainReaction) drawMessage() {
 	if cmd.GameState == GS_Start {
		l := cmd.Levels[cmd.ActiveLevel]
		str1 := fmt.Sprintf("Level %v - clear %v of %v", cmd.ActiveLevel + 1, l.RequiredTriggers, l.TotalTriggers)
		cmd.UIFont.DrawText2D(WinW/2 - 180, WinH/2, [3]byte{255,255,0}, 1.0, str1)
	} 
 	if cmd.GameState == GS_LevelWon {
		cmd.UIFont.DrawText2D(WinW/2 - 75, WinH/2, [3]byte{0,255,0}, 1.0, "You Win")
	}
 	if cmd.GameState == GS_LevelLost {
		cmd.UIFont.DrawText2D(WinW/2 - 80, WinH/2, [3]byte{255,0,0}, 1.0, "You Lose")
	}
 	if cmd.GameState == GS_GameCompleted {
 		cmd.UIFont.DrawText2D(WinW/2 - 300, WinH/2, [3]byte{0,255,0}, 1.0, "Congratulations - Game Completed")
	}	
}

func (cmd *ChainReaction) drawPoints() {

	for i := 0; i < len(cmd.Triggers); i++ {
		k := &cmd.Triggers[i]
		if k.State >= TS_ExplodeInflate && k.State <= TS_ExplodeHold && k.Points != 0 {

			str := fmt.Sprintf("+%v", k.Points)
 			cmd.UIFont.DrawText3D(k.Pos, [3]byte{0,255,0}, 0.375, str)
		}
 	}
	
}

const DEG2RAD float64 = 3.14159/180.0;

func drawCircle(t *trigger) {

	radius := float64(t.Size)
	pos := t.Pos

	gl.LoadIdentity()
	gl.Color4f(gl.GLfloat(t.Col[0]), gl.GLfloat(t.Col[1]), gl.GLfloat(t.Col[2]), gl.GLfloat(t.Col[3]))

	gl.LineWidth(2.0)
	gl.Translatef( gl.GLfloat(pos.X), gl.GLfloat(pos.Y), gl.GLfloat(pos.Z) )
	
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Begin(gl.POLYGON)
	
	for i := 0; i < 360; i++ {
		var degInRad float64 = float64(i)*DEG2RAD
		gl.Vertex3f( gl.GLfloat(math.Cos(degInRad) * radius), gl.GLfloat(math.Sin(degInRad)*radius), gl.GLfloat(0.0))
	}
	
	gl.End()
	gl.Disable(gl.BLEND)
}


func drawWin(d float) {	

	gl.LoadIdentity()
//	gl.LineWidth(5.0)
//	gl.Enable(gl.POINT_SMOOTH)
//	gl.Enable(gl.LINE_SMOOTH)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	if d > 1.0 {
		d = 1.0
	}

	c := WinColor

	gl.Color4f(gl.GLfloat(c[0]), gl.GLfloat(c[1]), gl.GLfloat(c[2]), gl.GLfloat(d))

	m := PlayArea * 4

	t := 1.0 * m
	b := -1.0 * m
	r := 1.0 * m
	l := -1.0 * m

	gl.Begin(gl.POLYGON)

	zo := -10.0

	gl.Vertex3f(gl.GLfloat(l), gl.GLfloat(t), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(r), gl.GLfloat(t), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(r), gl.GLfloat(b), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(l), gl.GLfloat(b), gl.GLfloat(zo))
	gl.Vertex3f(gl.GLfloat(l), gl.GLfloat(t), gl.GLfloat(zo))

	gl.End()
	gl.Disable(gl.BLEND)

}
