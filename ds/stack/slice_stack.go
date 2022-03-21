package stack

import "errors"

var (
	ErrEmptyStack    = errors.New("empty stack")
	ErrGoingOffStack = errors.New("going off the stack")
)

type pop func(count int) ([]int64, error)

func initStackInt64(values []int64) pop {
	idx := 0
	lenValuse := len(values)

	return func(count int) ([]int64, error) {
		if len(values) == 0 {
			return nil, ErrEmptyStack
		}

		if idx >= len(values) {
			return nil, ErrGoingOffStack
		}

		offset := idx + count
		begin := idx
		if offset > lenValuse {
			offset = idx + (lenValuse - idx)
		}

		idx = offset
		return values[begin:offset], nil
	}
}
