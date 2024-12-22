//go:build ignore

package main

import "fmt"

// Проверка четности числа.
func main() {
	number := 7
	if number%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
