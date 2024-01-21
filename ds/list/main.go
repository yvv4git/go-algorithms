package main

// importing fmt and container list packages
import (
	"container/list" // Двусвязный список
	"fmt"
)

// main method
func main() {
	/*
		В Go, структура list.List представляет двусвязный список. Этот список реализует контейнер,
		который содержит элементы в некотором порядке. Вот основные методы этой структуры и их сложности:
	*/
	var intList list.List
	intList.PushBack(11) // O(1) - помещает элемент e в конец списка
	intList.PushBack(23)
	intList.PushBack(34)

	// Идем по списку с начала.
	// Front() - O(1)
	// Next() - O(1)
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
