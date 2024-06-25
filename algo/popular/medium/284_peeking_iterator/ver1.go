package main

import "fmt"

// Iterator представляет собой итератор, который поддерживает методы Next и HasNext.
type Iterator struct {
	data []int
	pos  int
}

func NewIterator(data []int) *Iterator {
	return &Iterator{data: data, pos: 0}
}

func (it *Iterator) Next() int {
	if it.pos >= len(it.data) {
		panic("no more elements")
	}
	val := it.data[it.pos]
	it.pos++
	return val
}

func (it *Iterator) HasNext() bool {
	return it.pos < len(it.data)
}

// PeekingIterator представляет собой итератор, который поддерживает операцию Peek.
type PeekingIterator struct {
	iter        *Iterator
	peekedValue *int
}

func NewPeekingIterator(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter, peekedValue: nil}
}

// Peek возвращает следующий элемент, но не перемещает указатель итератора.
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (pi *PeekingIterator) Peek() int {
	if pi.peekedValue == nil {
		if !pi.iter.HasNext() {
			panic("no more elements")
		}
		val := pi.iter.Next()
		pi.peekedValue = &val
	}
	return *pi.peekedValue
}

// Next возвращает следующий элемент и перемещает указатель итератора.
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (pi *PeekingIterator) Next() int {
	if pi.peekedValue != nil {
		val := *pi.peekedValue
		pi.peekedValue = nil
		return val
	}
	return pi.iter.Next()
}

// HasNext возвращает true, если в итераторе есть еще элементы.
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (pi *PeekingIterator) HasNext() bool {
	return pi.peekedValue != nil || pi.iter.HasNext()
}

func main() {
	data := []int{1, 2, 3}
	iter := NewIterator(data)
	pi := NewPeekingIterator(iter)

	fmt.Println(pi.Next())    // 1
	fmt.Println(pi.Peek())    // 2
	fmt.Println(pi.Next())    // 2
	fmt.Println(pi.Next())    // 3
	fmt.Println(pi.HasNext()) // false
}
