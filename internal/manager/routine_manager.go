package manager

import (
	"sync"
)

// RoutineManager is a manager for co-routines.
// It creats a pool of sub-spider co-routines
type RoutineManager struct {
	lock *sync.Mutex
	used int
	cap  int
}

func NewRoutineManager(cap int) *RoutineManager {
	return &RoutineManager{
		lock: &sync.Mutex{},
		used: 0,
		cap:  cap,
	}
}

// create a sub spider co-routine
func (r *RoutineManager) GetOne() bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.used < r.cap {
		r.used += 1
		return true
	}
	return false
}

// release a sub spider co-routine
func (r *RoutineManager) FreeOne() {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.used -= 1
}

// get the number of running sub spider co-routine
func (r *RoutineManager) Used() int {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.used
}

// get the number of left sub spider co-routine
func (r *RoutineManager) Left() int {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.cap - r.used
}
