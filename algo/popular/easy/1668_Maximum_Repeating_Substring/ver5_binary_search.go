//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Читаем две строки из ввода.
	var sequence, word string
	fmt.Fscan(in, &sequence)
	fmt.Fscan(in, &word)

	result := maxRepeatingBinary(sequence, word)
	fmt.Println(result)
}

// maxRepeatingBinary находит максимальное k для word^k в sequence с помощью бинарного поиска.
func maxRepeatingBinary(sequence string, word string) int {
	/*
		INTUITION:
		Максимальное k не может превышать len(sequence) / len(word).
		Функция проверки "является ли word^k подстрокой sequence" монотонна:
		если word^k является подстрокой, то word^(k-1) тоже является подстрокой.

		Идея: используем бинарный поиск для нахождения максимального k
		в отсортированном (по k) пространстве возможных ответов.

		APPROACH: Бинарный поиск + проверка
		1. Определяем границы поиска: [0, len(sequence)/len(word)]
		2. Бинарный поиск на каждом шаге проверяет середину
		3. Если word^mid подстрока → ищем в правой половине
		4. Иначе → ищем в левой половине

		БАЗА: low = 0 (word^0 = "" всегда подстрока)
		Верхняя граница: high = len(sequence) / len(word)

		Переход:
		- mid = (low + high + 1) / 2
		- Если contains(sequence, word*mid): low = mid
		- Иначе: high = mid - 1

		ПОЧЕМУ ТАК:
		- Монотонность свойства позволяет использовать бинарный поиск
		- O(log(n/m)) итераций вместо O(n/m) в линейном поиске
		- Эффективно когда m (len(word)) << n (len(sequence))

		TIME COMPLEXITY: O(log(n/m) * n * m)
		SPACE COMPLEXITY: O(k * m) для хранения word*k
	*/

	n := len(sequence)
	m := len(word)

	// Максимально возможное k.
	maxPossibleK := n / m

	low, high := 0, maxPossibleK

	for low < high {
		// Среднее значение, смещенное вправо.
		mid := (low + high + 1) / 2

		// Проверяем, является ли word^mid подстрокой sequence.
		if isKRepeating(sequence, word, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}

// isKRepeating проверяет, является ли word^k подстрокой sequence.
func isKRepeating(sequence string, word string, k int) bool {
	if k == 0 {
		return true
	}

	// Строим word, повторенный k раз.
	repeated := strings.Repeat(word, k)
	return strings.Contains(sequence, repeated)
}
