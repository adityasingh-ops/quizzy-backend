package timer

import (
	"time"
)

type Timer struct {
	duration time.Duration
	done     chan bool
}

func NewTimer(seconds int) *Timer {
	return &Timer{
		duration: time.Duration(seconds) * time.Second,
		done:     make(chan bool),
	}
}

func (t *Timer) Start() {
	go func() {
		time.Sleep(t.duration)
		t.done <- true
	}()
}

func (t *Timer) Stop() {
	close(t.done)
}

func (t *Timer) Done() <-chan bool {
	return t.done
}