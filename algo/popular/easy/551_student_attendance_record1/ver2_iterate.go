//go:build ignore

package main

import "fmt"

func checkRecordV2(s string) bool {
	/*
		METHOD: Iterate
		В этом решении мы используем цикл for для итерации по строке s по одному символу за раз.
		Мы подсчитываем количество 'A' и последовательных символов 'L' вручную, не используя пакет strings.
		Если количество 'A' превышает 1 или количество подряд идущих 'L' превышает 2, мы немедленно возвращаем false.
		Если мы проитерировали весь цикл без вызова return false, это означает, что студент может быть принят,
		и мы возвращаем true.

		TIME COMPLEXITY: O(n), где n - количество символов в строке s.

		SPACE COMPLEXITY: O(1), так как мы не используем дополнительную память,
		которая бы растягивалась с увеличением размера входных данных.
	*/
	absentCount := 0
	lateCount := 0
	maxLateCount := 0

	for _, char := range s {
		if char == 'A' {
			absentCount++
			lateCount = 0 // Сбрасываем счетчик подряд идущих 'L'
		} else if char == 'L' {
			lateCount++
			if lateCount > maxLateCount {
				maxLateCount = lateCount
			}
		} else {
			lateCount = 0 // Сбрасываем счетчик подряд идущих 'L'
		}

		if absentCount > 1 || maxLateCount > 2 {
			return false
		}
	}

	return true
}

func main() {
	// Примеры использования
	fmt.Println(checkRecordV2("PPALLP")) // Вывод: true
	fmt.Println(checkRecordV2("PPALLL")) // Вывод: false
}
