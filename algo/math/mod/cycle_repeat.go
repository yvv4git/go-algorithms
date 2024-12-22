//go:build ignore

package main

import "fmt"

// Циклическое повторение через массив или срез.
func main() {
	items := []string{"A", "B", "C", "D"}
	index := 6
	fmt.Println(items[index%len(items)]) // Выведет "B"
}
