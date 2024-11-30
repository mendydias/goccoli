package core

import "testing"

func TestCapNil(t *testing.T) {
	var q *Deque = nil
	if q.Cap() != 0 {
		t.Errorf("Expected 0 capacity in a nil Deque, found %v", q.Cap())
	}
}

func TestLenNil(t *testing.T) {
	var q *Deque = nil
	if q.Len() != 0 {
		t.Errorf("Expected 0 length in a nil Deque, found %v", q.Len())
	}
}

func TestLenCapPushBackTail(t *testing.T) {
	q := new(Deque)
	val := 10
	q.PushBack(val)
	if q.Len() != 1 {
		t.Errorf("Expected a length of 1 but got %v", q.Len())
	}
	if q.Cap() != minCap {
		t.Errorf("Expected a starting min capacity of 16 but go %v", q.Cap())
	}
	tail := q.Tail()
	if tail != val {
		t.Errorf("Expected %d to be the tail, but found %v", val, tail)
	}
	head := q.Head()
	if head != val {
		t.Errorf("Expected %d to be in the head but found %v", val, head)
	}
}

func TestLenCapPushBackHeadTailFull(t *testing.T) {
	capacity := 4
	q := WithCapacity(capacity)
	val, val2, val3, val4 := 1, 2, 3, 4
	q.PushBack(val)
	q.PushBack(val2)
	q.PushBack(val3)
	q.PushBack(val4)
	if q.Len() != q.capacity {
		t.Errorf("Expected length to be at capacity, but found %d", q.Len())
	}
	head := q.Head()
	if head != val {
		t.Errorf("Expected the head item to be %d but found %d", val, head)
	}
	tail := q.Tail()
	if tail != val4 {
		t.Errorf("Expected the tail item to be %d but found %d", val4, tail)
	}
}

func TestLenCapPushBackHeadTailOverCap(t *testing.T) {
	capacity := 4
	q := WithCapacity(capacity)
	val, val2, val3, val4, val5 := 1, 2, 3, 4, 6
	q.PushBack(val)
	q.PushBack(val2)
	q.PushBack(val3)
	q.PushBack(val4)
	q.PushBack(val5)
	if q.Len() != q.capacity {
		t.Errorf("Expected length to be at capacity, but found %d", q.Len())
	}
	head := q.Head()
	if head != val2 {
		t.Errorf("Expected the head item to be %d but found %d", val, head)
	}
	tail := q.Tail()
	if tail != val5 {
		t.Errorf("Expected the tail item to be %d but found %d", val4, tail)
	}
	if q.Contains(val) {
		t.Errorf("Expected the deque not to contain the overwritten previous head, but found %d", val)
	}
}

func TestOverflowTwice(t *testing.T) {
	capacity := 3
	q := WithCapacity(capacity)
	vals := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		q.PushBack(vals[i])
	}
	head := q.Head()
	if head != 3 {
		t.Errorf("Expected head to be the %d, but found %d", 3, head)
	}
	tail := q.Tail()
	if tail != vals[5] {
		t.Errorf("Expected the tail to be the last item in the overflowed array, but found %d", tail)
	}
}

func TestPushFrontHeadTailLen(t *testing.T) {
	val := 10
	q := new(Deque)
	q.PushFront(val)
	if q.Len() != 1 {
		t.Errorf("Expected a length of 1 when pushed one item to the front, found %d", q.Len())
	}
	head := q.Head()
	if head != val {
		t.Errorf("Expected %d as the head item, but found %d", val, head)
	}
	tail := q.Tail()
	if tail != val {
		t.Errorf("Expected %d as the tail item, but found %d", val, tail)
	}
}

func TestPushFrontTwiceCorrectHeadCorrectTail(t *testing.T) {
	val, val2 := 10, 20
	q := new(Deque)
	q.PushFront(val)
	q.PushFront(val2)
	if q.Len() != 2 {
		t.Errorf("Expected a length of two for two items but found %d", q.Len())
	}
	head := q.Head()
	if head != val2 {
		t.Errorf("Expected the head item to be %d, the last pushed front, but found %d", val2, head)
	}
	tail := q.Tail()
	if tail != val {
		t.Errorf("Expected the tail item to be %d, the first pushed front, but found %d", val, tail)
	}
}

func TestPushFrontFullLen(t *testing.T) {
	capacity := 4
	q := WithCapacity(capacity)
	vals := []int{1, 2, 3, 4}
	for _, item := range vals {
		q.PushFront(item)
	}
	if q.Len() != capacity {
		t.Errorf("Expected length of 4 but found %d", q.Len())
	}
	head := q.Head()
	if head != vals[capacity-1] {
		t.Errorf("Expected head item to be %d but found %d", vals[capacity-1], head)
	}
	tail := q.Tail()
	if tail != vals[0] {
		t.Errorf("Expected tail tiem to be %d but found %d", vals[0], tail)
	}
}

func TestPushFrontOverCap(t *testing.T) {
	capacity := 4
	q := WithCapacity(capacity)
	vals := []int{1, 2, 3, 4, 5}
	for _, item := range vals {
		q.PushFront(item)
	}
	if q.Len() != capacity {
		t.Errorf("Expected length of %d but found %d", capacity, q.Len())
	}
	head := q.Head()
	if head != vals[4] {
		t.Errorf("Expected head item to be %d but found %d", vals[4], head)
	}
	tail := q.Tail()
	if tail != vals[1] {
		t.Errorf("Expected tail item to be %d but found %d", vals[1], tail)
	}
}

func TestPushFrontOverCapTwice(t *testing.T) {
	capacity := 3
	q := WithCapacity(capacity)
	vals := []int{1, 2, 3, 4, 5, 6}
	for _, item := range vals {
		q.PushFront(item)
	}
	head := q.Head()
	tail := q.Tail()
	if head != 6 {
		t.Errorf("Expected the head item to be %d but found %d", 6, head)
	}
	if tail != 3 {
		t.Errorf("Expected the tail item to be %d but found %d", 4, tail)
	}
}

func TestPushFrontPushBackSimult(t *testing.T) {
	val, val2 := 10, 20
	q := new(Deque)
	q.PushBack(val)
	q.PushFront(val2)
	head := q.Head()
	tail := q.Tail()
	if q.Len() != 2 {
		t.Errorf("Expected a length of 2 but found %d", q.Len())
	}
	if head != val2 {
		t.Errorf("Expected the head item to be %d, but found %d", val2, head)
	}
	if tail != val {
		t.Errorf("Expected the tail item to  be %d, but found %d", val, tail)
	}
}
