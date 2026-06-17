package main

import "fmt"

const base = 256
const prime = 101

func RabinKarpSearch(text, pattern string) []int {
	/*
		METHOD: Алгоритм Рабина-Карпа (Rabin-Karp)
		TIME COMPLEXITY:
		  - Средний случай: O(n + m)
		  - Худший случай: O(n × m) (при коллизиях хеша)
		SPACE COMPLEXITY: O(1) без учёта массива результатов

		Основан на вычислении хеша для паттерна и для каждой подстроки текста
		с использованием кольцевого (rolling) хеша. При совпадении хешей выполняется
		проверка посимвольно для исключения коллизии.
	*/
	if len(pattern) == 0 {
		return []int{}
	}

	n := len(text)
	m := len(pattern)

	if m > n {
		return []int{}
	}

	patHash := 0
	textHash := 0
	h := 1
	var matches []int

	for i := 0; i < m-1; i++ {
		h = (h * base) % prime
	}

	for i := 0; i < m; i++ {
		patHash = (patHash*base + int(pattern[i])) % prime
		textHash = (textHash*base + int(text[i])) % prime
	}

	for i := 0; i <= n-m; i++ {
		if patHash == textHash {
			j := 0
			for j < m && text[i+j] == pattern[j] {
				j++
			}
			if j == m {
				matches = append(matches, i)
			}
		}

		if i < n-m {
			textHash = (textHash - int(text[i])*h) % prime
			if textHash < 0 {
				textHash += prime
			}
			textHash = (textHash*base + int(text[i+m])) % prime
		}
	}

	return matches
}

func main() {
	/*
		Пример использования алгоритма Рабина-Карпа для поиска подстроки в тексте.
	*/
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	positions := RabinKarpSearch(text, pattern)

	fmt.Printf("Текст:    %s\n", text)
	fmt.Printf("Паттерн:  %s\n", pattern)
	if len(positions) > 0 {
		fmt.Printf("Найден на позициях: %v\n", positions)
		for _, pos := range positions {
			fmt.Printf("  %s\n", text[:pos]+"["+pattern+"]"+text[pos+len(pattern):])
		}
	} else {
		fmt.Println("Паттерн не найден в тексте.")
	}
}
