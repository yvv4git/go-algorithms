//go:build ignore

package main

import "fmt"

// Реализация хеш-функции с ограниченным диапазоном.
func hash(key int, size int) int {
	return key % size
}

func main() {
	key := 123
	size := 10
	fmt.Println("Хеш:", hash(key, size)) // Выведет 3
}
