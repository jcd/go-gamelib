package main

import "game/phys"
import . "game/common"

//import "fmt"
type Ropecmd struct {
	rope * phys.Rope
}

func (cmd *Ropecmd) Name() string {
	return "Rope"
}

func (cmd *Ropecmd) SleepTime() int {
	return 0
}

func (cmd *Ropecmd) OnSetup() {
	cmd.rope = phys.NewRope(100, 0.04)
	cmd.rope.RopeConnectionVel = Vector3{2.3,0.3,0.0}
	cmd.rope.RopeConnectionVel.ImulScale(0.25)
	cmd.rope.GroundFricConst = 0.97
	cmd.rope.GroundHeight = -3.0 
/*	
	sch.AddIntervalSimple(8.0, 10000, func() {
		cmd.rope.RopeConnectionVel = *cmd.rope.RopeConnectionVel.Neg()
	});
*/
}

func (cmd *Ropecmd) OnUpdate() {
	cmd.rope.ResetForces()
	cmd.rope.Solve()
	cmd.rope.SimulateRope(0.002)
}

func (cmd *Ropecmd) OnDraw() {
	l := len(cmd.rope.Masses)
	for i := 0; i < l; i++ {
		drawMass(cmd.rope.Masses[i], 0.1)
	}
}

func init() {
	cmds = cmds[0:len(cmds)+1]
	t := &Ropecmd{}
	cmds[len(cmds)-1] = t
}
