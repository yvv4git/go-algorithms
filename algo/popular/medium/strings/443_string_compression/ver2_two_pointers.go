package main

import "strconv"

func compressV2(chars []byte) int {
	l := len(chars)

	i, k, gcnt := 0, 0, 0

	for i < l {
		e := chars[i]

		for j := i; j < l; j++ {
			if chars[j] == e {
				gcnt++
			} else {
				break
			}
		}

		chars[k] = e
		k++
		if gcnt > 1 {
			myString := strconv.Itoa(gcnt)
			for _, ee := range myString {
				chars[k] = byte(ee)
				k++
			}
		}

		i += gcnt
		gcnt = 0
	}

	return k
}
