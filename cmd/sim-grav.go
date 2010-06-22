package main

import "game/phys"
import . "game/common"

type Simgrav struct {
	sim * phys.Gravitation
}

func (cmd *Simgrav) Name() string {
	return "Simulate gravity"
}

func (cmd *Simgrav) SleepTime() int {
	return 25
}

func (cmd *Simgrav) OnSetup() {
	cmd.sim = phys.NewGravitation(&Vector3{0.0, -9.81, 0.0})
}

func (cmd *Simgrav) OnUpdate() {
	cmd.sim.ResetForces()
	cmd.sim.Solve()
	cmd.sim.Simulate(0.02)
}

func (cmd *Simgrav) OnDraw() {
	drawMass(cmd.sim.Masses[0], 1.0)
}

func init() {
	cmds = cmds[0:len(cmds)+1]
	t := &Simgrav{}
	cmds[len(cmds)-1] = t
}
