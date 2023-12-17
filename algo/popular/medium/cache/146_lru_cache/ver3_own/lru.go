package ver3_own

import "sync"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	items    map[int]*Node
	head     *Node
	tail     *Node
	mutex    sync.RWMutex
}

// Constructor создает новый LRUCache с заданной емкостью.
// Сложность: O(1)
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		items:    make(map[int]*Node),
		head:     nil,
		tail:     nil,
	}
}

// Get возвращает значение, связанное с ключом, если ключ существует.
// Если ключ не существует, возвращается -1.
// Сложность: O(1)
func (c *LRUCache) Get(key int) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if node, ok := c.items[key]; ok {
		c.moveToFront(node)
		return node.value
	}
	return -1
}

// Put добавляет ключ и значение в кэш.
// Если кэш заполнен, удаляется последний использованный элемент.
// Сложность: O(1)
func (c *LRUCache) Put(key int, value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if node, ok := c.items[key]; ok {
		node.value = value
		c.moveToFront(node)
	} else {
		if len(c.items) == c.capacity {
			c.removeNode(c.tail)
		}
		node := &Node{key: key, value: value}
		c.addToFront(node)
		c.items[key] = node
	}
}

// moveToFront перемещает узел в начало списка.
// Сложность: O(1)
func (c *LRUCache) moveToFront(node *Node) {
	if node == c.head {
		return
	}

	node.prev.next = node.next
	if node == c.tail {
		c.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

	node.next = c.head
	node.prev = nil
	c.head.prev = node
	c.head = node
}

// addToFront добавляет узел в начало списка.
// Сложность: O(1)
func (c *LRUCache) addToFront(node *Node) {
	if c.head == nil {
		c.head = node
		c.tail = node
	} else {
		node.next = c.head
		c.head.prev = node
		c.head = node
	}
}

// removeNode удаляет узел из списка.
// Сложность: O(1)
func (c *LRUCache) removeNode(node *Node) {
	if node == c.head {
		c.head = node.next
	} else {
		node.prev.next = node.next
	}

	if node == c.tail {
		c.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

	delete(c.items, node.key)
}
