package main

func plusOne(digits []int) []int {
	extra := 1
	for i := len(digits) - 1; i >= 0; i-- {
		remain := (digits[i] + extra) % 10 // остаток
		extra = (digits[i] + extra) / 10   // целая часть
		digits[i] = remain
	}

	if extra != 0 {
		return append([]int{extra}, digits...) // создаем слайс, где в начале будет целая часть от деления
	}

	return digits
}
