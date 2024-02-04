package ver1

import "container/list"

// LRUCache - структура для хранения размера, списка и хэш-таблицы
type LRUCache struct {
	Capacity int
	List     *list.List
	Items    map[int]*list.Element
}

// Constructor - функция для создания нового LRUCache
// TIME COMPLEXITY: O(1)
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		List:     new(list.List),
		Items:    make(map[int]*list.Element, capacity),
	}
}

// Get - функция для получения значения по ключу
// TIME COMPLEXITY: O(1)
func (c *LRUCache) Get(key int) int {
	if node, ok := c.Items[key]; ok {
		value := node.Value.(*list.Element).Value.(Node).Value
		c.List.MoveToFront(node)
		return value
	}
	return -1
}

// Put - функция для добавления пары ключ-значение
// TIME COMPLEXITY: O(1)
func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.Items[key]; ok {
		c.List.MoveToFront(node)
		node.Value.(*list.Element).Value = Node{Key: key, Value: value}
	} else {
		if c.List.Len() == c.Capacity {
			idx := c.List.Back().Value.(*list.Element).Value.(Node).Key
			delete(c.Items, idx)
			c.List.Remove(c.List.Back())
		}
		node := &list.Element{
			Value: Node{
				Key:   key,
				Value: value,
			},
		}
		newNode := c.List.PushFront(node)
		c.Items[key] = newNode
	}
}
