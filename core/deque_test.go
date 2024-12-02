package core

import "testing"

func setupDequeWithItems() (*Deque, int) {
	capacity := 8
	q := WithCapacity(capacity)
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, item := range vals {
		q.PushBack(item)
	}
	return q, capacity
}

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

func TestPopBackLenCap(t *testing.T) {
	q, capacity := setupDequeWithItems()
	oldTail, _ := q.PopBack()
	if q.Cap() != capacity {
		t.Errorf("Expected the capacity to remain %d but found %d", capacity, q.Cap())
	}
	if q.Len() != capacity-1 {
		t.Errorf("Expected the length to be %d but found %d", capacity-1, q.Len())
	}
	tail := q.Tail()
	if tail != 7 {
		t.Errorf("Expected the tail item to be one less %d, but found %d", 7, tail)
	}
	head := q.Head()
	if head != 1 {
		t.Errorf("Expected the head to remain unchanged at %d but found %d", 1, head)
	}
	if oldTail != 8 {
		t.Errorf("Expected PopBack to return the last tail item %d, but found %d", 8, oldTail)
	}
}

func TestPopBackNilPointer(t *testing.T) {
	var q *Deque = nil
	_, err := q.PopBack()
	if err == nil {
		t.Errorf("Expected error when PopBack called on nil Deque, instead found %v", err)
	}
}

func TestOverFlowPopBack(t *testing.T) {
	q, capacity := setupDequeWithItems()
	q.PushBack(9)
	oldTail, err := q.PopBack()
	head := q.Head()
	tail := q.Tail()
	if q.Len() != capacity-1 {
		t.Errorf("Expected length %d but found %d", 7, q.Len())
	}
	if oldTail != 9 {
		t.Errorf("Expected old tail item to be %d but found %d", 9, oldTail)
	}
	if err != nil {
		t.Errorf("Expected no errors but found %v", err)
	}
	if tail != 8 {
		t.Errorf("Expected new tail item to be %d, but found %d", 8, tail)
	}
	if head != 2 {
		t.Errorf("Expected the head to be %d but found %d", 2, head)
	}
}

func TestPopFrontNilPointer(t *testing.T) {
	var q *Deque = nil
	_, err := q.PopFront()
	if err == nil {
		t.Errorf("Expected error when popping the front of a nil deque but found %v", err)
	}
}

func TestPopFrontLenCap(t *testing.T) {
	q, capacity := setupDequeWithItems()
	oldHead := q.Head()
	head, err := q.PopFront()
	newHead := q.Head()
	if q.Len() != capacity-1 {
		t.Errorf("Expected the length to be %d but found %d", capacity-1, q.Len())
	}
	if err != nil {
		t.Errorf("There cannot be any errors on a non-nil deque: %v", q)
	}
	if head != oldHead {
		t.Errorf("Expected %d to be popped but found %d", oldHead, head)
	}
	if newHead != 2 {
		t.Errorf("Expected the new head to point to 2 but found %d", newHead)
	}
}

func TestPopFrontOverflow(t *testing.T) {
	q, _ := setupDequeWithItems()
	q.PushFront(9) // At this point the head should be this
	oldHead := q.Head()
	head, _ := q.PopFront()
	newHead := q.Head()
	tail := q.Tail()
	if head != oldHead {
		t.Errorf("Expected %d to be popped off but found %d", oldHead, head)
	}
	if newHead != 1 {
		t.Errorf("Expected the new head to be %d but found %d", 1, newHead)
	}
	if tail != 7 {
		t.Errorf("Expected the tail to be same at %d but found %d", 7, tail)
	}
}

func TestPopBackSingleItem(t *testing.T) {
	q := new(Deque)
	val := 10
	q.PushBack(val)
	tail, _ := q.PopBack()
	if q.Len() != 0 {
		t.Errorf("Expected length to be 0 but found %d", q.Len())
	}
	if tail != val {
		t.Errorf("Expected %d to be popped out but found %d", val, tail)
	}
}

func TestPopFrontSingleItem(t *testing.T) {
	q := new(Deque)
	val := 10
	q.PushBack(val)
	head, _ := q.PopFront()
	if q.Len() != 0 {
		t.Errorf("Expected length to be 0 but found %d", q.Len())
	}
	if head != val {
		t.Errorf("Expected %d to be popped out but found %d", val, head)
	}
}

func TestPopBackPopFrontTillEmptyThenAddOnceLenCapHeadTail(t *testing.T) {
	q, capacity := setupDequeWithItems()
	for i := 0; i < capacity/2; i++ {
		q.PopBack()
	}
	for i := 0; i < capacity/2; i++ {
		q.PopFront()
	}

	q.PushBack(1)
	q.PushFront(2)
	q.PushBack(3)

	head := q.Head()
	tail := q.Tail()

	if q.Len() != 3 {
		t.Errorf("Expected length to be 3, but found %d", q.Len())
	}

	if head != 2 {
		t.Errorf("Expected head to be %d but found %d", 2, head)
	}

	if tail != 3 {
		t.Errorf("Expected tail to be %d but found %d", 3, tail)
	}

	if !q.Contains(1) {
		t.Error("Should contain 1")
	}

	if q.Contains(8) {
		t.Error("Should not contain 8")
	}
}

func TestPopFrontPushFront(t *testing.T) {
	q, _ := setupDequeWithItems()
	q.PopFront()
	val := 25
	q.PushFront(val)
	head := q.Head()
	if head != val {
		t.Errorf("The new head item should be %d but found %d", val, head)
	}
}

func TestPopBackPushBack(t *testing.T) {
	q, _ := setupDequeWithItems()
	q.PopBack()
	val := 25
	q.PushBack(val)
	tail := q.Tail()
	if tail != val {
		t.Errorf("The new head item should be %d but found %d", val, tail)
	}
}
