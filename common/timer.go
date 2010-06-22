package common

import (
	"time"
)

type Time float64

type Timer struct {
	TimeOffset int64
}

func NewTimer() Timer {
	return Timer{time.Nanoseconds()}
}

func (t *Timer) Now() Time {
	return Time(float64(time.Nanoseconds()-t.TimeOffset) / 1000000000.0)
}
