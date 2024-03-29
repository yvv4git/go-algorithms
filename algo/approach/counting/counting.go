package main

import "fmt"

/*
Подход counting, также известный как подсчет частоты элементов, используется для подсчета количества вхождений каждого элемента в наборе данных.
Этот подход широко используется в различных областях, включая компьютерные науки, математику, статистику и информатику.

Основные шаги в подходе counting:
1. Инициализация счетчика: Создается структура данных (часто это словарь, хэш-таблица или массив),
которая будет использоваться для хранения количества вхождений каждого элемента.
2. Проход по данным: Происходит проход по всем элементам данных, и для каждого элемента увеличивается соответствующее значение в структуре данных.
3. Получение результатов: После завершения прохода по данным, структура данных будет содержать количество вхождений каждого элемента.
*/

// Функция для подсчета элементов в массиве
func countElements(arr []int) map[int]int {
	// Создаем пустую map для подсчета
	countMap := make(map[int]int)

	// Проходим по каждому элементу в массиве
	for _, element := range arr {
		// Увеличиваем значение элемента на 1
		countMap[element]++
	}

	// Возвращаем map с подсчитанными элементами
	return countMap
}

func main() {
	// Пример использования
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	counts := countElements(numbers)

	// Выводим результаты
	for key, value := range counts {
		fmt.Printf("Элемент %d встречается %d раз(а)\n", key, value)
	}
}
