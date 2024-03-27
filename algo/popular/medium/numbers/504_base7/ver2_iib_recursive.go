package main

import "strconv"

func convertToBase7V2(num int) string {
	if num < 0 {
		return "-" + convertToBase7(num*-1)
	}

	if num < 7 {
		return strconv.Itoa(num)
	}

	return convertToBase7(num/7) + strconv.Itoa(num%7)
}
