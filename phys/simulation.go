package phys

import . "game/common"

//import "fmt"

type Simulator interface {
	ResetForces()
	Solve() 
	Simulate(dt Real)
}

type Simulation struct {
	N int
	Masses []*Mass
}

func NewSimulation(count int, def_mass Real) *Simulation {
	s := Simulation{0, make([]*Mass,0, count)}
	return &s
}

func (s *Simulation) Release() {
	for i := 0; i < s.N; i++ {
		s.Masses[i] = nil
	}
	s.N = 0
}

func (s *Simulation) GetMass(index int) *Mass {
	if index < 0 || index >= s.N {
		return nil
	}
	return s.Masses[index]
}

func (s *Simulation) AddMass(m *Mass) {
	s.Masses = s.Masses[0:len(s.Masses)+1]
	s.Masses[s.N] = m
	s.N += 1
}

func (s *Simulation) NewMass(mass Real) *Mass {
	m := &Mass{M:mass}
	s.Masses = s.Masses[0:len(s.Masses)+1]
	s.Masses[s.N] = m
	s.N += 1
	return m
}

func (s *Simulation) Simulate(dt Real) {
	for i := 0; i < s.N; i++ {
		s.Masses[i].Simulate(dt)
	}
}

func (s *Simulation) ResetForces() {
	reset := Vector3{}
	for i := 0; i < s.N; i++ {
		s.Masses[i].Force.Assign(&reset)
	}
}

// Constant velocity simulation
type ConstantVelocity struct {
	*Simulation
}

func NewConstantVelocity() *ConstantVelocity {
	s := &ConstantVelocity{NewSimulation(1,1.0)}
	m := s.NewMass(1.0)
	m.Vel = Vector3{X:1.0} // vel 1 in x dir
	return s
}

// Gravition simulation
type Gravitation struct {
	*Simulation
	Gravity Vector3
}

func (s *Gravitation) Solve() {
	for i := 0; i < s.N; i++ {
		m := s.Masses[i]
		f := s.Gravity.MulScale(m.M)
		m.ApplyForce(f)
	}
}

func NewGravitation(dir *Vector3) *Gravitation {
	s := &Gravitation{NewSimulation(1,1.0), *dir}
	m := s.NewMass(1.0)
	m.Pos = Vector3{-5.0, 0.0, -30.0}
	m.PosOld = m.Pos
	m.Vel = Vector3{5.0, 15.0, 0.0} 
	return s
}

// MassConnectedWithSpring simulation
type MassConnectedWithSpring struct {
	*Simulation
	SpringConstant Real
	ConnectionPos Vector3
}

func (s *MassConnectedWithSpring) Solve() {
	for i := 0; i < s.N; i++ {
		m := s.Masses[i]
		springVec := m.Pos.Sub(&s.ConnectionPos)
		f := springVec.Neg().MulScale(s.SpringConstant)
		m.ApplyForce(f)
	}
}

func NewMassConnectedWithSpring(spring_const Real, conn_pos *Vector3) *MassConnectedWithSpring {
	s := &MassConnectedWithSpring{NewSimulation(1,1.0), spring_const, *conn_pos}
	m := s.NewMass(1.0)
	m.Pos = *conn_pos.Add(&Vector3{10.0, 0.0, 0.0}) // offset the pos of the mass 10m from the conn_pos
	m.Vel = Vector3{0.0, 0.0, 0.0} 
	return s
}

