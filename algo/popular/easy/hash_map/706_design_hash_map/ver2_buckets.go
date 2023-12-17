package _06_design_hash_map

/*
Бакеты - это метод хеширования, который используется для решения коллизий.
Вместо хеширования всего ключа, мы разделяем ключ на несколько частей и хешируем каждую часть.
Это позволяет избежать коллизий и уменьшить количество коллизий.
*/

type Node struct {
	key, value int
	next       *Node
}

type MyHashMap2 struct {
	buckets []*Node
	size    int
}

func Constructor2() MyHashMap2 {
	size := 1000
	buckets := make([]*Node, size)
	return MyHashMap2{buckets: buckets, size: size}
}

func (m *MyHashMap2) hash(key int) int {
	return key % m.size
}

func (m *MyHashMap2) Put(key int, value int) {
	i := m.hash(key)
	node := m.buckets[i]
	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		node = node.next
	}
	newNode := &Node{key: key, value: value, next: m.buckets[i]}
	m.buckets[i] = newNode
}

func (m *MyHashMap2) Get(key int) int {
	i := m.hash(key)
	node := m.buckets[i]
	for node != nil {
		if node.key == key {
			return node.value
		}
		node = node.next
	}
	return -1
}

func (m *MyHashMap2) Remove(key int) {
	i := m.hash(key)
	node := m.buckets[i]
	if node == nil {
		return
	}
	if node.key == key {
		m.buckets[i] = node.next
		return
	}
	prev := node
	node = node.next
	for node != nil {
		if node.key == key {
			prev.next = node.next
			return
		}
		prev = node
		node = node.next
	}
}
