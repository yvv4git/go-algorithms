package main

import (
	"container/list"
	"fmt"
	"sync"
)

// entry представляет собой элемент в кэше
type entry struct {
	key   interface{}
	value interface{}
}

// LRUCache представляет собой LRU кэш
type LRUCache struct {
	capacity int                           // емкость кэша
	list     *list.List                    // двусвязный список для управления порядком элементов
	items    map[interface{}]*list.Element // map для быстрого доступа к элементам по ключу
	mu       sync.Mutex                    // мьютекс для потокобезопасности
}

// NewLRUCache создает новый LRU кэш с заданной емкостью
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		items:    make(map[interface{}]*list.Element),
	}
}

// Get получает значение по ключу из кэша
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.list.MoveToFront(elem) // перемещаем элемент в начало списка
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// Put добавляет значение в кэш с заданным ключом
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (c *LRUCache) Put(key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.list.MoveToFront(elem)          // перемещаем элемент в начало списка
		elem.Value.(*entry).value = value // обновляем значение
	} else {
		elem := c.list.PushFront(&entry{key, value}) // добавляем новый элемент в начало списка
		c.items[key] = elem                          // добавляем элемент в map
	}

	if c.list.Len() > c.capacity {
		c.removeOldest() // удаляем старейший элемент, если превышен размер кэша
	}
}

// removeOldest удаляет старейший (наименее используемый) элемент из кэша
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (c *LRUCache) removeOldest() {
	ent := c.list.Back() // получаем последний элемент из списка
	if ent != nil {
		c.removeElement(ent) // удаляем элемент
	}
}

// removeElement удаляет данный элемент из кэша
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (c *LRUCache) removeElement(e *list.Element) {
	c.list.Remove(e) // удаляем элемент из списка
	kv := e.Value.(*entry)
	delete(c.items, kv.key) // удаляем элемент из map
}

func main() {
	cache := NewLRUCache(2)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Получаем значение по ключу
	if value, ok := cache.Get("key1"); ok {
		fmt.Println("Value for key1:", value)
	}

	// Добавляем новое значение, вытесняя старое
	cache.Put("key3", "value3")

	// Проверяем, что старое значение удалено
	if _, ok := cache.Get("key2"); !ok {
		fmt.Println("Key2 was removed from cache")
	}
}
