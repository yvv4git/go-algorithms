package _94_decode_string

import (
	"strconv"
	"strings"
)

func decodeStringV2(s string) string {
	_, answer := decode(s, 0)
	return strings.Join(answer, "")
}

func decode(s string, i int) (int, []string) {
	num := 0
	st := []string{}
	for i < len(s) {
		char := string(s[i])
		if val, err := strconv.Atoi(char); err == nil {
			num = num*10 + val
		} else if char == "[" {
			k, l := decode(s, i+1)
			for j := 0; j < num; j++ {
				st = append(st, l...)
			}
			i = k
			num = 0
		} else if char == "]" {
			return i, st
		} else {
			st = append(st, char)
		}
		i += 1
	}
	return i, st
}
