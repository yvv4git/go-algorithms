package main

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)

	maxLen := lenA
	if lenB > lenA {
		maxLen = lenB
	}

	idx := maxLen - 1
	i := lenA - 1
	j := lenB - 1
	var tmp bool
	var result []rune
	for idx >= 0 {
		var bitA, bitB bool
		if i >= 0 && a[i] == '1' {
			bitA = true
		}
		if j >= 0 && b[j] == '1' {
			bitB = true
		}

		switch {
		case bitA && bitB:
			tmp = true
			result = append(result, '0')
		case (bitA || bitB) && tmp:
			result = append(result, '0')
		case (bitA || bitB) && !tmp:
			result = append(result, '1')
		case (!bitA && !bitB) && tmp:
			tmp = false
			result = append(result, '1')
		default:
			result = append(result, '0')
		}
		idx--
		i--
		j--
	}

	if tmp {
		result = append(result, '1')
	}

	// Развернем строку
	for a, b := 0, len(result)-1; a < b; a, b = a+1, b-1 {
		result[a], result[b] = result[b], result[a]
	}

	return string(result)
}
