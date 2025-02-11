package semaphore

import (
	"runtime"
	"sync/atomic"
)

type Semaphore struct {
	limit uint32
	a     atomic.Uint32
}

func NewSemaphore(limit uint32) *Semaphore {
	return &Semaphore{
		limit: limit,
	}
}

func (s *Semaphore) Acquire() {
	for {
		c := s.a.Load()
		if c < s.limit && s.a.CompareAndSwap(c, c+1) {
			return
		}
		runtime.Gosched()
	}
}

func (s *Semaphore) Release() {
	for {
		c := s.a.Load()
		if c == 0 || s.a.CompareAndSwap(c, c-1) {
			return
		}
		runtime.Gosched()
	}
}
