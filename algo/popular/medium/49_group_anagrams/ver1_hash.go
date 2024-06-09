package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	/*
		METHOD: Hashing
		TIME COMPLEXITY: O(n * m log m), где n - количество слов в strs, а m - средняя длина слова.
		SPACE COMPLEXITY: O(n * m), так как в худшем случае мы можем хранить каждое слово и его отсортированный анаграммный ключ.
	*/
	groups := make(map[string][]string)

	for _, str := range strs {
		// Сортируем буквы в слове и используем их в качестве ключа
		sortedStr := sortString(str)                       // Создаем хеш на основе специфического алгоритма
		groups[sortedStr] = append(groups[sortedStr], str) // Группируем анаграммы по ключу
	}

	// Создаем двумерный массив анаграмм
	/*
		[][]string{
				{"eat", "tea", "ate"}, <- группа-1
				{"tan", "nat"}, <- группа-2
				{"bat"},	<- группа-3
			}
	*/
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// sortString сортирует буквы в строке и возвращает отсортированную строку
func sortString(str string) string {
	s := strings.Split(str, "") // O(n) делит строку str на отдельные символы и сохраняет их в новый срез (slice) строк
	sort.Strings(s)             // O(n log n)
	return strings.Join(s, "")  // O(n) объединяем все символы в одну строку
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}
