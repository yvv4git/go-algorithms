package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Node структура для элемента связного списка
type Node struct {
	val  int
	next *Node
}

// RandomizedSet структура для хранения элементов
type RandomizedSet struct {
	head    *Node
	numsMap map[int]*Node
}

// Constructor для инициализации структуры
func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	return RandomizedSet{
		head:    nil,
		numsMap: make(map[int]*Node),
	}
}

// Insert метод для вставки элемента
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.numsMap[val]; exists {
		return false // Элемент уже существует
	}
	newNode := &Node{val: val, next: this.head}
	this.numsMap[val] = newNode
	this.head = newNode
	return true
}

// Remove метод для удаления элемента
func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.numsMap[val]; !exists {
		return false // Элемент не найден
	}
	node := this.numsMap[val]
	if node == this.head {
		this.head = this.head.next
	} else {
		prev := this.head
		for prev.next != node {
			prev = prev.next
		}
		prev.next = node.next
	}
	delete(this.numsMap, val)
	return true
}

// GetRandom метод для получения случайного элемента
func (this *RandomizedSet) GetRandom() int {
	count := 0
	current := this.head
	for current != nil {
		count++
		current = current.next
	}
	randomIndex := rand.Intn(count)
	current = this.head
	for i := 0; i < randomIndex; i++ {
		current = current.next
	}
	return current.val
}

// Пример использования
func main() {
	/*
		METHOD:

		TIME COMPLEXITY:

		SPACE COMPLEXITY:
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
