package main

import (
	"strconv"
	"unicode"
)

type stackNums []int

func (s *stackNums) push(v int) {
	*s = append(*s, v)
}

func (s *stackNums) pop() int {
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out
}

func (s stackNums) isEmpty() bool {
	return len(s) == 0
}

func calculateV3(s string) int {
	/*
		METHOD: Stack
		TIME COMPLEXITY: O(n)
		Space complexity: O(1), используем фиксированное количество памяти
	*/
	stack := new(stackNums)
	out := 0
	currentNum := 0
	lastSign := string('+')

	for i := 0; i < len(s); i++ {

		if unicode.IsDigit(rune(s[i])) {
			v, _ := strconv.Atoi(string(s[i]))
			currentNum = currentNum*10 + v
		}

		if (!unicode.IsDigit(rune(s[i])) && !unicode.IsSpace(rune(s[i]))) || (i+1 == len(s)) {
			if lastSign == "/" {
				stack.push(stack.pop() / currentNum)
			} else if lastSign == "*" {
				stack.push(stack.pop() * currentNum)
			} else if lastSign == "-" {
				stack.push(-currentNum)
			} else if lastSign == "+" {
				stack.push(currentNum)
			}

			lastSign = string(s[i])
			currentNum = 0
		}
	}

	for !stack.isEmpty() {
		out += stack.pop()
	}

	return out
}
