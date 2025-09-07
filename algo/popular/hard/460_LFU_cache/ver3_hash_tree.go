//go:build ignore

package main

import (
	"fmt"
)

// Public interface with original signatures
type LFUCache struct {
	impl *lfuCacheImpl
}

// Constructor создает новый LFU кэш с указанной емкостью
func Constructor(capacity int) LFUCache {
	return LFUCache{
		impl: newLFUCacheImpl(capacity),
	}
}

// Get возвращает значение по ключу, увеличивает счетчик использования
func (this *LFUCache) Get(key int) int {
	return this.impl.get(key)
}

// Put добавляет или обновляет пару ключ-значение, обновляет счетчики
func (this *LFUCache) Put(key int, value int) {
	this.impl.put(key, value)
}

// Implementation details
type lfuCacheImpl struct {
	capacity int
	kv       map[int]*treeNode
	root     *treeNode
	counter  int
}

type treeNode struct {
	key       int
	value     int
	frequency int
	stamp     int
	left      *treeNode
	right     *treeNode
}

func newLFUCacheImpl(capacity int) *lfuCacheImpl {
	return &lfuCacheImpl{
		capacity: capacity,
		kv:       make(map[int]*treeNode),
		root:     nil,
		counter:  0,
	}
}

// get внутренняя реализация Get - возвращает значение и обновляет частоту использования
func (impl *lfuCacheImpl) get(key int) int {
	if node, ok := impl.kv[key]; ok {
		impl.counter++
		// Удаляем узел из дерева перед обновлением частоты
		impl.root = impl.deleteNode(impl.root, node)
		// Обновляем частоту и метку времени
		node.frequency++
		node.stamp = impl.counter
		// Вставляем обратно с новыми параметрами
		impl.root = impl.insertNode(impl.root, node)
		return node.value
	}
	return -1
}

// put внутренняя реализация Put - добавляет/обновляет элемент, обрабатывает вытеснение
func (impl *lfuCacheImpl) put(key int, value int) {
	if impl.capacity == 0 {
		return
	}

	impl.counter++

	// Если ключ уже существует
	if node, ok := impl.kv[key]; ok {
		// Удаляем узел из дерева перед обновлением
		impl.root = impl.deleteNode(impl.root, node)
		// Обновляем значение и частоту
		node.value = value
		node.frequency++
		node.stamp = impl.counter
		// Вставляем обратно с новыми параметрами
		impl.root = impl.insertNode(impl.root, node)
		return
	}

	// Если достигнута емкость, вытесняем элемент
	if len(impl.kv) >= impl.capacity {
		// Находим минимальный узел (наименьшая частота, самый старый)
		minNode := impl.findMinNode(impl.root)
		// Удаляем из дерева и мапы
		impl.root = impl.deleteNode(impl.root, minNode)
		delete(impl.kv, minNode.key)
	}

	// Создаем новый узел
	node := &treeNode{
		key:       key,
		value:     value,
		frequency: 1,
		stamp:     impl.counter,
	}
	// Добавляем в мапу и дерево
	impl.kv[key] = node
	impl.root = impl.insertNode(impl.root, node)
}

// insertNode вставляет узел в дерево с учетом частоты и времени
func (impl *lfuCacheImpl) insertNode(root, node *treeNode) *treeNode {
	if root == nil {
		return node
	}

	// Сравниваем сначала частоту, затем время
	if node.frequency < root.frequency ||
		(node.frequency == root.frequency && node.stamp < root.stamp) {
		root.left = impl.insertNode(root.left, node)
	} else {
		root.right = impl.insertNode(root.right, node)
	}

	return root
}

// deleteNode удаляет узел из дерева
func (impl *lfuCacheImpl) deleteNode(root, node *treeNode) *treeNode {
	if root == nil {
		return nil
	}

	if node == root {
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}
		// У узла есть оба потомка
		minRight := impl.findMinNode(root.right)
		root.key = minRight.key
		root.value = minRight.value
		root.frequency = minRight.frequency
		root.stamp = minRight.stamp
		root.right = impl.deleteNode(root.right, minRight)
		return root
	}

	if node.frequency < root.frequency ||
		(node.frequency == root.frequency && node.stamp < root.stamp) {
		root.left = impl.deleteNode(root.left, node)
	} else {
		root.right = impl.deleteNode(root.right, node)
	}

	return root
}

// findMinNode находит узел с минимальной частотой (и самым старым временем)
func (impl *lfuCacheImpl) findMinNode(root *treeNode) *treeNode {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current
}

func main() {
	/*
		APPROACH: Hash table + Binary search tree
		Используется комбинация хеш-таблицы для хранения пар ключ-значение и бинарного дерева поиска для организации элементов по частоте использования и времени доступа.
		Дерево упорядочено сначала по частоте, затем по времени последнего использования. При переполнении вытесняется элемент с наименьшей частотой, а при равной частоте - самый старый.

		TIME COMPLEXITY: O(log n) время для get/put из-за операций с деревом
		SPACE COMPLEXITY: O(n) где n - емкость кэша
	*/
	fmt.Println("Демонстрация работы LFU кэша (реализация на дереве)")
	cache := Constructor(2)

	// Добавляем элементы
	cache.Put(1, 1)
	fmt.Println("Put(1,1) -> cache: [{1:1}]")

	cache.Put(2, 2)
	fmt.Println("Put(2,2) -> cache: [{1:1}, {2:2}]")

	// Получаем элемент
	val := cache.Get(1)
	fmt.Printf("Get(1) -> %d (ожидается 1)\n", val)

	// Добавляем новый элемент, должен вытеснить 2
	cache.Put(3, 3)
	fmt.Println("Put(3,3) -> cache: [{1:1}, {3:3}] (2 вытеснен как наименее используемый)")

	// Проверяем вытесненный элемент
	val = cache.Get(2)
	fmt.Printf("Get(2) -> %d (ожидается -1, так как 2 вытеснен)\n", val)

	// Проверяем новый элемент
	val = cache.Get(3)
	fmt.Printf("Get(3) -> %d (ожидается 3)\n", val)

	// Добавляем еще один элемент, должен вытеснить 1
	cache.Put(4, 4)
	fmt.Println("Put(4,4) -> cache: [{3:3}, {4:4}] (1 вытеснен как наименее используемый)")

	// Проверяем вытесненный элемент
	val = cache.Get(1)
	fmt.Printf("Get(1) -> %d (ожидается -1, так как 1 вытеснен)\n", val)

	// Проверяем оставшиеся элементы
	val = cache.Get(3)
	fmt.Printf("Get(3) -> %d (ожидается 3)\n", val)

	val = cache.Get(4)
	fmt.Printf("Get(4) -> %d (ожидается 4)\n", val)
}
