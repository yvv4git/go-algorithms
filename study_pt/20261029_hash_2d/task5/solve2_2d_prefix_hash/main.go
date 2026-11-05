package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	hashP1 uint64 = 0x9e3779b185ebca87
	hashQ1 uint64 = 0xc2b2ae3d27d4eb4f
	hashP2 uint64 = 0x165667b19e3779f9
	hashQ2 uint64 = 0x85ebca77c2b2ae63
)

func countNiceRectanglesV2(n, m int, grid []string) int {
	/*
		TASK: Посчитать количество прямоугольных частей картины, симметричных относительно своего центра.
		Картина - это N строк по M строчных латинских букв, каждая буква задаёт цвет клетки.
		Часть - это прямоугольник с границами r1..r2 по строкам и c1..c2 по столбцам.
		Симметрия относительно центра означает, что клетка (r, c) совпадает с клеткой (r1+r2-r, c1+c2-c).
		Пример: grid = ["abc", "acb"], ответ равен 8.
		Шесть частей из одной клетки симпатичны, плюс часть a/a в первом столбце и часть bc/cb в последних двух столбцах.
		Случаи: картина 1×1 даёт ответ 1; картина 1×2 "aa" даёт ответ 3; картина 1×2 "ab" даёт ответ 2.
		Граница: картина 100×100 из одинаковых букв даёт 5050*5050 симпатичных частей.
		Если все клетки разные, симпатичны только части из одной клетки, ответ N*M.

		METHOD: 2D prefix hash (двумерный префиксный хеш)
		Префиксный хеш pref[i][j] описывает прямоугольник от (0, 0) до (i-1, j-1).
		Строки учитываются с основанием P, столбцы - с основанием Q, степени хранятся в powP и powQ.
		Хеш произвольного прямоугольника получается из четырёх значений pref включением-исключением за O(1).
		Для повёрнутой картины строится отдельный префиксный хеш: rotPref[i][j] использует grid[N-1-i][M-1-j].
		Части (r1, r2, c1, c2) соответствует повёрнутая часть (N-1-r2, N-1-r1, M-1-c2, M-1-c1).
		Если оба хеша исходной и повёрнутой части совпали, часть симпатична.
		Две независимые пары оснований дают вероятность ложного совпадения около 2^-128.

		TIME COMPLEXITY: O(N^2*M^2)
		Перебираются все O(N^2*M^2) прямоугольников, хеш каждого считается за O(1).

		SPACE COMPLEXITY: O(N*M)
		Четыре префиксных таблицы и степени оснований занимают O(N*M).
	*/
	maxSize := n
	if m > maxSize {
		maxSize = m
	}

	powP1 := makePowers(hashP1, maxSize)
	powQ1 := makePowers(hashQ1, maxSize)
	powP2 := makePowers(hashP2, maxSize)
	powQ2 := makePowers(hashQ2, maxSize)

	pref1 := buildPrefix(n, m, grid, hashP1, hashQ1)
	pref2 := buildPrefix(n, m, grid, hashP2, hashQ2)
	rotPref1 := buildRotatedPrefix(n, m, grid, hashP1, hashQ1)
	rotPref2 := buildRotatedPrefix(n, m, grid, hashP2, hashQ2)

	ans := 0
	for r1 := range n {
		for r2 := r1; r2 < n; r2++ {
			for c1 := range m {
				for c2 := c1; c2 < m; c2++ {
					if rectangleHash(pref1, powP1, powQ1, r1, r2, c1, c2) !=
						rectangleHash(rotPref1, powP1, powQ1, n-1-r2, n-1-r1, m-1-c2, m-1-c1) {
						continue
					}
					if rectangleHash(pref2, powP2, powQ2, r1, r2, c1, c2) ==
						rectangleHash(rotPref2, powP2, powQ2, n-1-r2, n-1-r1, m-1-c2, m-1-c1) {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func buildPrefix(n, m int, grid []string, p, q uint64) [][]uint64 {
	pref := make([][]uint64, n+1)
	for i := range pref {
		pref[i] = make([]uint64, m+1)
	}
	for i := range n {
		for j := range m {
			value := uint64(grid[i][j])
			pref[i+1][j+1] = pref[i][j+1]*p + pref[i+1][j]*q - pref[i][j]*p*q + value
		}
	}
	return pref
}

func buildRotatedPrefix(n, m int, grid []string, p, q uint64) [][]uint64 {
	pref := make([][]uint64, n+1)
	for i := range pref {
		pref[i] = make([]uint64, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			value := uint64(grid[n-1-i][m-1-j])
			pref[i+1][j+1] = pref[i][j+1]*p + pref[i+1][j]*q - pref[i][j]*p*q + value
		}
	}
	return pref
}

func rectangleHash(pref [][]uint64, powP, powQ []uint64, r1, r2, c1, c2 int) uint64 {
	height := r2 - r1 + 1
	width := c2 - c1 + 1
	h := pref[r2+1][c2+1]
	h -= pref[r1][c2+1] * powP[height]
	h -= pref[r2+1][c1] * powQ[width]
	h += pref[r1][c1] * powP[height] * powQ[width]
	return h
}

func makePowers(base uint64, maxSize int) []uint64 {
	pow := make([]uint64, maxSize+1)
	pow[0] = 1
	for i := 1; i <= maxSize; i++ {
		pow[i] = pow[i-1] * base
	}
	return pow
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	if _, err := fmt.Fscan(reader, &n, &m); err != nil {
		return
	}

	grid := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	fmt.Fprintln(writer, countNiceRectanglesV2(n, m, grid))
}
