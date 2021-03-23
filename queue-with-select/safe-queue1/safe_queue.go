// we change Queueimpl7 to Queue
// SafeQueue is a wrapper of Queue, it is safe to used in concurrent model

// ref: https://github.com/golang/proposal/blob/master/design/27935-unbounded-queue-package.md

package queue

import (
	"sync"
)

type SafeQueue struct {
	q *Queueimpl7
	sync.Mutex
}

func NewSafe() *SafeQueue {
	sq := &SafeQueue{
		q: New(),
	}

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
