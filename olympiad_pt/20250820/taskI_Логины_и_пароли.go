//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	strings := make([]string, 2*n)
	for i := 0; i < 2*n; i++ {
		fmt.Fscan(in, &strings[i])
	}

	// Создаем слайс структур для хранения строк и их индексов
	type StringWithIndex struct {
		str   string
		index int
	}

	strs := make([]StringWithIndex, 2*n)
	for i := 0; i < 2*n; i++ {
		strs[i] = StringWithIndex{strings[i], i + 1}
	}

	// Сортируем строки по длине (от коротких к длинным)
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i].str) < len(strs[j].str)
	})

	// Создаем мапу для отслеживания использованных индексов
	used := make(map[int]bool)

	// Результат: пары (логин, пароль)
	result := make([][2]int, 0, n)

	// Проходим по всем строкам от самых коротких к самым длинным
	for i := 0; i < 2*n; i++ {
		if used[strs[i].index] {
			continue
		}

		login := strs[i].str
		loginIndex := strs[i].index
		used[loginIndex] = true

		// Ищем пароль: строку, которая начинается с логина
		for j := i + 1; j < 2*n; j++ {
			if used[strs[j].index] {
				continue
			}

			password := strs[j].str
			// Проверяем, начинается ли пароль с логина
			if len(password) >= len(login) && password[:len(login)] == login {
				result = append(result, [2]int{loginIndex, strs[j].index})
				used[strs[j].index] = true
				break
			}
		}
	}

	// Выводим результат
	for _, pair := range result {
		fmt.Fprintln(out, pair[0], pair[1])
	}
}
