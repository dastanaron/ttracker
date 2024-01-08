package timer

import "time"

type Timer struct {
	StartTime time.Time
	Duration  chan int
	done      chan int
}

func New() Timer {
	return Timer{}
}

func (timer *Timer) Start() {
	timer.StartTime = time.Now()
	timer.Duration = make(chan int)
	timer.done = make(chan int)
	go computeTimer(timer.Duration, timer.done)
}

func (timer *Timer) Stop() {
	timer.done <- -1
}

func computeTimer(c, done chan int) {
	duration := 0

	for {
		select {
		case <-done:
			close(done)
			close(c)
			return
		default:
			time.Sleep(1 * time.Second)
			duration++
			c <- duration
		}
	}
}
