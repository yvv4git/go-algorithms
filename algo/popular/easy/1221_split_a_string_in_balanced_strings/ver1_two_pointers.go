package main

import "fmt"

func balancedStringSplit(s string) int {
	/*
		METHOD: Counting Characters
		В этом коде мы используем подсчет символов 'L' и 'R' для поиска сбалансированных подстрок.

		В цикле мы проходим по каждому символу в строке s.
		Если текущий символ равен 'L', то мы увеличиваем счетчик countL.
		Если текущий символ равен 'R', то мы увеличиваем счетчик countR.
		Если countL равен countR, то мы нашли сбалансированную подстроку.

		В конце мы возвращаем количество сбалансированных подстрок.

		TIME COMPLEXITY: O(n)

		SPACE COMPLEXITY: O(1)
	*/
	// Инициализируем счетчики для символов 'L' и 'R'
	countL := 0
	countR := 0
	// Инициализируем счетчик для сбалансированных подстрок
	balancedCount := 0

	// Проходим по каждому символу в строке
	for _, char := range s {
		// Увеличиваем соответствующий счетчик в зависимости от символа
		if char == 'L' {
			countL++
		} else {
			countR++
		}

		// Если количество 'L' равно количеству 'R', значит мы нашли сбалансированную подстроку
		if countL == countR {
			// Увеличиваем счетчик сбалансированных подстрок
			balancedCount++
			// Сбрасываем счетчики для поиска следующей сбалансированной подстроки
			countL = 0
			countR = 0
		}
	}

	// Возвращаем количество сбалансированных подстрок
	return balancedCount
}

func main() {
	// Примеры использования функции
	s1 := "RLRRLLRLRL"
	fmt.Println(balancedStringSplit(s1)) // Вывод: 4

	s2 := "RLLLLRRRLR"
	fmt.Println(balancedStringSplit(s2)) // Вывод: 3

	s3 := "LLLLRRRR"
	fmt.Println(balancedStringSplit(s3)) // Вывод: 1
}
