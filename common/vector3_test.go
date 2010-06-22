package common

import (
	"fmt"
	"testing"
)

func TestCmp(t * testing.T) {
	v1 := Vector3{1.0,2.0,3.0}
	v2 := Vector3{1.0,2.0,3.0}
	if ! v1.Eq(&v2) {
		t.Error(v1, "should be equal to", v2)
	}
	v1.X = 2.0
	if v1.Eq(&v2) {
		t.Error(v1, "should not be equal to", v2)
	}
}

func TestLen(t * testing.T) {
	v1 := Vector3{1.0,0.0,0.0}
	l := v1.Len()
	if l != 1.0 {
		t.Error(v1, "should have len 1.0")
	}
}

func TestNormalize(t * testing.T) {
	v1 := Vector3{1.0,5.0,2.5}
	l := v1.Normalize().Len()
	if l != 1.0 {
		t.Error(v1, "should have len 1.0 after Normalize is called - got", fmt.Sprintf("%f",l))
	}
}

func TestNormalized(t * testing.T) {
	v1 := Vector3{1.0,5.0,2.5}
	l := v1.Normalized().Len()
	if l != 1.0 {
		t.Error(v1, "should have len 1.0 after Normalize is called - got", fmt.Sprintf("%f",l))
	}
}
