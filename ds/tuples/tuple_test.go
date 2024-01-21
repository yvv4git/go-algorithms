package main

import (
	"fmt"
	"testing"
)

// Tuple - это структура, которая эмулирует кортеж
type Tuple struct {
	First  int
	Second string
}

func TestTuple01(t *testing.T) {
	/*
		В Go нет встроенного типа данных, подобного кортежам в других языках программирования, таких как Python или JavaScript.
		Однако, вы можете использовать структуры или слайсы для эмуляции кортежей.
	*/
	// Создание экземпляра кортежа
	tuple := Tuple{First: 1, Second: "Hello"}

	// Вывод значений кортежа
	fmt.Println(tuple.First)  // Выведет: 1
	fmt.Println(tuple.Second) // Выведет: Hello
}
