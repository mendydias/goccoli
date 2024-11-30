package core

import (
	"fmt"
	"os"
	"slices"

	"github.com/rs/zerolog"
)

// This is my implementation of a double-ended queue.
// Deque (double-ended queue) is a data structure that supports insertion and deletion
// operations at both ends. It combines the functionality of both stacks and queues,
// allowing elements to be added or removed from either the front or back in constant O(1) time.
// Unlike a standard queue (FIFO) or stack (LIFO), a deque provides maximum flexibility
// for element access. Common use cases include sliding window problems, palindrome checking,
// and task scheduling where elements need to be processed from either end.
//
// Implementation uses a dynamic circular buffer to maintain O(1) operations at both ends
// while providing efficient memory usage and cache locality. When capacity is reached,
// the buffer automatically grows to accommodate new elements.
//
// The capacity must be a power of two
type Deque struct {
	buf      []int
	count    int
	head     int
	tail     int
	capacity int
}

var logger zerolog.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Str("module", "core/deque").Logger()

// A sufficiently large enough minimum capacity to prevent the data structure from having to overwrite too many elements at a time.
const minCap = 1024

func WithCapacity(capacity int) *Deque {
	logger.Debug().Msgf("Creating a deque with capacity %d", capacity)
	return initDeque(nil, capacity)
}

func (q *Deque) Cap() int {
	logger.Debug().Msg("Requesting capacity of deque")

	if q == nil {
		logger.Debug().Str("Deque", fmt.Sprintf("%v", q)).Msg("Deque provided to Cap() is a nil instance.")
		return 0
	}

	logger.Debug().Int("Capacity", len(q.buf)).Msg("Returning the capacity of the calling deque.")
	return len(q.buf)
}

func (q *Deque) Len() int {
	logger.Debug().Msg("Requesting length of deque")

	if q == nil {
		logger.Debug().Str("Deque", fmt.Sprintf("%v", q)).Msg("Deque provided to Len() is a nil instance.")
		return 0
	}

	logger.Debug().Int("Length", q.count).Msg("Returning the length of the calling deque.")
	return q.count
}

func (q *Deque) PushBack(data int) {
	logger.Debug().Msgf("Pushing %d to the tail of the deque", data)

	if q == nil || len(q.buf) == 0 {
		logger.Debug().Msg("Deque is nil or empty. Initializing.")
		initDeque(q, 0)
	}

	logger.Debug().Str("buffer", fmt.Sprintf("%v", q.buf)).Msg("Pre-push buffer state")
	index := q.getindex(1)
	q.buf[index] = data
	q.updateCount()
	q.updatePointers(index, 1)

	logger.Debug().
		Str("buffer", fmt.Sprintf("%v", q.buf)).
		Int("head", q.head).
		Int("tail", q.tail).
		Int("count", q.count).
		Msg("Post-push buffer state")
}

func (q *Deque) PushFront(data int) {
	logger.Debug().Int("data", data).Msg("Pushing an item to the front")

	if q == nil || len(q.buf) == 0 {
		logger.Debug().Msg("Deque is nil or empty. Initializing.")
		initDeque(q, 0)
	}

	logger.Debug().Str("buffer", fmt.Sprintf("%v", q.buf)).Msg("Pre-push buffer state")
	index := q.getindex(-1)
	q.buf[index] = data
	q.updateCount()
	q.updatePointers(index, -1)

	logger.Debug().
		Str("buffer", fmt.Sprintf("%v", q.buf)).
		Int("head", q.head).
		Int("tail", q.tail).
		Int("count", q.count).
		Msg("Post-push buffer state")
}

func (q *Deque) Head() int {
	logger.Debug().Msgf("Returning element at head = %d", q.head)

	return q.buf[q.head]
}

func (q *Deque) Tail() int {
	logger.Debug().Msgf("Returning element at tail = %d", q.tail)

	return q.buf[q.tail]
}

func (q *Deque) Contains(key int) bool {
	logger.Debug().Msgf("Looking for %d in %v", key, q.buf)

	return slices.Contains(q.buf, key)
}

func initDeque(q *Deque, capacity int) *Deque {
	if q == nil {
		logger.Debug().Msg("Deque is a nil instance. Allocating memory and creating an instance.")

		q = new(Deque)
	}

	if capacity == 0 {
		logger.Debug().Msg("Deque has an empty buffer. Creating min capacity buffer and initializing internal pointers for head and tail.")
		capacity = minCap
	}
	q.buf = make([]int, capacity)
	q.head = 0
	q.count = 0
	q.tail = 0

	return q
}

func (q *Deque) getindex(offset int) int {
	index := 0

	logger.Debug().Int("count", q.count).Msg("Index depends on the item count")
	if q.count == 0 {
		index = 0
	} else {
		capacity := len(q.buf)
		if offset > 0 {
			index = (((q.tail + offset) % capacity) + capacity) % capacity
		} else {
			index = (((q.head + offset) % capacity) + capacity) % capacity
		}
	}

	logger.Debug().Int("index", index).Msg("Returning next index to push back")
	return index
}

func (q *Deque) updateCount() {
	logger.Debug().Msg("Updating count")

	if q.count+1 >= len(q.buf) {
		logger.Debug().Msg("Item count is at capacity. Upper bounded to bufer capacity.")

		q.count = len(q.buf)
	} else {
		logger.Debug().Msg("Incrementing count")

		q.count++
	}
}

// Maintains the invariant: tail - head gives the total count of elements, while maintaining wrap around properties..
func (q *Deque) updatePointers(newIndex int, direction int) {
	logger.Debug().Msg("Updating head and tail pointers")

	if q.count == 1 {
		logger.Debug().Int("head", q.head).Int("tail", q.tail).Msg("Nothing to update, only one element inserted.")
		return
	}

	if direction > 0 {
		q.tail = newIndex
		logger.Debug().Int("new index", newIndex).Msg("Updating pointers to grow with the tail")

		if q.tail > q.head {
			q.head = 0
			logger.Debug().Int("head", q.head).Int("tail", q.tail).Msg("Tail pointer is ahead of head pointer, head is at the start")
		} else {
			q.head = (q.tail + 1) % len(q.buf)
			logger.Debug().Int("head", q.head).Int("tail", q.tail).Msg("Tail pointer is behind of head pointer, head is one greater than tail")
		}
	} else {
		q.head = newIndex
		logger.Debug().Int("new index", newIndex).Msg("Updating pointers to grow with the head")

		if q.head == q.tail {
			capacity := len(q.buf)
			q.tail = (((q.head - 1) % capacity) + capacity) % capacity
			logger.Debug().Int("tail", q.tail).Msg("Head and Tail clash. Recalculating ")
		} else {
			logger.Debug().Msg("Nothing to update with the tail.")
			q.tail = 0
		}
	}
}
