package main

import (
	"bufio"
	"fmt"
	"os"
)

func countNiceRectangles(n, m int, grid []string) int {
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

		METHOD: Bruteforce (полный перебор)
		Перебираются все возможные границы прямоугольника: r1 <= r2 и c1 <= c2.
		Для каждой части проверяются все её клетки (r, c).
		Парная клетка при повороте на 180 градусов имеет координаты (r1+r2-r, c1+c2-c).
		Если цвета хотя бы одной пары различаются, часть не является симпатичной.
		Если все пары совпали, счётчик ans увеличивается на единицу.

		TIME COMPLEXITY: O(N^3*M^3)
		Число прямоугольников O(N^2*M^2), проверка каждого занимает O(N*M).

		SPACE COMPLEXITY: O(1)
		Дополнительная память не зависит от размеров картины, кроме хранения самой сетки.
	*/
	ans := 0
	for r1 := range n {
		for r2 := r1; r2 < n; r2++ {
			for c1 := range m {
				for c2 := c1; c2 < m; c2++ {
					if isNice(grid, r1, r2, c1, c2) {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func isNice(grid []string, r1, r2, c1, c2 int) bool {
	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			if grid[r][c] != grid[r1+r2-r][c1+c2-c] {
				return false
			}
		}
	}
	return true
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

	fmt.Fprintln(writer, countNiceRectangles(n, m, grid))
}
