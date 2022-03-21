package stack

import "errors"

var (
	ErrEmptyStack    = errors.New("empty stack")
	ErrGoingOffStack = errors.New("going off the stack")
)

// Pop - used for pop values from stack
type Pop func(count int) ([]int64, error)

// NewStackInt64 - used for init instance of Pop
func NewStackInt64(values []int64) Pop {
	idx := 0
	lenValues := len(values)

	return func(count int) ([]int64, error) {
		if len(values) == 0 {
			return nil, ErrEmptyStack
		}

		if idx >= len(values) {
			return nil, ErrGoingOffStack
		}

		offset := idx + count
		begin := idx
		if offset > lenValues {
			offset = idx + (lenValues - idx)
		}

		idx = offset
		return values[begin:offset], nil
	}
}
