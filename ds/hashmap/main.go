package main

import (
	"fmt"
)

// Node представляет пару ключ-значение в хеш-таблице.
type Node struct {
	key   string
	value int
	next  *Node
}

// HashTable представляет простую реализацию хеш-таблицы.
type HashTable struct {
	buckets []*Node
	size    int
}

// HashCode возвращает простой хеш-код, который является суммой значений ASCII символов ключа.
// Временная сложность: O(n), где n - длина ключа.
func HashCode(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash
}

// Index возвращает индекс в хеш-таблице для заданного ключа.
// Временная сложность: O(1).
// Пространственная сложность: O(1).
func (ht *HashTable) Index(key string) int {
	return HashCode(key) % ht.size
}

// Insert добавляет пару ключ-значение в хеш-таблицу.
// Временная сложность: O(1) в среднем случае, O(n) в худшем случае, где n - количество коллизий.
// Пространственная сложность: O(1).
func (ht *HashTable) Insert(key string, value int) {
	index := ht.Index(key)
	node := &Node{key: key, value: value}
	if ht.buckets[index] == nil {
		ht.buckets[index] = node
	} else {
		current := ht.buckets[index]
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

// Get извлекает значение для заданного ключа из хеш-таблицы.
// Временная сложность: O(1) в среднем случае, O(n) в худшем случае, где n - количество коллизий.
// Пространственная сложность: O(1).
func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.Index(key)
	current := ht.buckets[index]
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

// NewHashTable создает новую хеш-таблицу заданного размера.
// Временная сложность: O(n), где n - размер хеш-таблицы.
// Пространственная сложность: O(n).
func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func main() {
	// Создаем хеш-таблицу размером 10
	hashTable := NewHashTable(10)

	// Добавляем пары ключ-значение
	hashTable.Insert("one", 1)
	hashTable.Insert("two", 2)
	hashTable.Insert("three", 3)

	// Получаем значение по ключу
	if value, ok := hashTable.Get("two"); ok {
		fmt.Println("Значение для ключа 'two':", value)
	} else {
		fmt.Println("Ключ 'two' не найден")
	}

	// Проверяем наличие ключа
	if _, ok := hashTable.Get("three"); ok {
		fmt.Println("Ключ 'three' существует в хеш-таблице")
	}

	// Удаление пары ключ-значение (в данном примере не реализовано)
}
