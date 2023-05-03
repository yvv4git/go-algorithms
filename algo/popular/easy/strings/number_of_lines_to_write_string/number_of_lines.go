package number_of_lines_to_write_string

func numberOfLines(widths []int, s string) []int {
	/*
		Time complexity : O(n), where N is length of 's'.
		Space complexity : O(n)
	*/
	const maxLength = 100

	pixels := maxLength
	rows := 1

	for _, ch := range s {
		pixels -= widths[ch-'a']

		if pixels < 0 {
			pixels = maxLength - widths[ch-'a']
			rows++
		}
		//fmt.Printf("ch[%s]=%v ch-a=%v pixels=%v \n", string(ch), ch, ch-'a', pixels)
	}

	return []int{
		rows,
		maxLength - pixels, // из 100(допустимая длина строки) вычитаем оставшееся свободное место, получаем занятое место
	}
}
