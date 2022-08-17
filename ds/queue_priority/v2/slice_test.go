package main

import (
	"fmt"
	"testing"
)

type intQ []int

// Интересный способ положить элемент в слайс.
func TestSliceByPointer(t *testing.T) {
	q := &intQ{1, 2, 3, 4, 5}

	*q = append(*q, 6)
	*q = append(*q, 7)
	*q = append(*q, 8)
	*q = append(*q, 9)
	*q = append(*q, 10)
	*q = append(*q, 11) // Увеличим cap в 2 раза. С 10 до 20

	t.Logf("%v len=%d cap=%d", q, len(*q), cap(*q)) // type TaskPQ []Task
}

func TestSliceCapIncrementing(t *testing.T) {
	s := []int{}
	fmt.Printf("%v len=%d cap=%d \n", s, len(s), cap(s)) // len=0 cap=0

	for i := 1; i <= 10; i++ {
		s = append(s, i)
		fmt.Printf("%d %v len=%d cap=%d \n", i, s, len(s), cap(s))
	}
	/*
		[] len=0 cap=0
		1 [1] len=1 cap=1	- сначала равны 1
		2 [1 2] len=2 cap=2	- увеличиваем в 2 раза
		3 [1 2 3] len=3 cap=4 - увеличиваем в 2 раза
		4 [1 2 3 4] len=4 cap=4
		5 [1 2 3 4 5] len=5 cap=8 - увеличиваем в 2 раза
		6 [1 2 3 4 5 6] len=6 cap=8
		7 [1 2 3 4 5 6 7] len=7 cap=8
		8 [1 2 3 4 5 6 7 8] len=8 cap=8
		9 [1 2 3 4 5 6 7 8 9] len=9 cap=16 - увеличиваем в 2 раза
		10 [1 2 3 4 5 6 7 8 9 10] len=10 cap=16
	*/
}
