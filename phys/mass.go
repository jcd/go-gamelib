package phys

import "fmt"
import . "game/common"

type Mass struct {
	M Real // Mass value
	Pos Vector3 
	Vel Vector3
	Force Vector3
	PosOld Vector3 
}

func (m *Mass) NewMass(mass Real) *Mass {
	return &Mass{M:mass}
}

func (m *Mass) ApplyForce(f *Vector3) {
	m.Force.Iadd(f)
}

func (m *Mass) Simulate(dt Real) {
	

	me := 6
	// Euler method
	if me == 1 {
		acc := m.Force.DivScale(m.M)
		m.Vel.Iadd( acc.ImulScale(dt))
		m.Pos.Iadd( m.Vel.MulScale(dt) )
	} 
	// Midpoint method
	if me == 2 {
		acc := m.Force.DivScale(m.M)
		d_acc := acc.Sub(&m.PosOld)
		m.PosOld = *acc

		m.Vel.Iadd( acc.ImulScale(dt) )
		m.Pos.Iadd( m.Vel.MulScale(dt).Iadd( d_acc.MulScale(dt*0.5) ))
	} 
	// Verlet
	if me == 3 {
		acc := m.Force.DivScale(m.M)
		vel := m.Pos.Sub(&m.PosOld) // Vel is the old m.Pos
		m.PosOld = m.Pos
		m.Pos.Iadd( vel.Add( acc.MulScale(dt*dt) ) )
		m.Vel = *m.Pos.Sub(&m.PosOld)
	}
	// Leapfrog
	if me == 4 {
		// acc := m.Force.DivScale(m.M)
		// m.Vel.Iadd( acc.Add(&m.PosOld).ImulScale(dt*0.5) )
		// m.PosOld = *acc
		// m.Pos.Iadd( m.Vel.MulScale(dt) ).Iadd(  m.PosOld.MulScale(dt*dt*0.5) )
		old_vel := m.Vel
		acc := m.Force.DivScale(m.M)
		m.Vel.Iadd( acc.MulScale(dt) ).Iadd(&old_vel).ImulScale(0.5)
		m.Pos.Iadd( m.Vel.MulScale(dt) )
		//		m.Vel.IAdd( acc.MulScale(dt*-0.5) )

		//m.PosOld = *acc
	}

	// Leapfrog 2 (velocity verlet)
	if me == 5 {
		prev_acc := m.PosOld

		// calc x(t+dt) = x(t) + v(t)dt+ 0.5*a(t)*dt^2
		m.Pos.Iadd( m.Vel.MulScale(dt) ).Iadd( prev_acc.MulScale(dt*dt*0.5) )
		
		// calc v(t+dt/2) = v(t) + (a(t)*dt) / 2
		vel_half := m.Vel
		vel_half.Iadd( prev_acc.MulScale(dt*0.5) )

		// calc a(t+dt) = acc
		acc := m.Force.DivScale(m.M)
		
		// calc v(t+dt) = v(t+dt/2) + (a(t)+dt)*dt / 2
		m.Vel = *vel_half.Add(acc.MulScale(0.5*dt))
		
		// Rememver acc for next step
		m.PosOld = *acc
	}
	// Verlet (position)
	if me == 6 {
		
		cur_pos := m.Pos
		
		// x(t+dt) = 2x(t) - x(t-dt) +a(t)dt^2 
		acc := m.Force.DivScale(m.M)
		m.Pos = *m.Pos.ImulScale(2.0).Isub(&m.PosOld).Iadd(acc.MulScale(dt*dt))
		
		// v(t+dt) = (x(t-dt) - x(t)) / dt
		m.Vel = *m.Pos.Sub(&cur_pos).DivScale(dt)
		
		// prepare x(t-dt) for next iteration
		m.PosOld = cur_pos
	}
}

func (m *Mass) String() string {
	return fmt.Sprintf("Mass{M: %f, Pos: %v, Vel: %v, Force: %v", 
		m.M, 
		m.Pos,
		m.Vel,
		m.Force)
}
