package remove_adjustment_duplicates

/*
Из строки надо удалить все дубликаты.
При чем не важно, как их крутить вертеть!
Например:
"abbaca" - ab и ba - это дубликаты.
*/
func removeDuplicates(s string) string {
	b := []byte(s)

	for i := len(b) - 1; i > 0; i-- {
		if i != len(b) && b[i] == b[i-1] {
			b = append(b[:i-1], b[i+1:]...)
		}
	}

	return string(b)
}
