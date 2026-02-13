package main

func getLongestSubsequenceTwoPass(words []string, groups []int) []string {
	/*
		INTUITION:
		Нужно найти самую длинную подпоследовательность из words, где
		соседние элементы имеют разные значения в groups (чередование 0 и 1).

		APPROACH: Two Pass
		Делаем два независимых прохода:
		1. Начинаем с первого элемента с группой 0
		2. Начинаем с первого элемента с группой 1
		Выбираем более длинный результат.

		ПОЧЕМУ ТАК:
		- Поскольку значения бинарные (только 0 или 1), подпоследовательность
		  должна начинаться либо с 0, либо с 1
		- Каждый проход строит максимально длинную подпоследовательность
		  для своего стартового значения
		- Один из проходов гарантированно даст оптимальный результат

		TIME COMPLEXITY: O(n) - два прохода по массиву
		SPACE COMPLEXITY: O(n) - для хранения результатов
	*/

	n := len(words)
	if n == 0 {
		return []string{}
	}

	// Строим подпоследовательность, начиная с указанной группы
	buildSequence := func(startGroup int) []string {
		result := []string{}
		lastGroup := -1 // -1 означает, что ещё ничего не добавили

		for i := 0; i < n; i++ {
			// Если это первый элемент с нужной группой или группа отличается от последней
			if lastGroup == -1 {
				// Ищем первый элемент с нужной стартовой группой
				if groups[i] == startGroup {
					result = append(result, words[i])
					lastGroup = groups[i]
				}
			} else {
				// Добавляем элемент, если его группа отличается от последней
				if groups[i] != lastGroup {
					result = append(result, words[i])
					lastGroup = groups[i]
				}
			}
		}

		return result
	}

	// Два прохода: начинаем с группы 0 и с группы 1
	seq0 := buildSequence(0)
	seq1 := buildSequence(1)

	// Возвращаем более длинную подпоследовательность
	if len(seq0) >= len(seq1) {
		return seq0
	}
	return seq1
}
