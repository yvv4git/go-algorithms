package main

import "strconv"

func restoreIpAddressesV2(s string) []string {
	var result []string
	for i := 1; i < 4 && i < len(s)-2; i++ {
		for j := i + 1; j < i+4 && j < len(s)-1; j++ {
			for k := j + 1; k < j+4 && k < len(s); k++ {
				s1, s2, s3, s4 := s[:i], s[i:j], s[j:k], s[k:]
				if isValidPart(s1) && isValidPart(s2) && isValidPart(s3) && isValidPart(s4) {
					result = append(result, s1+"."+s2+"."+s3+"."+s4)
				}
			}
		}
	}
	return result
}

func isValidPart(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}
