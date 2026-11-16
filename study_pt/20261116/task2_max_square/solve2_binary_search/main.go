package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func maxSquareSideV2(n, m int64, vert, hor []int64) int64 {
	/*
		TASK: Найти максимальную сторону квадратного участка в городе N x M, который не пересекает улицы.
		Улица - это линия нулевой ширины: вертикальная стоит на координате x, горизонтальная - на координате y.
		Квадрат должен целиком лежать в городе и не пересекать эти линии, поэтому он помещается в одну ячейку.
		Ячейка - это прямоугольник между двумя соседними вертикальными и двумя соседними горизонтальными линиями.
		Границы города 0, N и 0, M участвуют в разбиении наравне с улицами.
		Пример: N = 10, M = 10, vert = [3], hor = [5].
		При k = 5 по x есть зазор 7 >= 5 и по y есть зазор 5 >= 5, значит k = 5 достигимо.
		При k = 6 по y зазоры 5 и 5 меньше 6, значит k = 6 недостижимо, ответ 5.
		Случаи: vert = [] и hor = [] - улиц нет, ответ min(N, M).
		Случай: vert = [0, 5], hor = [3, 1, 2] при N = 5, M = 7 даёт ответ 4.
		Граница: улица на границе города даёт зазор 0 рядом с ней, на ответ это не влияет.
		Если подходящего квадрата нет вообще, ответ 0.

		METHOD: Binary search (бинарный поиск по ответу)
		Бинарный поиск - это метод, который за O(log L) итераций находит границу между подходящими и неподходящими значениями.
		Координаты улиц по каждой оси дополняются границами города и сортируются один раз.
		Ответ k ищется в диапазоне от 0 до min(N, M).
		Проверка hasGap(coords, k) возвращает true, если среди соседних отсортированных координат есть разность не меньше k.
		Если hasGap по x и hasGap по y оба true, k достигимо, и берётся больший k.
		Если хотя бы одна проверка false, k слишком велико, и берётся меньший k.
		Инвариант поиска: lo всегда достижимо, hi - недостижимо или равно итоговому ответу.
		Порядок улиц на входе не важен, потому что координаты всё равно сортируются.

		TIME COMPLEXITY: O((Vert + Hor) * log(Vert + Hor) + log(min(N, M)) * (Vert + Hor))
		Сортировка выполняется один раз, бинарный поиск делает O(log min(N, M)) проверок по O(Vert + Hor).

		SPACE COMPLEXITY: O(Vert + Hor)
		Сохраняются координаты улиц с границами по обеим осям.
	*/
	sortedX := sortedWithBorders(vert, n)
	sortedY := sortedWithBorders(hor, m)

	lo := int64(0)
	hi := min(m, n)
	for lo < hi {
		mid := lo + (hi-lo+1)/2
		if hasGap(sortedX, mid) && hasGap(sortedY, mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// sortedWithBorders возвращает отсортированный срез координат улиц с добавленными границами 0 и limit.
func sortedWithBorders(coords []int64, limit int64) []int64 {
	all := make([]int64, 0, len(coords)+2)
	all = append(all, 0, limit)
	all = append(all, coords...)
	slices.Sort(all)
	return all
}

// hasGap сообщает, есть ли среди соседних элементов отсортированного среза разность не меньше k.
func hasGap(sorted []int64, k int64) bool {
	for i := 1; i < len(sorted); i++ {
		if sorted[i]-sorted[i-1] >= k {
			return true
		}
	}
	return false
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

	fmt.Fprintln(writer, maxSquareSideV2(n, m, vert, hor))
}
