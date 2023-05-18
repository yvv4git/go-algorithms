package first_unique_character_in_string

func firstUniqCharV1(s string) int {
	/*
		Method: Hash.
		Time complexity : O(n)
		Space complexity : O(1) - так как allocated константное количество данных.
	*/
	// This is a constant space allocation (ie: always 26)
	var counts = make([]int, 26)

	// Insert all the character's count into counts array.
	// Получается здесь надо пройти 1 раз по N элементам.
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++ // Для каждого символа определяем его позицию в массиве и на этом месте инкриминируем счетчик.
	}

	// Find the first element with count 1.
	// Здесь надо пройти второй раз по N элементам.
	for i := 0; i < len(s); i++ {
		if counts[s[i]-'a'] == 1 { // Мы ищем именно не повторяющийся элемент, у повторяющихся счетчик > 1.
			return i
		}
	}

	return -1
}
