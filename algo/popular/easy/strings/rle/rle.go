package main

import (
	"strconv"
)

func rleString(in string) string {
	var result string
	var counter int
	var last rune

	if len(in) == 0 {
		return ""
	}

	for _, c := range in {
		counter += 1

		if (last != c) && counter == 1 {
			result += string(c)
			last = c
		} else if (last != c) && counter > 1 {
			result += strconv.Itoa(counter - 1)
			counter = 1
			result += string(c)
			last = c
		}
	}

	// Обработаем последний символ
	result += strconv.Itoa(counter)
	return result

}
