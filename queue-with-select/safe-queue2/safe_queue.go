// we change Queueimpl7 to Queue
// SafeQueue is a wrapper of Queue, it is safe to used in concurrent model

// ref: https://github.com/golang/proposal/blob/master/design/27935-unbounded-queue-package.md

package queue

import (
	"sync"
	"time"
)

const (
	signalInterval = 200
	signalChanSize = 10
)

type SafeQueue struct {
	q *Queueimpl7
	sync.Mutex
	C chan struct{}
}

func NewSafe() *SafeQueue {
	sq := &SafeQueue{
		q: New(),
		C: make(chan struct{}, signalChanSize),
	}

	go func() {
		ticker := time.NewTicker(time.Millisecond * signalInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if sq.q.Len() > 0 {
					// send signal to indicate there are message waiting to be handled
					select {
					case sq.C <- struct{}{}:
						//signaled
					default:
						// not block this goroutine
					}
				}
			}
		}

	}()

	return sq
}

func (s *SafeQueue) Len() int {
	s.Lock()
	n := s.q.Len()
	s.Unlock()
	return n
}

func (s *SafeQueue) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()

	s.q.Push(v)
}

func (s *SafeQueue) Pop() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Pop()
}

func (s *SafeQueue) Front() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Front()
}
