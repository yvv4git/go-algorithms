package main

func isOneBitCharacterV2(bits []int) bool {
	// Вспомогательная функция для рекурсии
	var isValid func(int) bool
	isValid = func(index int) bool {
		// Базовый случай: если дошли до начала массива, то последний символ корректен
		if index == 0 {
			return true
		}
		// Если текущий элемент равен 1, то следующий символ должен быть двухбитным
		// Поэтому переходим к предыдущему символу
		if bits[index-1] == 1 {
			return isValid(index - 2)
		}
		// Если текущий элемент равен 0, то это однобитный символ
		// Поэтому переходим к предыдущему символу
		return isValid(index - 1)
	}
	// Вызываем вспомогательную функцию с индексом последнего элемента
	return isValid(len(bits) - 1)
}

//func main() {
//	bits := []int{1, 1, 1, 0}
//	fmt.Println(isOneBitCharacter(bits)) // Выводит: true
//}
