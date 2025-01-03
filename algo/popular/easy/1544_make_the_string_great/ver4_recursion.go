package main

func makeGoodV4(s string) string {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	// Проходим по строке и ищем "плохую" пару
	for i := 0; i < len(s)-1; i++ {
		if abs(int(s[i])-int(s[i+1])) == 32 {
			// Если найдена "плохая" пара, удаляем её и рекурсивно вызываем функцию для новой строки
			return makeGood(s[:i] + s[i+2:])
		}
	}
	// Если "плохих" пар нет, возвращаем строку
	return s
}
