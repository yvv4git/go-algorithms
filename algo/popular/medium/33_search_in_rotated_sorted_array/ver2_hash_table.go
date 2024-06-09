package main

// Функция для поиска элемента в отсортированном сдвинутом массиве с использованием хэш-таблицы
func searchV2(nums []int, target int) int {
	/*
		METHOD: Hash table
		TIME COMPLEXITY: O(n), так как надо пройтись по всем элементам.
		SPACE COMPLEXITY: O(n), так как мы храним дополнительную хэш-таблицу размера n.
	*/
	// Создаем хэш-таблицу для быстрого поиска
	hashMap := make(map[int]int)
	for i, num := range nums {
		hashMap[num] = i
	}

	// Проверяем наличие целевого элемента в хэш-таблице
	if index, ok := hashMap[target]; ok {
		return index
	}

	// Если элемент не найден, возвращаем -1
	return -1
}

//func main() {
//	nums := []int{4, 5, 6, 7, 0, 1, 2}
//	target := 0
//	fmt.Println(search(nums, target)) // Выводим индекс искомого элемента
//}
