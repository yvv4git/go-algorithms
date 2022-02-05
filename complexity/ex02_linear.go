//go:build linear

package main

import "github.com/pterm/pterm"

// Array
var arr = [5]int{1, 2, 3, 4, 5}

func main() {
	pterm.Info.Println(searchInArraySequentially(5))
}

// Linear complexity - O(n)
// O(n) ограничивает нашу функцию сверху.
func searchInArraySequentially(want int) int {
	result := 0
	for _, value := range arr {
		result += 1
		if value == want {
			return result
		}
	}

	return result
}
