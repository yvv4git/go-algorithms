package main

// Stack Структура для представления стека
type Stack []rune

// Push Метод для добавления элемента в стек
func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

// Pop Метод для удаления элемента из стека и возврата удаленного элемента
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// IsEmpty Метод для проверки, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Функция для поиска самой длинной последовательности одинаковых символов в строке ДНК
func longestRepetitionV2(dna string) int {
	/*
		METHOD: Stack
		TIME COMPLEXITY: O(n), где n - количество символов в строке ДНК.
		Это обусловлено тем, что алгоритм проходит по каждому символу в строке ровно один раз,
		и для каждого символа выполняется постоянное количество операций (добавление в стек или удаление из стека).
		SPACE COMPLEXITY: O(n), так как в худшем случае (когда все символы уникальны) стек может заполниться до n элементов.
		Однако, в общем случае, количество элементов в стеке будет меньше n, так как мы удаляем элементы из стека только тогда, когда находим последовательность разных символов.
	*/
	var stack Stack
	maxLength := 0
	currentLength := 0

	for _, char := range dna {
		// Если стек пуст или последний добавленный символ не совпадает с текущим символом,
		// то добавляем текущий символ в стек и обновляем текущую длину
		if stack.IsEmpty() || stack[len(stack)-1] != char {
			stack.Push(char)
			currentLength = 1
		} else {
			// Если последний добавленный символ совпадает с текущим символом,
			// увеличиваем текущую длину и проверяем, не превышает ли она максимальную длину
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

//func main() {
//	dna := "ATTCGGGA"
//	fmt.Println(longestRepetition(dna)) // Вывод: 3
//}
