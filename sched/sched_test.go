package sched

import (
	"fmt"
	"testing"
	"time"
	. "game/common"
)

func say(s string, start Time, timer Timer) {
	fmt.Println("Saying", s, "With now ", timer.Now(), "-", start, "=", timer.Now() - start);
}


func stopat(s string, stop_time Time) func(ch TaskChan, sc *Scheduler) {
	return func(ch TaskChan, sc *Scheduler) {
		<-ch
		for sc.now < stop_time {
			fmt.Println("Saying", s, "with now ", sc.now, " stopping at", stop_time);
			sc.C <- TaskResult{RUNNING,ch};
			<-ch
		}
		fmt.Println("Finally saying", s, "with now ", sc.now, " stopping at", stop_time);
		sc.C <- TaskResult{COMPLETED,ch};
	}
}

func TestAdd(t *testing.T) {

	fmt.Println("Test Add")

	sc := NewScheduler(100)
	snow := sc.now
	sc.Add( stopat("stop at 1.5 ", snow + 1.5) )

	ticker := time.NewTicker(500000000)

	for sc.Update() != 0 {
		<-ticker.C
	}
}

func TestAddLater(t *testing.T) {

	fmt.Println("Test AddLater")

	sc := NewScheduler(100)
	snow := sc.now
  	sc.AddLater(1.5, stopat("stop at 1.5 ", snow + 2.5) )
	sc.AddLater(4.5,  stopat("stop at 4.5 ", snow + 5.5) )

	ticker := time.NewTicker(500000000)

	for sc.Update() != 0 {
		<-ticker.C
	}
}

func TestAddSimple(t *testing.T) {

	fmt.Println("Test Simple")

	sc := NewScheduler(100)
	snow := sc.now
        ch := sc.AddSimple( func() { say("lalala", snow, sc.timer) } )

	ticker := time.NewTicker(500000000)

	i := 4
	for sc.Update() != 0 {
		<-ticker.C
		if i == 1 {
			ch<-COMPLETED

		}
		i--
	}
}

func TestAddLaterSimple(t *testing.T) {

	fmt.Println("Test AddLaterSimple")

	sc := NewScheduler(100)
	snow := sc.now
  	ch1 := sc.AddLaterSimple(1.5, func() { say("1.5 seconds later", snow, sc.timer) } )
	ch2 := sc.AddLaterSimple(4.5, func() { say("4.5 seconds later", snow, sc.timer) } )

	ticker := time.NewTicker(500000000)

	i := 12
	for sc.Update() != 0 {
		<-ticker.C
		if i == 1 {
 			ch1<-COMPLETED
 			ch2<-COMPLETED
		}
		i--
	}
}

func TestAddIntervalSimple(t *testing.T) {

	fmt.Println("Test AddIntervalLaterSimple")

	sc := NewScheduler(100)
	snow := sc.now
  	ch1 := sc.AddIntervalSimple(1.5, 3, func() { say("1.5 seconds later", snow, sc.timer) } )

	ticker := time.NewTicker(500000000)

	i := 12
	for sc.Update() != 0 {
		<-ticker.C
		if i == 1 {
 			ch1<-COMPLETED
		}
		i--
	}
}

func TestAddIntervalManySimple(t *testing.T) {

	fmt.Println("Test AddIntervalLaterSimple")

	sc := NewScheduler(100)
	snow := sc.now
  	ch1 := sc.AddIntervalSimple(0.5, 5, func() { say("A 1.5 seconds later", snow, sc.timer) } )
  	ch2 := sc.AddIntervalSimple(1.5, 3, func() { say("B 1.5 seconds later", snow, sc.timer) } )
  	ch3 := sc.AddIntervalSimple(1.0, 30, func() { say("C 1.5 seconds later", snow, sc.timer) } )

	ticker := time.NewTicker(500000000)

	i := 12
	for sc.Update() != 0 {
		<-ticker.C
		if i == 1 {
 			ch1<-COMPLETED
 			ch2<-COMPLETED
 			ch3<-COMPLETED
		}
		i--
	}
}
