package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomizedSet структура для хранения элементов
type RandomizedSet struct {
	numsMap  map[int]int // Хеш-таблица для хранения значений и их индексов
	numsList []int       // Список для хранения значений
}

// Constructor для инициализации структуры
func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	return RandomizedSet{
		numsMap:  make(map[int]int),
		numsList: []int{},
	}
}

// Insert метод для вставки элемента
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.numsMap[val]; exists {
		return false // Элемент уже существует
	}
	this.numsList = append(this.numsList, val) // Добавляем элемент в конец списка
	this.numsMap[val] = len(this.numsList) - 1 // Сохраняем индекс элемента в хеш-таблице
	return true
}

// Remove метод для удаления элемента
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.numsMap[val]
	if !exists {
		return false // Элемент не найден
	}
	last := this.numsList[len(this.numsList)-1]          // Получаем последний элемент в списке
	this.numsList[index] = last                          // Заменяем удаляемый элемент последним элементом
	this.numsMap[last] = index                           // Обновляем индекс последнего элемента в хеш-таблице
	this.numsList = this.numsList[:len(this.numsList)-1] // Удаляем последний элемент из списка
	delete(this.numsMap, val)                            // Удаляем элемент из хеш-таблицы
	return true
}

// GetRandom метод для получения случайного элемента
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (this *RandomizedSet) GetRandom() int {
	return this.numsList[rand.Intn(len(this.numsList))] // Возвращаем случайный элемент из списка
}

// Пример использования
func main() {
	/*
		METHOD: Map / Hash table
		Выбран подход с использованием хеш-таблицы (map) и списка (slice) по нескольким причинам, которые обеспечивают эффективную реализацию требуемых операций:
		1. Быстрая вставка и удаление:
		Хеш-таблица позволяет быстро проверять наличие элемента и добавлять новые элементы за O(1) в среднем случае
		При удалении элемента, хеш-таблица позволяет быстро найти индекс элемента в списке, что необходимо для его удаления.

		2. Быстрое получение случайного элемента:
		Список (slice) позволяет быстро получать случайный элемент за O(1),
		так как доступ к элементу по индексу в списке является константной операцией.

		3. Оптимизация удаления:
		Для удаления элемента из списка за O(1) используется трюк с заменой удаляемого элемента последним элементом в списке
		и последующим удалением последнего элемента. Это позволяет избежать дорогостоящей операции перемещения элементов в списке.

		4. Баланс между временной и пространственной сложностью:
		Использование хеш-таблицы и списка обеспечивает баланс между временной сложностью операций и использованием памяти.
		Хотя пространственная сложность составляет O(n), где n — количество элементов, это является разумной ценой за обеспечение требуемой временной сложности операций.

		TIME COMPLEXITY:
		- Insert: O(1)
		- Remove: O(1)
		- GetRandom: O(1)

		SPACE COMPLEXITY:
		- O(n), где n — количество элементов в структуре. Хранение элементов в списке и хеш-таблице.
	*/
	obj := Constructor()
	fmt.Println(obj.Insert(1))   // true
	fmt.Println(obj.Remove(2))   // false
	fmt.Println(obj.Insert(2))   // true
	fmt.Println(obj.GetRandom()) // 1 или 2
	fmt.Println(obj.Remove(1))   // true
	fmt.Println(obj.Insert(2))   // false
	fmt.Println(obj.GetRandom()) // 2
}
