package timer

import "time"

type Timer struct {
	StartTime time.Time
	Seconds   chan int
	done      chan int
}

func New() Timer {
	return Timer{}
}

func (timer *Timer) Start() {
	timer.StartTime = time.Now()
	timer.Seconds = make(chan int)
	timer.done = make(chan int)
	go computeTimer(timer.Seconds, timer.done)
}

func (timer *Timer) Stop() {
	timer.done <- -1
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
