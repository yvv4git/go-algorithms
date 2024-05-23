package main

import (
	"fmt"
)

// toLowerCase переводит все символы строки в нижний регистр без использования стандартной библиотеки.
func toLowerCase(str string) string {
	/*
		METHOD: Iterate
		В Go, то можно реализовать функцию toLowerCase самостоятельно, используя ASCII-значения символов.
		В таблице ASCII буквы верхнего регистра находятся в диапазоне от 65 до 90, а буквы нижнего регистра - от 97 до 122.
		Разница между значениями верхнего и нижнего регистров равна 32. Поэтому, чтобы перевести букву в нижний регистр,
		нужно к ее ASCII-коду прибавить 32.

		TIME COMPLEXITY: O(n), где n - количество символов в строке, так как мы проходим по каждому символу строки.

		SPACE COMPLEXITY: O(n), так как мы создаем новую строку той же длины.
	*/
	result := make([]rune, len(str))
	for i, char := range str {
		if char >= 'A' && char <= 'Z' {
			result[i] = char + 32 // Перевод символа в нижний регистр
		} else {
			result[i] = char // Оставить символ без изменений, если он не является буквой верхнего регистра
		}
	}

	return string(result)
}

func main() {
	str := "Hello, World!"
	fmt.Println(toLowerCase(str)) // Выведет: "hello, world!"
}
