package v1

type Set struct {
	data map[interface{}]struct{}
}

// New - Creates a new empty set.
func New() *Set {
	return &Set{make(map[interface{}]struct{})}
}

// Insert - Inserts an element into the set.
func (s *Set) Insert(val interface{}) {
	s.data[val] = struct{}{}
}

// Remove - Removes an element from the set. If none was present, nothing is done.
func (s *Set) Remove(val interface{}) {
	delete(s.data, val)
}

// Size - Returns the number of elements in the set.
func (s *Set) Size() int {
	return len(s.data)
}

// Exists - Checks whether an element is inside the set.
func (s *Set) Exists(val interface{}) bool {
	_, ok := s.data[val]
	return ok
}

// Do - Executes a function for every element in the set.
func (s *Set) Do(f func(interface{})) {
	for val, _ := range s.data {
		f(val)
	}
}

// Reset - Clears the contents of a set.
func (s *Set) Reset() {
	*s = *New()
}
