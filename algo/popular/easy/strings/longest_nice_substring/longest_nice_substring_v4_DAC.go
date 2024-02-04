package longest_nice_substring

import "unicode"

func longestNiceSubstringV4(s string) string {
	/*
		METHOD: Divide And Conquer
		Time complexity: O(n)
		Space complexity: O(n)

		Разделяй и властвуй в информатике — схема разработки алгоритмов, заключающаяся в рекурсивном разбиении решаемой
		задачи на две или более подзадачи того же типа, но меньшего размера, и комбинировании их решений для получения
		ответа к исходной задаче; разбиения выполняются до тех пор, пока все подзадачи не окажутся элементарными.
		Это мощный инструмент для решения концептуально сложных задач: все, что требуется для этого,
		это найти случай разбивания задачи на подзадачи, решения тривиальных случаев и объединения подзадач в исходную задачу.
	*/
	if len(s) < 2 {
		return ""
	}

	runesHash := make(map[rune]struct{}, 0)
	for _, r := range s {
		runesHash[r] = struct{}{}
	}

	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		_, existUpper := runesHash[unicode.ToUpper(r)]
		_, existLower := runesHash[unicode.ToLower(r)]
		if existUpper && existLower {
			continue
		}

		left := longestNiceSubstringV4(s[:i])
		right := longestNiceSubstringV4(s[i+1:])
		if len(left) >= len(right) {
			return left
		} else {
			return right
		}
	}

	return s
}
