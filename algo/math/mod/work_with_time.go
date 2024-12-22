//go:build ignore

package main

import "fmt"

// Работа с временными интервалами&
func main() {
	seconds := 12345
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	fmt.Printf("Минуты: %d, Остаток секунд: %d\n", minutes, remainingSeconds)
}
