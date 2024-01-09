package main

import (
	"strconv"
	"strings"
)

func myAtoiV1(str string) int {
	str = strings.TrimSpace(str)
	i, _ := strconv.Atoi(str)
	return i
}
