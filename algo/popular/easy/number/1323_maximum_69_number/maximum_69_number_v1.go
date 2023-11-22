package _323_maximum_69_number

import (
	"strconv"
	"strings"
)

func maximum69NumberV1(num int) int {
	r, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return r
}
