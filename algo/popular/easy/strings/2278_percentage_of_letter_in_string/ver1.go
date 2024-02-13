package main

import (
	"fmt"
	"math"
)

// percentageLetter вычисляет процентное соотношение символа letter в строке s.
// Возвращает целое число - процент в виде целого числа, округленного до ближайшего целого.
func percentageLetter(s string, letter byte) int {
	// Инициализируем счетчик вхождений символа.
	count := 0

	// Проходимся по каждому символу в строке.
	for i := 0; i < len(s); i++ {
		// Если текущий символ совпадает с искомым, увеличиваем счетчик.
		if s[i] == letter {
			count++
		}
	}

	// Вычисляем процентное соотношение и округляем в меньшую сторону.
	percentage := float64(count) / float64(len(s)) * 100
	return int(math.Floor(percentage))
}

func main() {
	// Пример использования функции percentageLetter.
	s := "foobar"
	letter := byte('o')
	percentage := percentageLetter(s, letter)
	fmt.Printf("Процентное соотношение символа '%c' в строке '%s' равно %d%%\n", letter, s, percentage)
}
