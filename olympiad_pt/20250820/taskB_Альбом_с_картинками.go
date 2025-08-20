//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var w, h int
	var n, a, b int

	// Чтение размеров страницы
	fmt.Scan(&w, &h)

	// Чтение количества картин и их размеров
	fmt.Scan(&n, &a, &b)

	// Проверяем, помещается ли хотя бы одна картина на страницу
	if a > w || b > h {
		fmt.Println(-1)
		return
	}

	// Вычисляем максимальное количество картин на одной странице
	// по горизонтали и вертикали
	horizontal := w / a
	vertical := h / b

	// Общее количество картин на одной странице
	perPage := horizontal * vertical

	// Если на страницу не помещается ни одна картина
	if perPage == 0 {
		fmt.Println(-1)
		return
	}

	// Вычисляем минимальное количество страниц
	pages := (n + perPage - 1) / perPage // Эквивалентно math.Ceil(n / perPage)

	fmt.Println(pages)
}
