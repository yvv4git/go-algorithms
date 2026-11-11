package main

import (
	"bufio"
	"fmt"
	"os"
)

func correctWithRepetition(mode int, s string) string {
	/*
		TASK: Восстановить исходную строку x после не более одной ошибки в закодированной строке y.
		Программа запускается дважды, поэтому mode = 1 означает кодирование, а mode = 2 - декодирование.
		При mode = 1 функция получает x и возвращает y, в которой каждый бит повторён три раза.
		Пример: x = "01010", тогда y = "000111000111000".
		При mode = 2 функция получает z, отличающуюся от y не более чем в одном бите, и возвращает x.
		Пример: z = "000101000111000" (испорчена позиция 5), тогда ответ x = "01010".
		Границы: x = "" даёт y = "", а x = "1" даёт y = "111".
		Если z отличается от y в двух и более позициях, условие задачи нарушено и корректность не гарантируется.
		При mode, не равном 1 или 2, функция возвращает пустую строку.

		METHOD: Repetition code (троирование битов)
		Строка x кодируется так: 0 -> 000, а 1 -> 111.
		Декодер делит z на тройки и выбирает в каждой тройке бит, который встречается чаще.
		Одна ошибка меняет только одну тройку, а два других бита в ней остаются равными исходному.
		Поэтому голосование большинства всегда возвращает правильный бит.

		TIME COMPLEXITY: O(n)
		Кодирование и декодирование проходят по строке один раз.

		SPACE COMPLEXITY: O(n)
		При кодировании строится строка длины 3n, при декодировании - строка длины n.
	*/
	switch mode {
	case 1:
		return repetitionEncode(s)
	case 2:
		return repetitionDecode(s)
	default:
		return ""
	}
}

func repetitionEncode(x string) string {
	y := make([]byte, 0, 3*len(x))
	for i := 0; i < len(x); i++ {
		for k := 0; k < 3; k++ {
			y = append(y, x[i])
		}
	}
	return string(y)
}

func repetitionDecode(z string) string {
	x := make([]byte, 0, len(z)/3)
	for i := 0; i+2 < len(z); i += 3 {
		sum := int(z[i]-'0') + int(z[i+1]-'0') + int(z[i+2]-'0')
		if sum >= 2 {
			x = append(x, '1')
		} else {
			x = append(x, '0')
		}
	}
	return string(x)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var mode int
	var s string
	if _, err := fmt.Fscan(reader, &mode, &s); err != nil {
		return
	}

	fmt.Fprintln(writer, correctWithRepetition(mode, s))
}
