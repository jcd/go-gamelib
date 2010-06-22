package phys

import (
	"testing"
	. "game/common"

//	"fmt"
)

func TestRope(t *testing.T) {

	s := NewRope(80, 0.05)
	
	var dt Real = 0.1
	for i := 0; i < 1000; i++ {
		
		if i % 100 == 0 {
			t.Log("Mass[79] is %s", s.Masses[79])
		}
		s.ResetForces()
		//		fmt.Println("r", s.Masses[79].Force)

		s.Solve()
		s.SimulateRope(dt)
	}
}

