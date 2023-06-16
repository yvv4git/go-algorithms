package binary_watch

import "strconv"

func readBinaryWatchV3(turnedOn int) []string {
	/*
		Method: Not backtrack
		Time complexity: ???
		Space complexity: O(1)
	*/
	const (
		hours   = 12
		minutes = 60
	)
	result := make([]string, 0)

	for h := 0; h < hours; h++ {
		for m := 0; m < minutes; m++ {
			if bitCount(int64(m))+bitCount(int64(h)) == turnedOn {
				temp := ""
				if m < 10 {
					temp = "0" + strconv.Itoa(m)
				} else {
					temp = strconv.Itoa(m)
				}
				result = append(result, strconv.Itoa(h)+":"+temp)
			}
		}
	}

	return result
}

func bitCount(n int64) int {
	result := 0
	temp := strconv.FormatInt(n, 2)
	for i := 0; i < len(temp); i++ {
		if temp[i] == '1' {
			result++
		}
	}

	return result
}
