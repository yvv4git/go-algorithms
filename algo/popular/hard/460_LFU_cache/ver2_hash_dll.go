//go:build ignore

package main

import (
	"container/list"
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
	kv       map[int]*cacheItem
	freqMap  map[int]*list.List
	minFreq  int
}

type cacheItem struct {
	key       int
	value     int
	frequency int
	element   *list.Element
}

func newLFUCacheImpl(capacity int) *lfuCacheImpl {
	return &lfuCacheImpl{
		capacity: capacity,
		kv:       make(map[int]*cacheItem),
		freqMap:  make(map[int]*list.List),
		minFreq:  0,
	}
}

// get внутренняя реализация Get - возвращает значение и обновляет частоту использования
func (impl *lfuCacheImpl) get(key int) int {
	if item, ok := impl.kv[key]; ok {
		// Увеличиваем частоту использования
		impl.incrementFrequency(item)
		return item.value
	}
	return -1
}

// put внутренняя реализация Put - добавляет/обновляет элемент, обрабатывает вытеснение
func (impl *lfuCacheImpl) put(key int, value int) {
	if impl.capacity == 0 {
		return
	}

	// Если ключ уже существует
	if item, ok := impl.kv[key]; ok {
		item.value = value
		impl.incrementFrequency(item)
		return
	}

	// Если достигнута емкость, вытесняем элемент
	if len(impl.kv) >= impl.capacity {
		// Получаем список элементов с минимальной частотой
		lruList := impl.freqMap[impl.minFreq]
		// Удаляем последний элемент (LRU)
		item := lruList.Back().Value.(*cacheItem)
		lruList.Remove(item.element)
		delete(impl.kv, item.key)
	}

	// Создаем новый элемент с частотой 1
	item := &cacheItem{
		key:       key,
		value:     value,
		frequency: 1,
	}

	// Добавляем в соответствующую частотную группу
	if impl.freqMap[1] == nil {
		impl.freqMap[1] = list.New()
	}
	item.element = impl.freqMap[1].PushFront(item)
	impl.kv[key] = item
	impl.minFreq = 1
}

// incrementFrequency увеличивает частоту использования элемента
func (impl *lfuCacheImpl) incrementFrequency(item *cacheItem) {
	// Удаляем элемент из текущей частотной группы
	lruList := impl.freqMap[item.frequency]
	lruList.Remove(item.element)

	// Обновляем минимальную частоту если нужно
	if item.frequency == impl.minFreq && lruList.Len() == 0 {
		impl.minFreq++
	}

	// Увеличиваем частоту
	item.frequency++

	// Добавляем в новую частотную группу
	if impl.freqMap[item.frequency] == nil {
		impl.freqMap[item.frequency] = list.New()
	}
	item.element = impl.freqMap[item.frequency].PushFront(item)
}

func main() {
	/*
		APPROACH: Хеш-таблица + Двусвязные списки
		Используется комбинация хеш-таблицы для хранения пар ключ-значение и двусвязных списков для организации элементов по частоте использования.
		Элементы группируются по частоте, внутри групп поддерживается порядок LRU.
		При переполнении сначала вытесняются элементы с наименьшей частотой, а при равной частоте - самый давно использованный.

		TIME COMPLEXITY: O(1) среднее время для get/put благодаря хеш-таблице
		SPACE COMPLEXITY: O(n) где n - емкость кэша
	*/
	fmt.Println("Демонстрация работы LFU кэша (реализация на двусвязных списках)")
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
