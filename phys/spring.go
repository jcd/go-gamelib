package phys

import . "game/common"

type Spring struct {
	Mass1 *Mass
	Mass2 *Mass
	
	SpringConst Real
	SpringLen   Real
	FrictionConst Real
}

func NewSpring(mass1 *Mass, mass2 *Mass,
	spring_const Real, spring_len Real, fric_const Real) *Spring {
	s := &Spring{mass1, mass2, spring_const, spring_len, fric_const}
	return s
}

func (s *Spring) Solve() *Vector3 {

	spring_vec := s.Mass1.Pos.Sub(&s.Mass2.Pos)

	r := spring_vec.Len()

	f := Vector3{}

	if (r != 0.0) {
		spring_vec.IdivScale(r).ImulScale(r - s.SpringLen).ImulScale(-s.SpringConst)
		f.Iadd(spring_vec)
//		fmt.Println("fd", f)
	}

	f.Iadd( s.Mass1.Vel.Sub(&s.Mass2.Vel).Neg().ImulScale(s.FrictionConst) )

//	fmt.Println(f)
	
	s.Mass1.ApplyForce(&f)
	s.Mass2.ApplyForce(f.Neg())
	return &f
}
