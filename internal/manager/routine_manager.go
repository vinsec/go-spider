package manager

import (
	"sync/atomic"
)

type RoutineManager struct {
	used int32
	cap  int32
}

func NewRoutineManager(cap int) *RoutineManager {
	return &RoutineManager{
		used: 0,
		cap:  int32(cap),
	}
}

func (r *RoutineManager) GetOne() bool {
	if atomic.LoadInt32(&r.used) < r.cap {
		atomic.AddInt32(&r.used, 1)
		return true
	}
	return false
}

func (r *RoutineManager) FreeOne() {
	atomic.AddInt32(&r.used, -1)
}

func (r *RoutineManager) Used() int {
	return int(atomic.LoadInt32(&r.used))
}

func (r *RoutineManager) Left() int {
	return int(r.cap - atomic.LoadInt32(&r.used))
}
