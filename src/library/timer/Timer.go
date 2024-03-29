package timer

import "time"

type Timer struct {
	StartTime time.Time
	Duration  int
	Tick      chan int
	done      chan int
}

func New() Timer {
	return Timer{}
}

func (timer *Timer) Start() {
	timer.StartTime = time.Now()
	timer.Tick = make(chan int)
	timer.done = make(chan int)
	go computeTimer(timer.Tick, timer.done)
}

func (timer *Timer) Stop() {
	timer.done <- -1
}

func (timer *Timer) Clear() {
	timer.Tick = make(chan int)
	timer.done = make(chan int)
	timer.Duration = 0
}

func computeTimer(c, done chan int) {
	seconds := 0

	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(1 * time.Second)
			seconds++
			c <- seconds
		}
	}
}
