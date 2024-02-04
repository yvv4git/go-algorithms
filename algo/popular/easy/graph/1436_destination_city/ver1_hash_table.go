package _436_destination_city

// Функция для нахождения города назначения
func destCityV1(paths [][]string) string {
	/*
		METHOD: Adjacency List(список смежности) on hash table
		TIME COMPLEXITY: O(n) - где n есть количество путей
		Space complexity: O(n) - храним все пути в хеш-таблице
	*/
	// Создаем хеш-таблицу для подсчета вхождений каждого города
	// Здесь мы используем пустую структуру, так как мы не храним никаких значений,
	// а просто используем ее для подсчета вхождений
	sources := make(map[string]struct{})

	// Проходим по всем путям и добавляем источники в хеш-таблицу
	for i := range paths {
		sources[paths[i][0]] = struct{}{}
	}

	// Проходим по всем путям еще раз и проверяем, не является ли конечная точка источником
	// Если не является, то это наш конечный пункт
	for i := range paths {
		if _, ok := sources[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}

	return ""
}
