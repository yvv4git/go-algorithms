package main

import "fmt"

func main() {
	/*
		Вставка в начало - O(1).
		Вставка в конец - O(n).
		Вставка в середину - O(n).
		Разворот односвязного списка - O(n). Так как надо пройти по всем узлам и поменять Next на предыдущий узел.
	*/
	list := &List{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	fmt.Println("Исходный список:")
	list.Print()

	list.Reverse()

	fmt.Println("Развернутый список:")
	list.Print()
}
