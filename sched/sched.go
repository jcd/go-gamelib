package sched

/*

*/
import (
	"fmt"
	"container/vector"
	"container/heap"
	. "game/common"
)

type TaskState int 

const (
	RUNNING = 0
	COMPLETED = 1
	FAILED = 2
	SUSPENDED = 3
	RESUMED = 4
)

type TaskChan chan TaskState // Channel to send commands to task

type TaskResult struct {
	State TaskState // State of the task identified by the channel c
	C TaskChan // The channel identifying that task
}

type schedulerChan chan TaskResult // Channel to receive status from task

type logChan chan string // Channel to send log messages to main thread for outputting main log

type laterTask struct {
	count int // Number of times to run
	dt Time   // Interval between repeats
	at Time   // When to run next
	fun func(TaskChan,*Scheduler)  // Function to call
	C TaskChan  // Channel to read commands
}

type laterHeap struct { 
	vector.Vector  
}
func (h *laterHeap) Less(i, j int) bool { return h.At(i).(laterTask).at < h.At(j).(laterTask).at }


type Scheduler struct {
	timer Timer
	before Time
	start_time Time
	now Time
	dt Time
	tasks *vector.Vector
	next_tasks *vector.Vector
	later_tasks *laterHeap
	C schedulerChan
	LOG logChan
}

func (s *Scheduler) AddLater(num Time, fun func(TaskChan,*Scheduler)) TaskChan {
	c := make(TaskChan);
	heap.Push(s.later_tasks, laterTask{1, 0.0, s.now + num, fun, c})
	return c;
}

func (s *Scheduler) Add(fun func(TaskChan,*Scheduler)) TaskChan {
	c := make(TaskChan);
	go fun(c,s)
	s.next_tasks.Push(c)
	return c;
}

// If count is zero then infinitely
func (s *Scheduler) AddInterval(dt Time, count int, fun func(TaskChan,*Scheduler)) TaskChan {
	c := make(TaskChan);
	fmt.Println("Adding interval for ", dt, s.now + dt)
	heap.Push(s.later_tasks, laterTask{count, dt, s.now, fun, c})
	return c;
}

func (s *Scheduler) AddSimple(fun func()) TaskChan {
	c := make(TaskChan);
	go func (tc TaskChan, sc *Scheduler) { 
		val := <-tc 
		for { 
			if val != RUNNING { 
				break
			} 
			fun() 
			sc.C<-TaskResult{RUNNING,tc}
			val = <-tc 
		} 
		sc.C<-TaskResult{val,tc} 
	}(c,s)
	s.next_tasks.Push(c)
	return c;
}

func (s *Scheduler) AddLaterSimple(num Time, fun func()) TaskChan {
	c := make(TaskChan);
	f := func (tc TaskChan, sc *Scheduler) { 
		val := <-tc 
		for { 
			if val != RUNNING { 
				break
			} 
			fun() 
			sc.C<-TaskResult{RUNNING,tc}
			val = <-tc 
		} 
		sc.C<-TaskResult{val,tc} 
	}
	fmt.Println("Adding for ", num, s.now + num)
	heap.Push(s.later_tasks, laterTask{1, 0.0, s.now + num, f, c})
	return c;
}

// If count is zero then infinitely
func (s *Scheduler) AddIntervalSimple(dt Time, count int, fun func()) TaskChan {
	c := make(TaskChan);
	f := func (tc TaskChan, sc *Scheduler) { 
		fmt.Println("intr1 ", tc, fun)
		
		val := <-tc 
		fmt.Println("intr2 ", tc, fun)
		if val != RUNNING { 
		fmt.Println("intr3 ", tc, fun)
			sc.C<-TaskResult{val,tc}
			return
		fmt.Println("intr4 ", tc, fun)
		} 
		fmt.Println("intr5 ", tc, fun)
		fun()
		fmt.Println("intr6 ", tc, fun, val)
		sc.C<-TaskResult{COMPLETED,tc}
		fmt.Println("intr7 ", tc, fun)
	}
	fmt.Println("Adding interval for ", dt, s.now + dt)
	heap.Push(s.later_tasks, laterTask{count, dt, s.now, f, c})
	return c;
}

func (s *Scheduler) Update() int {

	fmt.Println("tick")

	// Swap task queues and clear old one to be ready
	// to fill in upcoming tasks
	a := s.tasks
	s.tasks = s.next_tasks
	s.next_tasks = a
	s.next_tasks.Cut(0, s.next_tasks.Len())

	// Setup time now and get delta time since last update
	s.before = s.now
	s.now = s.timer.Now()
	s.dt = s.now - s.before

	// Add relevant tasks from later_tasks
//	if s.later_tasks.Len() != 0 {
		// fmt.Println("Handling later tasks")
	for s.later_tasks.Len() != 0 {
		// for t := s.later_tasks.At(0).(laterTask); t.at <= s.now ; t = s.later_tasks.At(0).(laterTask) {

		// start the task and remove from later_tasks
		if t := s.later_tasks.At(0).(laterTask); t.at <= s.now {
			// fmt.Println("Spawning later_task ", t.at, s.now)
			fmt.Println("got ", t)
			go t.fun(t.C,s)
			s.tasks.Push(t.C)
			heap.Pop(s.later_tasks)
			if t.count != 1 {
				// Put back in later_tasks to be run again later
				// use s.now as base an not t.at else we might
				// get the task run many times in case of fluctuations in time
				t.at = s.now + t.dt
				t.count--
				fmt.Println("pushing ", t)
				heap.Push(s.later_tasks, t)
			}
		} else {
			// No task ready to start
			break;
		}
	}

	// Spread tasks out on all cores
	for _, t := range *s.tasks {
		// Signal the task to continue
		if v, ok := t.(TaskChan); ok {
			fmt.Println("RUN")
			v<-RUNNING
		} else {
			fmt.Println("Warning: Cannot get channel for task")
		}
	}

	fmt.Println("TASKS ", s.tasks.Len())
	for i := s.tasks.Len(); i > 0; i--  {
		fmt.Println("FETCH")
		msg := <-s.C // Wait for tasks to complete
		fmt.Println("FETCH",msg.State)
		if msg.State == RUNNING {
			// Schedule for next tick
			fmt.Println("RERUN")
			s.next_tasks.Push(msg.C)
		}
	}
	fmt.Println("TASKSDONE ", s.tasks.Len())
	

	more_logs := true
	for more_logs {
		var msg string
		select {
		case msg = <- s.LOG:
			msg += "3"
			fmt.Println("Fdas")
		default:
			more_logs = false
		}
	}

	fmt.Println("End sched ",s.later_tasks.Len(), s.next_tasks.Len())

	// s.tasks has been run and can be seen as 'empty' now
	return s.later_tasks.Len() + s.next_tasks.Len()
	// All done
}

func NewScheduler(buf_size int) *Scheduler {
	h := new(laterHeap)
        heap.Init(h)
	s :=  &Scheduler{tasks : new(vector.Vector), 
		next_tasks : new(vector.Vector), 
		later_tasks : h,
	        timer : NewTimer(),
	        C : make(schedulerChan),
	        LOG : make(logChan)}
	s.start_time = s.timer.Now()
	s.now = s.start_time
	s.before = s.now
	return s
}
