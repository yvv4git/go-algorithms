package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countPatternOccurrences(n, m, l int, patterns [][]string, image []string) []int {
	/*
		TASK: Посчитать, сколько раз каждый контрольный элемент L×L встречается в изображении N×M.
		Изображение и контрольные элементы состоят из символов 0 и 1.
		Элемент должен совпадать с некоторым окном L×L точно, без поворотов и отражений.
		Пример: image = ["11", "11"], patterns = [["1"]], L = 1.
		Единица встречается в четырёх клетках, поэтому ответ [4].
		Случаи по L: L = 1 - считаем отдельные клетки; L = 2 - считаем квадраты 2×2.
		Граница: L = N = M - элемент сравнивается с изображением целиком.
		Если L > N или L > M, ни одно окно не помещается, ответ состоит из нулей.

		METHOD: Bruteforce (полный перебор)
		Окно L×L задаётся левым верхним углом (i, j), где 0 <= i <= N-L и 0 <= j <= M-L.
		Для каждого контрольного элемента patterns[p] перебираются все такие пары (i, j).
		Внутри проверяются все L строк окна: image[i+x][j:j+L] сравнивается с patterns[p][x].
		Если хотя бы одна строка не совпала, окно отбрасывается и сравнение прерывается.
		Если совпали все L строк, счётчик ans[p] увеличивается на единицу.
		Ответ - массив ans, где ans[p] равно числу вхождений элемента с индексом p.

		TIME COMPLEXITY: O(K*N*M*L^2)
		Для каждого из K элементов проверяется O(N*M) положений окна, каждое сравнение занимает O(L^2).

		SPACE COMPLEXITY: O(K*L^2)
		Дополнительная память нужна только для массива ответов ans длины K.
	*/
	ans := make([]int, len(patterns))
	for p, pattern := range patterns {
		for i := 0; i+l <= n; i++ {
			for j := 0; j+l <= m; j++ {
				matches := true
				for x := range l {
					if image[i+x][j:j+l] != pattern[x] {
						matches = false
						break
					}
				}
				if matches {
					ans[p]++
				}
			}
		}
	}
	return ans
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

	ans := countPatternOccurrences(n, m, l, patterns, image)
	for i, count := range ans {
		fmt.Fprintln(writer, i+1, count)
	}
}
