package main

func minHeightShelves(books [][]int, shelfWidth int) int {
	/*
		METHOD:
		Для решения задачи используется динамическое программирование снизу вверх (bottom-up dynamic programming).
		Мы создаем массив dp, где dp[i] хранит минимальную высоту, которую можно достичь, используя первые i книг.
		Для каждой книги мы проверяем, можно ли добавить её на текущую полку или нужно начать новую полку.
		Если текущую книгу можно добавить к предыдущим, мы обновляем dp[i] на основе предыдущих значений dp и текущей конфигурации книг.

		TIME COMPLEXITY:
		O(n^2), где n — количество книг. Это связано с тем, что для каждой книги мы перебираем все предыдущие книги.

		SPACE COMPLEXITY:
		O(n), так как мы используем массив dp размером n для хранения промежуточных результатов.
	*/
	n := len(books)
	// dp[i] будет хранить минимальную высоту для первых i книг
	dp := make([]int, n+1)

	// Инициализация dp[0] как 0, так как нет книг
	dp[0] = 0

	// Заполняем dp
	for i := 1; i <= n; i++ {
		// Инициализация dp[i] максимальным значением
		dp[i] = 1<<31 - 1
		width := 0
		height := 0
		// Перебираем предыдущие книги
		for j := i; j > 0; j-- {
			width += books[j-1][0]
			if width > shelfWidth {
				break
			}
			height = max(height, books[j-1][1])
			dp[i] = min(dp[i], dp[j-1]+height)
		}
	}

	return dp[n]
}

// Вспомогательная функция для нахождения максимума
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Вспомогательная функция для нахождения минимума
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}