package main

// FibV2 вычисляет n-ое число Фибоначчи с использованием мемоизации.
// Временная сложность: O(n)
// Пространственная сложность: O(n)
func fibV2(n int, cache map[int]int) int {
	// Базовые случаи
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	// Проверка наличия значения в кэше
	if val, found := cache[n]; found {
		return val
	}

	// Рекурсивные вызовы с мемоизацией
	result := fibV2(n-1, cache) + fibV2(n-2, cache)
	cache[n] = result
	return result
}

//func main() {
//	cache := make(map[int]int)
//	fmt.Println("Fibonacci result:", fibV2(5, cache))
//}
