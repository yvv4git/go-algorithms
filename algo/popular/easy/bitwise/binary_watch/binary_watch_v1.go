package binary_watch

import (
	"fmt"
	"strconv"
)

func readBinaryWatchV1(num int) []string {
	/*
		Method: Backtrace
		Time complexity: ???
		Space complexity: ???
	*/
	var res []string
	doRead(&res, num, 0, 0, 1)
	return res
}

func doRead(res *[]string, num int, hour, minute int, pos uint) {
	if num == 0 {
		*res = append(*res, strconv.Itoa(hour)+":"+fmt.Sprintf("%02d", minute))
		return
	}

	for i := pos; i <= 10; i++ {
		if i >= 5 {
			if minute+1<<(i-5) >= 60 {
				return
			}
			doRead(res, num-1, hour, minute+1<<(i-5), i+1)
		} else {
			if hour+1<<(i-1) >= 12 {
				continue
			}
			doRead(res, num-1, hour+1<<(i-1), minute, i+1)
		}
	}
}
