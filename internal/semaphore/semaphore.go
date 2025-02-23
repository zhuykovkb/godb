package semaphore

type Semaphore struct {
	limit chan struct{}
}

func NewSemaphore(limit uint32) Semaphore {
	return Semaphore{make(chan struct{}, limit)}
}

func (s *Semaphore) Acquire() {
	if s == nil || s.limit == nil {
		return
	}
	s.limit <- struct{}{}
}

func (s *Semaphore) Release() {
	if s == nil || s.limit == nil {
		return
	}
	<-s.limit
}
