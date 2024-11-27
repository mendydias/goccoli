package core

import "testing"

func TestCapNil(t *testing.T) {
	var q *Deque = nil
	if q.Cap() != 0 {
		t.Errorf("Expected 0 capacity in a nil Deque, found %v", q.Cap())
	}
}
