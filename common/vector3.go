package common

import "math"
import "fmt"

type Real float

type Vector3 struct {
	X Real
	Y Real
	Z Real
}

func (v *Vector3) Iadd(a *Vector3) *Vector3 {
	v.X = v.X + a.X
	v.Y = v.Y + a.Y
	v.Z = v.Z + a.Z
	return v
}

func (v *Vector3) Add(a *Vector3) *Vector3 {
	return &Vector3{v.X + a.X, v.Y + a.Y, v.Z + a.Z}
}

func (v *Vector3) Isub(a *Vector3) *Vector3 {
	v.X = v.X - a.X
	v.Y = v.Y - a.Y
	v.Z = v.Z - a.Z
	return v
}

func (v *Vector3) Sub(a *Vector3) *Vector3 {
	return &Vector3{v.X - a.X, v.Y - a.Y, v.Z - a.Z}
}

func (v *Vector3) Imul(a *Vector3) *Vector3 {
	v.X = v.X * a.X
	v.Y = v.Y * a.Y
	v.Z = v.Z * a.Z
	return v
}

func (v *Vector3) Mul(a *Vector3) *Vector3 {
	return &Vector3{v.X * a.X, v.Y * a.Y, v.Z * a.Z}
}

func (v *Vector3) MulScale(a Real) *Vector3 {
	return &Vector3{v.X * a, v.Y * a, v.Z * a}
}

func (v *Vector3) Idiv(a *Vector3) *Vector3 {
	v.X = v.X / a.X
	v.Y = v.Y / a.Y
	v.Z = v.Z / a.Z
	return v
}

func (v *Vector3) Div(a *Vector3) *Vector3 {
	return &Vector3{v.X / a.X, v.Y / a.Y, v.Z / a.Z}
}

func (v *Vector3) DivScale(a Real) *Vector3 {
	return &Vector3{v.X / a, v.Y / a, v.Z / a}
}

func (v *Vector3) ImulScale(a Real) *Vector3 {
	v.X = v.X * a
	v.Y = v.Y * a
	v.Z = v.Z * a
	return v
}

func (v *Vector3) IdivScale(a Real) *Vector3 {
	v.X = v.X / a
	v.Y = v.Y / a
	v.Z = v.Z / a
	return v
}

func (v *Vector3) Neg() *Vector3 {
	return &Vector3{-v.X,-v.Y,-v.Z}
}

func (v *Vector3) Assign(a *Vector3) *Vector3 {
	v.X = a.X
	v.Y = a.Y
	v.Z = a.Z
	return v
}

func (v *Vector3) Len() Real {
	return Real(float64(math.Sqrt(float64(v.X)*float64(v.X) + 
		    float64(v.Y)*float64(v.Y) + 
		    float64(v.Z)*float64(v.Z) )))
}

func (v *Vector3) Normalize() *Vector3 {
	l := v.Len()
	if (l == 0.0) {
		return nil
	}
	v.IdivScale(l)
	return v
}

func (v *Vector3) Normalized() *Vector3 {
	l := v.Len()
	if (l == 0.0) {
		return &Vector3{v.X, v.Y, v.Z}
	}
	return &Vector3{v.X / l, v.Y / l, v.Z / l}
}

func (v *Vector3) Eq(v1 *Vector3) bool {
	return v.X == v1.X && v.Y == v1.Y && v.Z == v1.Z 
}

func (v *Vector3) Lt(v1 *Vector3) bool {
	return v.X < v1.X && v.Y < v1.Y && v.Z < v1.Z
}

func (v *Vector3) Gt(v1 *Vector3) bool {
	return v.X > v1.X && v.Y > v1.Y && v.Z > v1.Z 
}

func (v *Vector3) String() string {
	return fmt.Sprintf("Vector3(%f,%f,%f)", v.X, v.Y, v.Z)
}
