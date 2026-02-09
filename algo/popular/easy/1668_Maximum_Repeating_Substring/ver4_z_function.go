//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Читаем две строки из ввода.
	var sequence, word string
	fmt.Fscan(in, &sequence)
	fmt.Fscan(in, &word)

	result := maxRepeatingZ(sequence, word)
	fmt.Println(result)
}

// maxRepeatingZ находит максимальное k для word^k в sequence с помощью Z-функции.
func maxRepeatingZ(sequence string, word string) int {
	/*
		INTUITION:
		Z-функция Z[i] показывает длину наибольшего префикса строки,
		который также является префиксом суффикса, начинающегося с i.
		Это позволяет эффективно находить все вхождения word в sequence.

		Идея: Z-значения в позициях, соответствующих началу word в sequence,
		показывают длину совпадения. Если Z[i] >= len(word), значит word
		полностью совпадает с subsequence в этой позиции.

		APPROACH: Z-функция
		1. Строим строку: word + "$" + sequence
		2. Вычисляем Z-функцию для этой строки
		3. Проверяем Z-значения в позициях len(word)+1... для совпадений word

		БАЗА: Z[0] = 0 (по определению)

		Переход Z-алгоритма:
		- Поддерживаем окно [l, r] с максимальным r
		- Если i <= r: используем Z[i-r] = min(r-i+1, Z[i-l])
		- Иначе: расширяем окно вручную

		ПОЧЕМУ ТАК:
		- Z-функция работает за O(n+m)
		- Позволяет находить все вхождения word за один проход
		- Эффективна для поиска паттерна в тексте

		TIME COMPLEXITY: O(n + m) где n = len(sequence), m = len(word)
		SPACE COMPLEXITY: O(n + m) для хранения Z-функции
	*/

	m := len(word)
	combined := word + "$" + sequence
	z := make([]int, len(combined))

	maxK := 0
	currentK := 0

	// Вычисляем Z-функцию.
	l, r := 0, 0
	for i := 1; i < len(combined); i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}

		// Расширяем окно вручную.
		for i+z[i] < len(combined) && combined[z[i]] == combined[i+z[i]] {
			z[i]++
		}

		// Обновляем окно.
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	// Проверяем Z-значения для позиций word + "$".
	for i := m + 1; i < len(combined); i++ {
		// Z[i] - это длина совпадения word с subsequence[i-(m+1):]
		if z[i] >= m {
			currentK++
			if currentK > maxK {
				maxK = currentK
			}
		} else {
			currentK = 0
		}
	}

	return maxK
}
