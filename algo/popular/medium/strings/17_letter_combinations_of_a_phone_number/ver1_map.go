package _7_letter_combinations_of_a_phone_number

// letterCombinationsV1 функция, которая генерирует все возможные комбинации букв, которые могут быть получены нажатием цифр на клавиатуре телефона.
// Аргументы:
//
//	digits - строка, содержащая цифры от 2 до 9 включительно.
//
// Возвращает:
//
//	[]string - массив строк, представляющих все возможные комбинации букв.
func letterCombinationsV1(digits string) []string {
	/*
		METHOD: Iterate
		TIME COMPLEXITY: O(4^n)
		Space complexity: O(4^n)

		Time complexity
		Временная сложность функции letterCombinationsV1 зависит от количества цифр в входной строке digits.
		Пусть n - количество цифр в digits. Для каждой цифры мы проходим по всем возможным буквы,
		которые могут быть получены нажатием этой цифры. Таким образом, временная сложность функции letterCombinationsV1 составляет O(4^n).

		Space complexity
		Пространственная сложность функции letterCombinationsV1 также зависит от количества цифр в входной строке digits.
		Пусть n - количество цифр в digits. В худшем случае, когда каждая цифра может быть преобразована в 4 буквы (например, цифры 7 и 9 на клавиатуре телефона),
		мы будем хранить 4^n комбинаций букв в результате. Таким образом, пространственная сложность функции letterCombinationsV1 составляет O(4^n).

		Объяснение:
		Если у нас есть n цифр, и каждая цифра может быть преобразована в m букв, то общее количество комбинаций будет равно m^n
	*/
	// Если входная строка пуста, то возвращаем пустой массив.
	if len(digits) == 0 {
		return []string{}
	}

	// Инициализируем результат как пустую строку.
	result := []string{""}

	// Для каждой цифры в входной строке.
	for _, digit := range digits {
		// Получаем все возможные буквы для этой цифры.
		letters := phoneMap[string(digit)]
		// Инициализируем временный массив.
		temp := []string{}

		// Для каждой строки в результате.
		for _, oldStr := range result {
			// Для каждой буквы.
			for _, letter := range letters {
				// Добавляем новую строку, которая является результатом конкатенации старой строки и буквы.
				temp = append(temp, oldStr+string(letter))
			}
		}

		// Обновляем результат.
		result = temp
	}

	// Возвращаем результат.
	return result
}

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
