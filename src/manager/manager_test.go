package manager

import "testing"

const cap = 5

func TestManager(t *testing.T) {
	manager := NewRoutineManager(cap)
	left := manager.Left()
	if left != cap {
		t.Error("test Left() failed")
	}

	ok := manager.GetOne()
	if !ok || manager.Left() != cap-1 {
		t.Error("test GetOne() failed")
	}

	used := manager.Used()
	if used != 1 {
		t.Error("test Used() failed")
	}

	manager.FreeOne()
	if manager.Left() != cap {
		t.Error("test FreeOne() failed")
	}
}
