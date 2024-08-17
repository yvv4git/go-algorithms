package main

import "strconv"

// Специальные последовательности байтов
const oneByte = 0                 // 0000 0000
const continuation = 1 << 7       // 1000 0000
const twoByte = 1<<7 | 1<<6       // 1100 0000
const threeByte = twoByte | 1<<5  // 1110 0000
const fourByte = threeByte | 1<<4 // 1111 0000

// Состояния
const (
	Empty int = iota + 1
	InTwo
	InThree
	InFour
)

func validUtf8V2(data []int) bool {
	/*
		METHOD: Bitwise / Побитовый анализ

		TIME COMPLEXITY: O(n), где n — количество байтов в массиве data.
		Мы проходим по каждому байту ровно один раз.

		SPACE COMPLEXITY: O(1), так как используется только константное количество дополнительной памяти (переменные state и count).
	*/
	state := Empty
	count := 0

	for i := 0; i < len(data); i++ {
		switch state {
		case Empty:
			// Однобайтовый символ
			if (0^data[i])>>7 == 0 {
				continue
			}

			// Проверка на продолжающий байт
			if (continuation^data[i])>>6 == 0 {
				return false
			}

			// Проверка на двухбайтовый символ
			if (twoByte^data[i])>>5 == 0 {
				state = InTwo
				count = 1
				continue
			}

			// Проверка на трехбайтовый символ
			if (threeByte^data[i])>>4 == 0 {
				state = InThree
				count = 2
				continue
			}

			// Проверка на четырехбайтовый символ
			if (fourByte^data[i])>>3 == 0 {
				state = InFour
				count = 3
				continue
			}

			// Что-то еще, некорректный байт
			return false
		case InTwo, InThree, InFour:
			// Проверка на продолжающий байт
			if (continuation^data[i])>>6 != 0 {
				return false
			}

			count--

			if count == 0 {
				state = Empty
			}
		default:
			panic("необработанное состояние: " + strconv.Itoa(state))
		}
	}

	return count == 0
}
