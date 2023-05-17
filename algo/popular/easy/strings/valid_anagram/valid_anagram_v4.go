package valid_anagram

func isAnagramV4(s string, t string) bool {
	/*
		Method: Array
		Time complexity : O(n+m)
		Space complexity : O(1), в примере пишут, что O(1), но мне видеться как O(n), так как надо в буфер добавить n элементов.
	*/
	buffer := make([]int, 26)

	for _, c := range s {
		buffer[c-'a']++ // Получаем позицию символа в алфавите.
	}

	for _, c := range t {
		buffer[c-'a']-- // Получаем позицию символа в алфавите.
	}

	for _, v := range buffer {
		if v != 0 {
			return false
		}
	}

	return true
}
