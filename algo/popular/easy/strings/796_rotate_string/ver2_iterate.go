package main

// Функция для проверки, является ли строка B циклическим сдвигом строки A
func rotateStringV2(s string, goal string) bool {
	/*
		METHOD:	Iterate
		Обратите внимание, что этот подход может быть неэффективным для больших строк, поскольку требует проверки всех возможных сдвигов.
		В реальных условиях, где требуется эффективное решение, можно использовать алгоритмы,
		основанные на строковых алгоритмах, такие как алгоритм Кнута-Морриса-Пратта (KMP),
		который позволяет эффективно искать подстроку в строке с использованием предварительно обработанных данных.
		Однако, для понимания алгоритма KMP требуется базовое понимание алгоритмов строки.

		TIME COMPLEXITY: O(n^2), где n - длина строки s.
		Это связано с тем, что для каждой позиции i в строке s мы выполняем операцию циклического сдвига, которая занимает O(n) времени.
		Таким образом, если строка s имеет длину n, мы выполняем n таких операций, что общее время выполнения составляет O(n^2).

		SPACE COMPLEXITY: O(n), так как мы создаем новую строку для хранения результата каждого сдвига.
		Таким образом, пространственная сложность линейно зависит от размера входных данных.
	*/
	// Проверяем, что длины строк равны
	if len(s) != len(goal) {
		return false
	}

	// Если строки равны, то они являются циклическим сдвигом друг друга
	if s == goal {
		return true
	}

	// Проверяем все возможные сдвиги строки s
	for i := 0; i < len(s); i++ {
		if rotateBy(s, i) == goal {
			return true
		}
	}

	return false
}

// Функция для получения циклического сдвига строки на k позиций
func rotateBy(s string, k int) string {
	runes := []rune(s)
	rotated := make([]rune, len(runes))

	for i := 0; i < len(runes); i++ {
		rotated[(i+k)%len(runes)] = runes[i]
	}

	return string(rotated)
}

//func main() {
//	A := "abcde"
//	B := "cdeab"
//	fmt.Println(rotateString(A, B)) // Вывод: true
//}
