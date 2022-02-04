//go:build logn

package main

import "github.com/pterm/pterm"

// Array
var arr = [5]int{1, 2, 3, 4, 5}

func main() {
	pterm.Info.Println("Binary search status[5]: ", binarySearch(5, arr[:]))
	pterm.Info.Println("Binary search status[6]: ", binarySearch(6, arr[:]))
}

// Logarithmic complexity - O(log n)
func binarySearch(want int, data []int) bool {
	low := 0
	high := len(data) - 1

	for low <= high {
		median := (low + high) / 2
		if data[median] < want {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(data) || data[low] != want {
		return false
	}

	return true
}
