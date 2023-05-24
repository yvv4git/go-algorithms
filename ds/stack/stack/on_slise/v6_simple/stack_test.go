package v6_simple

import (
	"fmt"
	"testing"
)

func TestStack01(t *testing.T) {
	s := new(stack)

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := 0; i < 10; i++ {
		v := s.Pop()
		fmt.Printf("v: %v \n", v)
	}
	fmt.Printf("s_len:%d s_cap:%d  s:%#v \n", len(*s), cap(*s), *s) // s_len:0 s_cap:16  s:v6_simple.stack{}
}

func TestStack02(t *testing.T) {
	s := new(stack)
	s.Push(777)
	fmt.Printf("Peak: %v \n", s.Peak())
	fmt.Printf("Peak: %v \n", s.Peak()) // Операция Peek() ничего не вытаскивает, только смотрит, что находится на верхушке стека.
}
