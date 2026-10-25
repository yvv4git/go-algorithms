package main

func findLength(nums1 []int, nums2 []int) int {
	/*
		TASK: Даны два целочисленных массива nums1 и nums2. Нужно вернуть
		максимальную длину подмассива, который встречается в обоих массивах.

		METHOD: Dynamic programming (динамическое программирование)
		Храним dp[j] - длину наибольшего общего подмассива, который
		заканчивается в текущем элементе nums1 и элементе nums2[j-1].
		Если nums1[i-1] == nums2[j-1], то dp[j] = diag + 1, иначе dp[j] = 0,
		где diag - значение dp[j-1] из предыдущей строки.
		Используем один массив и храним только предыдущую строку таблицы.

		TIME COMPLEXITY: O(n * m)
		- Таблица размера n x m заполняется одним проходом.
		- Каждая ячейка обрабатывается за O(1).

		SPACE COMPLEXITY: O(min(n, m))
		- Хранится один массив dp длины min(n, m) + 1.
		- Более короткий массив всегда становится nums2.
	*/
	// Более короткий массив храним в nums2, чтобы dp занимал O(min(n, m)).
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	dp := make([]int, len(nums2)+1)
	best := 0

	for i := 1; i <= len(nums1); i++ {
		// diag хранит dp[i-1][j-1] из предыдущей строки таблицы.
		diag := 0
		for j := 1; j <= len(nums2); j++ {
			// top - старое значение dp[j], то есть dp[i-1][j].
			top := dp[j]

			if nums1[i-1] == nums2[j-1] {
				dp[j] = diag + 1
				if dp[j] > best {
					best = dp[j]
				}
			} else {
				dp[j] = 0
			}

			diag = top
		}
	}

	return best
}
