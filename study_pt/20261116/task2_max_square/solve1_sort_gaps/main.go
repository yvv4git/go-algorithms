package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func maxSquareSide(n, m int64, vert, hor []int64) int64 {
	/*
		TASK: Найти максимальную сторону квадратного участка в городе N x M, который не пересекает улицы.
		Улица - это линия нулевой ширины: вертикальная стоит на координате x, горизонтальная - на координате y.
		Квадрат должен целиком лежать в городе и не пересекать эти линии, поэтому он помещается в одну ячейку.
		Ячейка - это прямоугольник между двумя соседними вертикальными и двумя соседними горизонтальными линиями.
		Границы города 0, N и 0, M участвуют в разбиении наравне с улицами.
		Пример: N = 10, M = 10, vert = [3], hor = [5].
		По x зазоры равны 3 и 7, по y зазоры равны 5 и 5, ответ min(7, 5) = 5.
		Случаи: vert = [] и hor = [] - улиц нет, ответ min(N, M).
		Случай: vert = [0, 5], hor = [3, 1, 2] при N = 5, M = 7 даёт ответ 4.
		Граница: улица на границе города даёт зазор 0 рядом с ней, на ответ это не влияет.
		Если подходящего квадрата нет вообще, ответ 0.

		METHOD: Sort gaps (сортировка и зазоры)
		Координаты улиц по каждой оси дополняются границами города и сортируются.
		Разность соседних отсортированных координат - это зазор, ширина ячейки по этой оси.
		Максимальный зазор по x равен maxX, максимальный зазор по y равен maxY.
		Квадрат должен поместиться в одной ячейке, поэтому его сторона не превышает ни maxX, ни maxY.
		Ответ равен min(maxX, maxY): сторона min(maxX, maxY) помещается в ячейке с шириной maxX и высотой maxY.
		Порядок улиц на входе не важен, потому что координаты всё равно сортируются.

		TIME COMPLEXITY: O((Vert + Hor) * log(Vert + Hor))
		Сортировка координат по обеим осям доминирует над одним проходом для зазоров.

		SPACE COMPLEXITY: O(Vert + Hor)
		Сохраняются координаты улиц с границами по обеим осям.
	*/
	maxX := maxGap(vert, n)
	maxY := maxGap(hor, m)
	if maxX < maxY {
		return maxX
	}
	return maxY
}

// maxGap возвращает максимальную разность соседних координат после сортировки и добавления границ 0 и limit.
// Строки coords содержат координаты улиц в произвольном порядке, limit - размер города по этой оси.
func maxGap(coords []int64, limit int64) int64 {
	all := make([]int64, 0, len(coords)+2)
	all = append(all, 0, limit)
	all = append(all, coords...)
	sort.Slice(all, func(i, j int) bool { return all[i] < all[j] })

	best := int64(0)
	for i := 1; i < len(all); i++ {
		if gap := all[i] - all[i-1]; gap > best {
			best = gap
		}
	}
	return best
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int64
	var vertCount, horCount int
	if _, err := fmt.Fscan(reader, &n, &m, &vertCount); err != nil {
		return
	}

	vert := make([]int64, vertCount)
	for i := range vert {
		if _, err := fmt.Fscan(reader, &vert[i]); err != nil {
			return
		}
	}

	if _, err := fmt.Fscan(reader, &horCount); err != nil {
		return
	}

	hor := make([]int64, horCount)
	for i := range hor {
		if _, err := fmt.Fscan(reader, &hor[i]); err != nil {
			return
		}
	}

	fmt.Fprintln(writer, maxSquareSide(n, m, vert, hor))
}
