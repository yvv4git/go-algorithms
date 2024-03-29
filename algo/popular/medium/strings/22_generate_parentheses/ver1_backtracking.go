package main

import "fmt"

// Функция для генерации всех возможных комбинаций скобок.
func generateParenthesis(n int) []string {
	/*
		METHOD: Backtracking & Recursion
		TIME COMPLEXITY: O(4^n / sqrt(n))
		SPACE COMPLEXITY: O(4^n / sqrt(n))

		Time complexity
		Временная сложность этого алгоритма составляет O(4^n / sqrt(n)),
		так как для каждой позиции в строке (от 0 до 2*n) мы можем выбрать либо открывающую, либо закрывающую скобку.
		Таким образом, для каждой позиции у нас есть 2 возможных варианта действий, и в итоге у нас получается 2^(2n) возможных комбинаций.
		Однако, поскольку нам нужны только "корректные" комбинации (которые удовлетворяют правилу валидности скобочных последовательностей),
		фактически существует только C(2n, n) (комбинаторика) возможных комбинаций, где C - это биномиальные коэффициенты.
		Таким образом, временная сложность сокращается до O(4^n / sqrt(n)).

		Space complexity
		Пространственная сложность составляет O(4^n / sqrt(n)),
		так как в худшем случае глубина рекурсии может достигать n (когда мы открываем все скобки и не закрываем их).
		Однако, так как мы используем стек для отслеживания состояния,
		фактически мы используем линейное количество дополнительного пространства, не считая пространства,
		необходимого для хранения результата.

		Method
		В данном случае используется метод под названием "backtracking".
		Этот метод используется для решения задач, в которых требуется перебрать все возможные комбинации или варианты.
		В данном случае, мы пытаемся найти все возможные комбинации правильных скобочных последовательностей,
		где каждая открывающая скобка имеет соответствующую закрывающую.

		Backtracking работает путем построения решения пошагово, проверки того, что это решение может быть закончено, и,
		если нет, возврата к предыдущему шагу и попытки альтернативного пути.
		В данном случае, мы используем рекурсивный подход, добавляя открывающую скобку, если она еще не достигла максимума,
		и закрывающую скобку, если количество открывающих скобок больше количества закрывающих.
		Если мы достигли максимальной длины скобочной последовательности, мы добавляем ее в результат.

		Backtracking является мощным инструментом для решения задач, где требуется перебрать все возможные варианты,
		но может быть неэффективным, если количество возможных вариантов очень велико.
		В этом случае, как правило, применяются другие алгоритмы, такие как динамическое программирование или метод ветвей и границ.
	*/
	// Создаем пустой срез для хранения результатов.
	result := make([]string, 0)
	// Вызываем функцию backtrack для генерации комбинаций.
	backtrack(&result, "", 0, 0, n)
	// Возвращаем результат.
	return result
}

// Рекурсивная функция для генерации комбинаций скобок.
func backtrack(result *[]string, current string, open int, close int, max int) {
	// Если длина текущей строки равна максимальному количеству пар скобок умноженному на 2,
	// добавляем текущую строку в результат.
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	// Если количество открытых скобок меньше максимального,
	// добавляем открывающую скобку и вызываем backtrack с увеличенным количеством открытых скобок.
	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}
	// Если количество закрытых скобок меньше количества открытых,
	// добавляем закрывающую скобку и вызываем backtrack с увеличенным количеством закрытых скобок.
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}

func main() {
	n := 3
	result := generateParenthesis(n)
	for _, s := range result {
		fmt.Println(s)
	}
}
