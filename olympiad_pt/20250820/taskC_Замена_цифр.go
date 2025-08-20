//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(firstLine[0])
	k, _ := strconv.Atoi(firstLine[1])

	scanner.Scan()
	numbers := strings.Fields(scanner.Text())

	// Преобразуем строки в числа
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(numbers[i])
	}

	// Создаем слайс для хранения информации о потенциальной выгоде от замен
	type change struct {
		value    int // величина выгоды
		position int // позиция в числе
		number   int // индекс числа в массиве
	}

	var changes []change

	// Для каждого числа вычисляем возможную выгоду от замен каждой цифры
	for idx, num := range arr {
		// Преобразуем число в строку для работы с цифрами
		s := strconv.Itoa(num)
		length := len(s)

		// Для каждой позиции в числе вычисляем, сколько мы можем выиграть,
		// заменив цифру на 9
		for pos, char := range s {
			digit := int(char - '0')
			gain := (9 - digit) * pow10(length-pos-1)
			changes = append(changes, change{gain, pos, idx})
		}
	}

	// Сортируем изменения по убыванию выгоды
	// Используем простую пузырьковую сортировку, так как n небольшое
	for i := 0; i < len(changes)-1; i++ {
		for j := i + 1; j < len(changes); j++ {
			if changes[i].value < changes[j].value {
				changes[i], changes[j] = changes[j], changes[i]
			}
		}
	}

	// Применяем до k самых выгодных замен
	totalChanges := min(k, len(changes))
	for i := 0; i < totalChanges; i++ {
		if changes[i].value > 0 { // Замена имеет смысл только если выгода положительная
			arr[changes[i].number] += changes[i].value
		}
	}

	// Считаем итоговую сумму
	sum := 0
	for _, num := range arr {
		sum += num
	}

	fmt.Println(sum)
}

func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
