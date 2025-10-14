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

	// Читаем количество элементов
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)

	// Читаем массив (даже если n=0)
	scanner.Scan()
	arrayStr := scanner.Text()

	// Читаем опорный элемент
	scanner.Scan()
	xStr := scanner.Text()
	x, _ := strconv.Atoi(xStr)

	// Если массив пустой
	if n == 0 {
		fmt.Println("0")
		fmt.Println("0")
		return
	}

	// Разбиваем строку на элементы
	elements := strings.Fields(arrayStr)

	// Считаем элементы меньшие x
	countLess := 0
	for _, elem := range elements {
		num, _ := strconv.Atoi(elem)
		if num < x {
			countLess++
		}
	}

	// Выводим результат
	fmt.Println(countLess)
	fmt.Println(n - countLess)
}
