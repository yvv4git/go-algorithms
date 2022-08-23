package main

import "fmt"

func main() {
	var s Stack

	s.Push(10)
	s.Push(20)
	s.Push(30)

	v, ok := s.Pop()
	fmt.Println(v, ok)
}
