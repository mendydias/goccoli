package core

import (
	"errors"
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

// Private helper functions
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
	q.capacity = capacity

	return q
}

// Calculates the next index pointer
func (q *Deque) next() int {
	index := 0

	if q.count != 0 {
		index = (q.tail + 1) & (len(q.buf) - 1)
	}

	logger.Debug().Int("index", index).Msg("Returning next index")
	return index
}

// Calculates the previous index pointer
func (q *Deque) prev(reference int) int {
	index := 0

	if q.count != 0 {
		index = (reference - 1) & (len(q.buf) - 1)
	}

	logger.Debug().Int("index", index).Msg("Returning previous index")
	return index
}

// Updates the element count
func (q *Deque) updateCount() {
	logger.Debug().Msg("Updating count")

	if q.count < len(q.buf) {
		logger.Debug().Msg("Incrementing count")
		q.count++
		return
	}

	logger.Debug().Msg("Item count is at capacity. Upper bounded to buffer capacity.")
}

// The following functions maintain the invariant: tail - head gives the total count of elements, while maintaining wrap around properties
func (q *Deque) incrementPointersRight(newIndex int) {
	logger.Debug().Int("new index", newIndex).Msg("Updating pointers to the right. Tail grows.")

	if q.count > 1 {
		q.tail = newIndex
		if q.tail > q.head {
			q.head = 0
		} else {
			q.head = q.next()
		}
	}

	logger.Debug().Int("head", q.head).Int("tail", q.tail).Int("count", q.count).Msg("New pointers after update right.")
}

func (q *Deque) incrementPointersLeft(newIndex int) {
	logger.Debug().Int("new index", newIndex).Msg("Updating pointers to the left. Head grows.")

	if q.count > 1 {
		q.head = newIndex
		if q.head > q.tail {
			q.tail = 0
		} else {
			q.tail = q.prev(q.head)
		}
	}

	logger.Debug().Int("head", q.head).Int("tail", q.tail).Int("count", q.count).Msg("New pointers after update left.")
}

func WithCapacity(capacity int) *Deque {
	// Check if capacity is not a power of two and get the upper bound closest power of two capacity
	logger.Debug().Int("capacity", capacity).Msg("Checking if given capacity is a power of 2")
	if capacity&(capacity-1) != 0 {
		capacity |= capacity >> 1
		capacity |= capacity >> 2
		capacity |= capacity >> 4
		capacity |= capacity >> 8
		capacity |= capacity >> 16
		capacity++
		logger.Debug().Int("new capacity", capacity).Msg("Capacity is not a power of two. Getting the next biggest power of two.")
	}

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
	index := q.next()
	q.buf[index] = data
	q.updateCount()
	q.incrementPointersRight(index)

	logger.Debug().
		Str("buffer", fmt.Sprintf("%v", q.buf)).
		Int("head", q.head).
		Int("tail", q.tail).
		Int("count", q.count).
		Msg("Post-push buffer state")
}

func (q *Deque) PopBack() (int, error) {
	if q == nil {
		err := errors.New("Illegal pop back operation on nil deque.")
		logger.Error().Msg(err.Error())
		return 0, err
	}

	logger.Debug().Msg("Removing the last element. Tail shrinks.")

	if q.count != 0 {
		tail := q.Tail()
		if q.count == 1 {
			logger.Debug().Msg("case: item count is 1")
			q.tail = 0
		} else if q.tail < q.head {
			logger.Debug().Msg("case: item count is greater than 1 and tail is behind head")
			q.tail = q.prev(q.tail)
		} else {
			logger.Debug().Msg("case: item count is greater than 1 and tail is ahead of head")
			q.tail = q.prev(q.tail)
		}
		q.count--

		logger.Debug().
			Int("count", q.count).
			Int("head", q.head).
			Int("tail", q.tail).
			Str("buffer", fmt.Sprintf("%v", q.buf)).
			Msg("Item count is greater than 0. Post pop state.")

		return tail, nil
	} else {
		logger.Debug().
			Int("count", q.count).
			Int("head", q.head).
			Int("tail", q.tail).
			Str("buffer", fmt.Sprintf("%v", q.buf)).
			Msg("Item count is 0. Post pop state.")

		return 0, nil
	}
}

func (q *Deque) PushFront(data int) {
	logger.Debug().Int("data", data).Msg("Pushing an item to the front")

	if q == nil || len(q.buf) == 0 {
		logger.Debug().Msg("Deque is nil or empty. Initializing.")
		initDeque(q, 0)
	}

	logger.Debug().Str("buffer", fmt.Sprintf("%v", q.buf)).Msg("Pre-push buffer state")
	index := q.prev(q.head)
	q.buf[index] = data
	q.updateCount()
	q.incrementPointersLeft(index)

	logger.Debug().
		Str("buffer", fmt.Sprintf("%v", q.buf)).
		Int("head", q.head).
		Int("tail", q.tail).
		Int("count", q.count).
		Msg("Post-push buffer state")
}

func (q *Deque) PopFront() (int, error) {
	if q == nil {
		err := errors.New("Illegal PopFront on nil deque.")
		logger.Error().Msg(err.Error())
		return 0, err
	}

	logger.Debug().Msg("Removing the item at the current head. Head shrinks.")

	if q.count != 0 {
		head := q.Head()
		q.head = q.prev(q.head)
		return head, nil
	} else {
		return 0, nil
	}
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
