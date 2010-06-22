package phys

import . "game/common"

type Rope struct {
	*Simulation
	Springs []*Spring
	G Vector3 // gravity
	RopeConnectionPos Vector3 
	RopeConnectionVel Vector3 // moving the rope conn
	GroundRepulsionConst Real // how much the ground repel the masses
	GroundFricConst Real      // friction the ground applies to masses
	GroundAbsorbtionConst Real // absorbtion friction applied to masses by ground
	GroundHeight Real 
	AirFricConst Real
}

func NewRope(nodes int, def_mass Real) *Rope {

	r := &Rope{Simulation: NewSimulation(1,1.0),
	        G: Vector3{0.0, -9.81, 0.0}, 
		GroundRepulsionConst: 100.0,
		GroundFricConst: 0.2,
		GroundAbsorbtionConst: 2.0,
		GroundHeight: -1.5,
	        AirFricConst: 0.02 }

	var spring_const Real = 7000.0
	var spring_len Real = 0.05
	var spring_inner_fric Real = 0.2

	r.Masses = make([]*Mass, nodes)
	r.N = nodes

	r.RopeConnectionPos = Vector3{5.0,0.0,-10.0}
	r.RopeConnectionVel = Vector3{0.0,0.0,0.0}

	// Create masses
	for i := 0; i < nodes; i++ {
		p := Vector3{r.RopeConnectionPos.X + spring_len * Real(i), 0.0, -10.0}
//		r.Masses[i] = &Mass{M: def_mass, Pos: p, PosOld: Vector3{0.0, 0.0, 0.0}}
		r.Masses[i] = &Mass{M: def_mass, Pos: p, PosOld: p}
	}

	r.Springs = make([]*Spring, nodes-1)

	for i := 0; i < nodes - 1; i++ {
		
		r.Springs[i] = NewSpring(r.Masses[i], r.Masses[i+1], spring_const, spring_len, spring_inner_fric)
	}

	return r
}


func (r *Rope) Solve() {

	l := len(r.Springs)
	
	for i := 0; i < l; i++ {
		r.Springs[i].Solve()
//		jj := r.Springs[i].Solve()
//		if i == 78 {
//			fmt.Println("solve", jj)
//		}
	}

//	fmt.Println(r.Masses[78])
//	fmt.Println(r.Masses[79])
//	return

	l = len(r.Masses)
	for i := 0; i < l; i++ {
		m := r.Masses[i]
		m.ApplyForce(r.G.MulScale(m.M)) // Gravity
		m.ApplyForce(m.Vel.Neg().MulScale(r.AirFricConst)) // Air


		// Is the mass below ground level
		if m.Pos.Y < r.GroundHeight {
			v := m.Vel
			v.Y = 0.0

			// The velocity in y direction is omited
			// because we will apply a friction force to
			// create a sliding effect. Sliding is
			// parallel to the ground. Velocity in y
			// direction will be used in the absorption
			// effect.
			m.ApplyForce(v.Neg().MulScale(r.GroundFricConst))
			
			v = m.Vel // get the velocity
			v.X = 0;  // omit the x and z components of the velocity
			v.Z = 0;  // we will use v in the absorption effect
			
			// above, we obtained a velocity which is vertical to the ground and it will be used in
			// the absorption force
			if v.Y < 0 { //let's absorb energy only when a mass collides towards the ground
				m.ApplyForce(v.Neg().MulScale(r.GroundAbsorbtionConst)) //the absorption force is applied
			}

			//The ground shall repel a mass like a spring.
			//By "Vector3D(0, groundRepulsionConstant, 0)" we create a vector in the plane normal direction
			//with a magnitude of groundRepulsionConstant.
			//By (groundHeight - masses[a]->pos.y) we repel a mass as much as it crashes into the ground.
			force := Vector3{0.0, r.GroundRepulsionConst, 0.0}
			force.ImulScale(r.GroundHeight - m.Pos.Y)
			
			//The ground repulsion force is applied
			m.ApplyForce(&force)
		}
	}
}

func (r *Rope) SimulateRope(dt Real) {

	r.Simulate(dt)

	r.RopeConnectionPos.Iadd(r.RopeConnectionVel.MulScale(dt))
	
	if r.RopeConnectionPos.Y < r.GroundHeight {
		r.RopeConnectionPos.Y = r.GroundHeight
		r.RopeConnectionVel.Y = 0.0
	}

	r.Masses[0].Pos = r.RopeConnectionPos
	r.Masses[0].Vel = r.RopeConnectionVel
}
