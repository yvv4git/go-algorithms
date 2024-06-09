package main

func calculateV2(s string) int {
	var ret, num int
	stack := []int{}
	op := '+'

	// Appending an arbitary operator makes the parsing easier
	s += "+"

	for _, ch := range s {
		if ch == ' ' {
			continue
		} else if '0' <= ch && ch <= '9' {
			num = 10*num + int(ch-'0')
		} else {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}

			num = 0
			op = rune(ch)
		}
	}

	for _, v := range stack {
		ret += v
	}

	return ret
}
