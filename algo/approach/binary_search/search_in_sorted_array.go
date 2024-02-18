package main

// Функция бинарного поиска
func binarySearch(arr []int, target int) int {
	/*
		METHOD: Binary search
		TIME COMPLEXITY: O(log n)
		SPACE COMPLEXITY: O(1)
	*/
	// Инициализируем границы поиска
	left := 0
	right := len(arr) - 1

	// Пока левая граница не больше правой
	for left <= right {
		// Находим середину текущего отрезка
		mid := left + (right-left)/2

		// Если искомый элемент найден, возвращаем его индекс
		if arr[mid] == target {
			return mid
		}

		// Если искомый элемент больше середины, ищем в правой половине
		if arr[mid] < target {
			left = mid + 1
		} else {
			// Иначе ищем в левой половине
			right = mid - 1
		}
	}

	// Если элемент не найден, возвращаем -1
	return -1
}

//func main() {
//	// Отсортированный массив
//	arr := []int{2, 3, 4, 10, 40}
//	// Элемент, который нужно найти
//	target := 10
//
//	// Вызываем функцию бинарного поиска
//	result := binarySearch(arr, target)
//
//	// Выводим результат
//	if result != -1 {
//		fmt.Printf("Элемент найден в позиции %d\n", result)
//	} else {
//		fmt.Println("Элемент не найден в массиве")
//	}
//}
