package number

import (
	"math/rand"
	"testing"
	"time"
)

/*

Guessing game / Игра в угадайку.

Сложность: Лёгкая

Условие задачи: играем в угадайку по следующей схеме:

Выбирается число от 1 до n. Надо отгадать загаданное число. После каждой неудачной попытки говорится больше или меньше заданное число.

Надо реализовать API:

-1: загаданное число больше выбранного;
1: загаданное число меньше выбранного;
0: загаданное число и выбранное совпали.

Необходимо вернуть загаданное число.

Пример:

Ввод:  n = 10, pick = 6
Вывод: 6

*/

func guessNumber(n int) int {
	res := randomGuess(n)

	low, high := 1, n

	for low <= high {
		mid := low + (high-low)/2
		if mid == res {
			return mid
		} else if res < mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0
}

func randomGuess(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-1) + 1
}

// TEST
func TestGuessNumber(t *testing.T) {
	result := guessNumber(10)
	t.Logf("result: %d", result)
}
