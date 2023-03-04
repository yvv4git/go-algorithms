package main

// Stack - сначала определим интерфейс.
type Stack interface {
	Push(val interface{})
	Pop() interface{}
	Size() int
}

type StackSlice struct {
	data []interface{}
	size int
}

func (s *StackSlice) Push(val interface{}) {
	s.data = append(s.data, val)
	s.size += 1
}

func (s *StackSlice) Pop() interface{} {
	if s.size == 0 {
		return nil
	}

	result := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size -= 1

	return result
}

func (s *StackSlice) Size() int {
	return s.size
}
