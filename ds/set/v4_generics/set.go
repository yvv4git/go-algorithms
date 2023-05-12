package main

import "fmt"

type Set[E comparable] map[E]struct{}

func NewSet[E comparable](values ...E) Set[E] {
	s := Set[E]{}

	for _, v := range values {
		s[v] = struct{}{}
	}

	return s
}

func (s Set[E]) Add(values ...E) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

// На мой взгляд, все методы ниже являются не обязательными.

func (s Set[E]) Members() []E {
	result := make([]E, 0, len(s))

	for v := range s {
		result = append(result, v)
	}

	return result
}

func (s Set[E]) String() string {
	return fmt.Sprintf("%v", s.Members())
}

func (s Set[E]) Union(s2 Set[E]) Set[E] {
	result := NewSet(s.Members()...)
	result.Add(s2.Members()...)

	return result
}

func (s Set[E]) Intersection(s2 Set[E]) Set[E] {
	result := NewSet[E]()
	for _, v := range s.Members() {
		if s2.Contains(v) {
			result.Add(v)
		}
	}

	return result
}
