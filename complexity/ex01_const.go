//go:build constant

package main

import "github.com/pterm/pterm"

// Array
var arr = [5]int{1, 2, 3, 4, 5}

func main() {
	pterm.Info.Println(accessArrayByIndex(1))
}

// Constant complexity - O(1)
func accessArrayByIndex(key int) int {
	return arr[key]
}
