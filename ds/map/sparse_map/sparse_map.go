package main

import (
	"fmt"
)

// SparseMap представляет собой структуру данных для хранения пар "ключ-значение"
// с целыми числами в качестве ключей.
type SparseMap struct {
	data map[int]interface{}
}

// NewSparseMap создает и возвращает новый экземпляр SparseMap.
func NewSparseMap() *SparseMap {
	return &SparseMap{
		data: make(map[int]interface{}),
	}
}

// Put добавляет или обновляет значение по заданному ключу.
func (sm *SparseMap) Put(key int, value interface{}) {
	sm.data[key] = value
}

// Get возвращает значение по заданному ключу.
// Если ключ не найден, возвращается nil.
func (sm *SparseMap) Get(key int) interface{} {
	return sm.data[key]
}

// Remove удаляет значение по заданному ключу.
func (sm *SparseMap) Remove(key int) {
	delete(sm.data, key)
}

// Contains проверяет, существует ли ключ в SparseMap.
func (sm *SparseMap) Contains(key int) bool {
	_, exists := sm.data[key]
	return exists
}

// Size возвращает количество элементов в SparseMap.
func (sm *SparseMap) Size() int {
	return len(sm.data)
}

// Keys возвращает слайс всех ключей в SparseMap.
func (sm *SparseMap) Keys() []int {
	keys := make([]int, 0, len(sm.data))
	for key := range sm.data {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	// Создаем новый SparseMap
	sparseMap := NewSparseMap()

	// Добавляем элементы
	sparseMap.Put(1, "One")
	sparseMap.Put(100, "One Hundred")
	sparseMap.Put(2000, "Two Thousand")

	// Получаем значение по ключу
	fmt.Println(sparseMap.Get(100)) // Вывод: One Hundred

	// Проверяем наличие ключа
	fmt.Println(sparseMap.Contains(2)) // Вывод: false

	// Удаляем элемент
	sparseMap.Remove(100)
	fmt.Println(sparseMap.Contains(100)) // Вывод: false

	// Получаем все ключи
	fmt.Println(sparseMap.Keys()) // Вывод: [1 2000]

	// Получаем размер SparseMap
	fmt.Println(sparseMap.Size()) // Вывод: 2
}
