package strings

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func FindIntersection(strArr []string) string {
	// code goes here
	if len(strArr) < 2 {
		return ""
	}

	a1Str := strings.Split(spaceMap(strArr[0]), ",")
	a2Str := strings.Split(spaceMap(strArr[1]), ",")

	var a1 []int
	var a2 []int
	for _, value := range a1Str {
		v, err := strconv.Atoi(value)
		if err != nil {
			return ""
		}
		a1 = append(a1, v)
	}

	for _, value := range a2Str {
		v, err := strconv.Atoi(value)
		if err != nil {
			return ""
		}
		a2 = append(a2, v)
	}

	var resultArr []int

	if len(a1) > len(a2) {
		for _, value := range a1 {
			if binarySearch(a2, value) {
				resultArr = append(resultArr, value)
			}
		}
	} else {
		for _, value := range a2 {
			if binarySearch(a1, value) {
				resultArr = append(resultArr, value)
			}
		}
	}

	var result string
	for _, value := range resultArr {
		result = fmt.Sprintf("%s,%d", result, value)
	}

	return strings.Trim(result, ",")
}

func binarySearch(list []int, target int) bool {
	var low, max = 0, len(list) - 1
	if max < low {
		return false
	}

	for low <= max {
		var mid = low + (max-low)/2
		if list[mid] == target {
			return true
		}

		if target > list[mid] {
			low = mid + 1
		} else {
			max = mid - 1
		}
	}

	return false
}

func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func TestIntersectionTwoDigitalString(t *testing.T) {
	tests := []struct {
		name     string
		inStrArr []string
		want     string
	}{
		{
			name:     "CASE-1",
			inStrArr: []string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"},
			want:     "1,4,13",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindIntersection(tt.inStrArr)
			t.Logf("%v", result)
		})
	}
}
