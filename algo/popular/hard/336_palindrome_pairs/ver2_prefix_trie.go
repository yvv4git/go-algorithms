package main

// TrieNode - узел префиксного дерева (Trie)
type TrieNode struct {
	childrens    [26]*TrieNode // Массив указателей на дочерние узлы
	subPalindrom []int         // Индексы слов, которые являются палиндромами в поддереве
	wordIndex    int           // Индекс слова, если узел является концом слова, иначе -1
}

// Trie - префиксное дерево (Trie)
type Trie struct {
	root *TrieNode // Корневой узел
}

var res [][]int // Результирующий список пар индексов

// Инициализация префиксного дерева
func initTrie() *Trie {
	return &Trie{
		root: &TrieNode{wordIndex: -1},
	}
}

// Проверка, является ли подстрока палиндромом
func isPalindromeV2(word string, i int, j int) bool {
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// Вставка слова в префиксное дерево
func (t Trie) insert(wordIndex int, word string) {
	current := t.root
	for i := len(word) - 1; i >= 0; i-- {
		// Проверка, является ли подстрока (0, i) палиндромом
		if isPalindromeV2(word, 0, i) {
			current.subPalindrom = append(current.subPalindrom, wordIndex)
		}

		// Добавление символа в узел префиксного дерева
		charIndex := word[i] - 'a'
		if current.childrens[charIndex] == nil {
			current.childrens[charIndex] = &TrieNode{wordIndex: -1}
		}
		current = current.childrens[charIndex]
	}
	current.wordIndex = wordIndex
}

// Поиск пар палиндромов в префиксном дереве
func (t Trie) search(i int, word string) {
	current := t.root

	// Пустая строка "" в корне
	if current.wordIndex != -1 && current.wordIndex != i && isPalindromeV2(word, 0, len(word)-1) {
		res = append(res, []int{i, current.wordIndex})
	}

	for j := 0; j < len(word); j++ {
		charIndex := word[j] - 'a'
		current = current.childrens[charIndex]
		if current == nil {
			break
		}

		// Найден префикс слова, остаток слова является палиндромом -> abcxxx + reverse(abc)
		if current.wordIndex != -1 && current.wordIndex != i && isPalindromeV2(word, j+1, len(word)-1) {
			res = append(res, []int{i, current.wordIndex})
		}
	}

	// current = последний символ слова
	// Слово cba, предположим, 'cbaxyx' в дереве, reverse('xyxabc')
	if current != nil && len(current.subPalindrom) > 0 {
		for _, j := range current.subPalindrom {
			if i != j {
				res = append(res, []int{i, j})
			}
		}
	}
}

// Основная функция для поиска всех пар индексов, образующих палиндром
func palindromePairsV2(words []string) [][]int {
	/*
		МЕТОД: Trie
		Использование префиксного дерева (Trie) для хранения слов в обратном порядке.
		Это позволяет эффективно проверять, могут ли два слова образовать палиндром.

		ВРЕМЕННАЯ СЛОЖНОСТЬ: O(n * m^2), где n — количество слов, m — средняя длина слова.
		Это связано с тем, что для каждого слова мы проверяем все возможные подстроки и их инверсии.

		ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ: O(n * m), где n — количество слов, m — средняя длина слова.
		Это связано с использованием префиксного дерева для хранения слов и их индексов.
	*/
	trie := initTrie()

	// O(n * m^2)
	// n : количество слов
	// m : длина каждого слова
	for i := 0; i < len(words); i++ {
		trie.insert(i, words[i])
	}

	res = [][]int{}
	// O(n * m^2)
	// n : количество слов
	// m : длина каждого слова
	for i := 0; i < len(words); i++ {
		trie.search(i, words[i])
	}
	return res
}
