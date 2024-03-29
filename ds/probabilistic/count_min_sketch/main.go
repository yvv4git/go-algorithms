package main

func main() {
	/*
		Функция Count проходит по каждой строке матрицы и для каждой строки вычисляет индекс, используя хэш-функцию.
		Затем она получает значение счетчика в соответствующей ячейке матрицы и сравнивает его с текущим минимальным значением.
		Если значение счетчика меньше текущего минимального, оно становится новым минимальным значением.
		В конце функция возвращает минимальное значение, которое представляет приблизительное количество элемента в наборе данных.

		Таким образом, когда мы вызываем cms.Count("item1"), функция возвращает 2, потому что элемент "item1" был добавлен в Count-Min Sketch два раза.
		Надо учитывать, что это приблизительное количество, так как Count-Min Sketch использует несколько хэш-функций для уменьшения вероятности ошибки.
	*/
	// Использование CountMinSketch
	cms := NewCountMinSketch(1000, 5)

	cms.Add("item1")
	cms.Add("item2")
	cms.Add("item1")

	count := cms.Count("item1")
	println(count) // Вывод: 2
}
