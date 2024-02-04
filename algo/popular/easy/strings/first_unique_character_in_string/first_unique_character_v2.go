package first_unique_character_in_string

func firstUniqCharV2(s string) int {
	/*
		METHOD: Hash.
		Time complexity : O(n)
		Space complexity : O(n) - так как придется хранить hash для всех N.
	*/
	counter := map[rune]int{}

	// Считаем количество повторов.
	for _, v := range s {
		counter[v]++
	}

	// Идем по N элементам и возвращаем индекс первого уникального элемента.
	for index, symbol := range s {
		if counter[symbol] == 1 {
			return index
		}
	}

	return -1
}
