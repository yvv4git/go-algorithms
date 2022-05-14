package main

import "strings"

func wordPattern(s string, words string) bool {
	t := strings.Fields(words)
	m := make(map[string]string)
	if len(t) != len(s) {
		return false
	}

	for i, v := range s {
		if val, ok := m[string(v)]; ok {
			if string(t[i]) != string(val) {
				return false
			}
		} else {
			for _, ele := range m {
				if ele == string(t[i]) {
					return false
				}
			}
			m[string(v)] = string(t[i])
		}
	}

	return true
}
