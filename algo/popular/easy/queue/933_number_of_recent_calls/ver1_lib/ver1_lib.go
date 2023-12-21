package ver1_lib

import "container/list"

type RecentCounter struct {
	queue *list.List
}

func Constructor() RecentCounter {
	return RecentCounter{queue: list.New()}
}

// Ping - временная сложность метода Ping - O(1) в среднем, так как мы просто добавляем элемент в очередь и
// удаляем элементы из очереди, если они не входят в окно разрешенного времени.
// Пространственная сложность - O(n), где n - количество вызовов метода Ping.
// Это связано с тем, что мы храним все вызовы метода Ping в очереди.
func (this *RecentCounter) Ping(t int) int {
	// Добавляем текущее время в очередь
	this.queue.PushBack(t)

	// Удаляем все элементы, которые не входят в окно разрешенного времени
	for this.queue.Len() > 0 && this.queue.Front().Value.(int) < t-3000 {
		this.queue.Remove(this.queue.Front())
	}

	// Возвращаем размер очереди, которая содержит только разрешенные временные интервалы
	return this.queue.Len()
}
