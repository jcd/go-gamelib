package phys

import (
	"testing"
	. "game/common"
)


func TestSpring(t * testing.T) {
	
	m1 := Mass{M:1.0}
	m2 := Mass{M:1.0}

	m1.Pos = Vector3{-1.0, 0.0, -30.0}
	m2.Pos = Vector3{1.0, 0.0, -30.0}

	s := NewSpring(&m1, &m2, 0.5, 1.0, 3.2)

	var dt Real = 0.3

	for i := 0; i < 50; i++ {
		t.Log("Mass[0] is %s", m1, "Mass[2] is %s", m2)
		s.Solve()
		m1.Simulate(dt)
		m2.Simulate(dt)
	}
	
}
