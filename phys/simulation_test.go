package phys

import (
	"testing"
	. "game/common"
)

func TestConstantVelocity(t *testing.T) {

	s := NewConstantVelocity()
	
	var dt Real = 0.3
	for i := 0; i < 10; i++ {
		t.Log("Mass[0] is %s", s.Masses[0])
		s.ResetForces()
		s.Simulate(dt)
	}
}

func TestGravitation(t *testing.T) {

	s := NewGravitation(&Vector3{0.0, -9.81, 0.0})
	
	var dt Real = 0.3
	for i := 0; i < 10; i++ {
		t.Log("Mass[0] is %s", s.Masses[0])
		s.ResetForces()
		s.Solve()
		s.Simulate(dt)
	}
}

func TestMassConnectedWithSpring(t *testing.T) {

	s := NewMassConnectedWithSpring(0.2, &Vector3{0.0,-5.0,0.0})
	
	var dt Real = 0.3
	for i := 0; i < 40; i++ {
		t.Log("Mass[0] is %s", s.Masses[0])
		s.ResetForces()
		s.Solve()
		s.Simulate(dt)
	}
}

