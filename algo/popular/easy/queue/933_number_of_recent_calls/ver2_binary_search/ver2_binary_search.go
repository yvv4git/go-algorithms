package ver2

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{pings: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	/*
		TIME COMPLEXITY: O(log n)
		Временная сложность этого алгоритма - O(log n) на каждый вызов метода Ping,
		потому что для каждого вызова метода Ping мы выполняем бинарный поиск в массиве.

		SPACE COMPLEXITY: O(n)
		Пространственная сложность этого алгоритма - O(n), где n - количество вызовов метода Ping.
		Это происходит потому, что мы храним все временные метки в массиве.
	*/
	this.pings = append(this.pings, t)
	return this.search(t)
}

func (this *RecentCounter) search(t int) int {
	/*
		Метод search использует бинарный поиск для поиска первого элемента в списке, который больше или равен t-3000.
		Этот элемент и все элементы, которые следуют за ним, являются недавними вызовами.
	*/
	// Инициализируем две переменные, которые будут использоваться в качестве границ для бинарного поиска
	left, right := 0, len(this.pings)-1

	// Пока левая граница меньше правой, выполняем бинарный поиск
	for left < right {
		// Вычисляем средний индекс
		mid := left + (right-left)/2

		// Если временная метка в середине больше или равна t-3000, то двигаем правую границу к середине
		if this.pings[mid] >= t-3000 {
			right = mid
		} else {
			// Иначе, двигаем левую границу к середине + 1
			left = mid + 1
		}
	}

	// Возвращаем количество недавних вызовов, которые были сделаны в течение последних 3 секунд
	return len(this.pings) - left
}
