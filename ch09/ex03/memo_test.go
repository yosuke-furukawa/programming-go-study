package ex03

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	Sequential(t, m, done)
}

func TestCancel(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	go func(done chan struct{}) {
		time.Sleep(time.Second / 10) // sleep 100 ms
		close(done)
	}(done)
	Sequential(t, m, done)
}
