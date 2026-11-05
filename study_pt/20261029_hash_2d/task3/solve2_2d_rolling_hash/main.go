package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	hashP1 uint64 = 0x9e3779b185ebca87
	hashQ1 uint64 = 0xc2b2ae3d27d4eb4f
	hashP2 uint64 = 0x165667b19e3779f9
	hashQ2 uint64 = 0x85ebca77c2b2ae63
)

type hashKey struct {
	h1 uint64
	h2 uint64
}

func countPatternOccurrencesV2(n, m, l int, patterns [][]string, image []string) []int {
	/*
		TASK: Посчитать, сколько раз каждый контрольный элемент L×L встречается в изображении N×M.
		Изображение и контрольные элементы состоят из символов 0 и 1.
		Элемент должен совпадать с некоторым окном L×L точно, без поворотов и отражений.
		Пример: image = ["11", "11"], patterns = [["1"]], L = 1.
		Единица встречается в четырёх клетках, поэтому ответ [4].
		Случаи по L: L = 1 - считаем отдельные клетки; L = 2 - считаем квадраты 2×2.
		Граница: L = N = M - элемент сравнивается с изображением целиком.
		Если L > N или L > M, ни одно окно не помещается, ответ состоит из нулей.

		METHOD: 2D rolling hash (двумерный скользящий хеш)
		Хеш квадрата L×L - это число, которое зависит от всех его клеток.
		Строки учитываются с основанием P, столбцы - с основанием Q.
		Для каждого элемента patterns[p] пара хешей h1 и h2 считается функцией hashPattern в том же порядке, что и окна изображения.
		Для изображения поддерживается вертикальный хеш v1[c] для окна из L строк в каждом столбце c.
		При добавлении новой строки v1[c] = v1[c]*P + bit, а при выходе старой строки вычитается bit*P^L.
		Горизонтальный хеш h1 обновляется как h1 = h1*Q + v1[c] с вычитанием v1[c-L]*Q^L.
		Пара hashKey{h1, h2} ищется в patternIndex, и счётчики всех совпавших элементов увеличиваются.
		Две независимые пары оснований дают вероятность ложного совпадения около 2^-128.

		TIME COMPLEXITY: O(N*M + K*L^2)
		Каждая из N*M клеток обрабатывается за O(1), хеши K элементов строятся за O(K*L^2).

		SPACE COMPLEXITY: O(M + K*L^2) без учёта хранения изображения
		Вертикальные хеши v1 и v2 занимают O(M), хеш-таблица элементов - O(K*L^2).
	*/
	ans := make([]int, len(patterns))
	if l > n || l > m {
		return ans
	}

	patternIndex := make(map[hashKey][]int, len(patterns))
	for idx, pattern := range patterns {
		key := hashKey{
			h1: hashPattern(pattern, l, hashP1, hashQ1),
			h2: hashPattern(pattern, l, hashP2, hashQ2),
		}
		patternIndex[key] = append(patternIndex[key], idx)
	}

	powL1 := powUint(hashP1, l)
	powL2 := powUint(hashP2, l)
	qowL1 := powUint(hashQ1, l)
	qowL2 := powUint(hashQ2, l)

	v1 := make([]uint64, m)
	v2 := make([]uint64, m)
	for r := range n {
		row := image[r]
		for c := range m {
			bit := uint64(row[c] - '0')
			v1[c] = v1[c]*hashP1 + bit
			v2[c] = v2[c]*hashP2 + bit
		}

		if r >= l {
			oldRow := image[r-l]
			for c := 0; c < m; c++ {
				bit := uint64(oldRow[c] - '0')
				v1[c] -= bit * powL1
				v2[c] -= bit * powL2
			}
		}

		if r < l-1 {
			continue
		}

		var h1, h2 uint64
		for c := 0; c < m; c++ {
			h1 = h1*hashQ1 + v1[c]
			h2 = h2*hashQ2 + v2[c]
			if c >= l {
				h1 -= v1[c-l] * qowL1
				h2 -= v2[c-l] * qowL2
			}
			if c < l-1 {
				continue
			}
			key := hashKey{h1: h1, h2: h2}
			for _, idx := range patternIndex[key] {
				ans[idx]++
			}
		}
	}
	return ans
}

func hashPattern(rows []string, l int, p, q uint64) uint64 {
	var h uint64
	for r := 0; r < l; r++ {
		var rowHash uint64
		for c := 0; c < l; c++ {
			rowHash = rowHash*q + uint64(rows[r][c]-'0')
		}
		h = h*p + rowHash
	}
	return h
}

func powUint(base uint64, exp int) uint64 {
	result := uint64(1)
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func readInput() (int, int, int, [][]string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024), 1<<20)

	nextLine := func() string {
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				return line
			}
		}
		return ""
	}

	var n, m, k, l int
	fmt.Sscan(nextLine(), &n, &m, &k, &l)

	patterns := make([][]string, k)
	for p := 0; p < k; p++ {
		rows := make([]string, l)
		for r := 0; r < l; r++ {
			rows[r] = nextLine()
		}
		patterns[p] = rows
	}

	image := make([]string, n)
	for i := 0; i < n; i++ {
		image[i] = nextLine()
	}
	return n, m, l, patterns, image
}

func main() {
	n, m, l, patterns, image := readInput()

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	ans := countPatternOccurrencesV2(n, m, l, patterns, image)
	for i, count := range ans {
		fmt.Fprintln(writer, i+1, count)
	}
}
