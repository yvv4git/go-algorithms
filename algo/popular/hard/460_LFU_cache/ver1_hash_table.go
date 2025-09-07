package main

import (
	"container/heap"
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
	kv       map[int]int
	kf       map[int]int
	fh       *frequencyHeap
	counter  int
	keyStamp map[int]int
}

type item struct {
	key       int
	frequency int
	stamp     int
	index     int
}

type frequencyHeap []*item

func (h frequencyHeap) Len() int { return len(h) }
func (h frequencyHeap) Less(i, j int) bool {
	if h[i].frequency == h[j].frequency {
		return h[i].stamp < h[j].stamp
	}
	return h[i].frequency < h[j].frequency
}
func (h frequencyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *frequencyHeap) Push(x interface{}) {
	n := len(*h)
	it := x.(*item)
	it.index = n
	*h = append(*h, it)
}

func (h *frequencyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	it := old[n-1]
	it.index = -1
	*h = old[0 : n-1]
	return it
}

func newLFUCacheImpl(capacity int) *lfuCacheImpl {
	return &lfuCacheImpl{
		capacity: capacity,
		kv:       make(map[int]int),
		kf:       make(map[int]int),
		fh:       &frequencyHeap{},
		keyStamp: make(map[int]int),
	}
}

// get внутренняя реализация Get - возвращает значение и обновляет частоту использования
func (impl *lfuCacheImpl) get(key int) int {
	if val, ok := impl.kv[key]; ok {
		impl.kf[key]++
		impl.counter++
		impl.keyStamp[key] = impl.counter
		for _, it := range *impl.fh {
			if it.key == key {
				it.frequency++
				it.stamp = impl.counter
				heap.Fix(impl.fh, it.index)
				break
			}
		}
		return val
	}
	return -1
}

// put внутренняя реализация Put - добавляет/обновляет элемент, обрабатывает вытеснение
func (impl *lfuCacheImpl) put(key int, value int) {
	if impl.capacity == 0 {
		return
	}

	impl.counter++

	if _, ok := impl.kv[key]; ok {
		impl.kv[key] = value
		impl.kf[key]++
		impl.keyStamp[key] = impl.counter
		for _, it := range *impl.fh {
			if it.key == key {
				it.frequency++
				it.stamp = impl.counter
				heap.Fix(impl.fh, it.index)
				break
			}
		}
		return
	}

	if len(impl.kv) >= impl.capacity {
		it := heap.Pop(impl.fh).(*item)
		delete(impl.kv, it.key)
		delete(impl.kf, it.key)
		delete(impl.keyStamp, it.key)
	}

	impl.kv[key] = value
	impl.kf[key] = 1
	impl.keyStamp[key] = impl.counter
	heap.Push(impl.fh, &item{
		key:       key,
		frequency: 1,
		stamp:     impl.counter,
	})
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	/*
		APPROACH: HASH TABLE + MIN-HEAP
		Используется комбинация хеш-таблицы для хранения пар ключ-значение с доступом O(1) и минимальной кучи для отслеживания частоты использования элементов.
		При переполнении кэша вытесняется элемент с наименьшей частотой использования, а при равной частоте - самый старый.
		TIME COMPLEXITY: O(1) среднее время для get/put благодаря хеш-таблице
		SPACE COMPLEXITY: O(n) где n - емкость кэша
	*/
	fmt.Println("Демонстрация работы LFU кэша с емкостью 2")
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
