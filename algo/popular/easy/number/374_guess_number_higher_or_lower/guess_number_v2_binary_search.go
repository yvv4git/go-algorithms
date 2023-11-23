package _74_guess_number_higher_or_lower

func guessNumberV2(n int) int {
	return binarySearch(0, n)
}

func binarySearch(start, finish int) int {
	if start > finish { // Если начало диапазона больше конца, то функция возвращает -1, что означает, что загаданное число не найдено.
		return -1
	}
	result := (finish-start)/2 + start // Вычисляем среднее значение между start & finish.

	switch guess(result) { // Обрати внимание, что на сервере LeetCode знает, какое число загадано. У меня это глобальная переменная pick.
	case -1:
		return binarySearch(start, result-1) // -1: загаданное число больше выбранного
	case 1:
		return binarySearch(result+1, finish) // 1: загаданное число меньше выбранного
	default:
		return result // Загаданное число и угаданное совпали
	}
}
