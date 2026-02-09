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

	result := maxRepeating(sequence, word)
	fmt.Println(result)
}

func maxRepeating(sequence string, word string) int {
	/*
		INTUITION:
		Нужно найти максимальное k, при котором word, повторенное k раз,
		является подстрокой sequence.

		Алгоритм: итеративно увеличиваем количество повторений word
		и проверяем, является ли полученная строка подстрокой sequence.

		APPROACH: Greedy / Iterative Checking
		Последовательно проверяем word, word*2, word*3... пока они
		являются подстроками sequence.

		ПОЧЕМУ ТАК:
		- Ограничения задачи маленькие (sequence и word до 100 символов)
		- strings.Contains эффективен для таких размеров
		- Простое и понятное решение без сложных структур данных

		TIME COMPLEXITY: O(k * n * m) где k - ответ, n - len(sequence), m - len(word)
		SPACE COMPLEXITY: O(k * m) для хранения текущей проверяемой строки
	*/

	count := 0
	// current - текущая строка word, повторенная count+1 раз
	current := word

	// Пока current является подстрокой sequence, увеличиваем счетчик
	for strings.Contains(sequence, current) {
		count++
		// Добавляем еще одно word к текущей строке для следующей проверки
		current += word
	}

	return count
}
