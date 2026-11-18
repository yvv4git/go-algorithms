package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

type Point struct {
	x int64
	y int64
}

func dist(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func greedyOrder(points []Point) []int {
	n := len(points)
	order := make([]int, 0, n)
	if n == 0 {
		return order
	}
	visited := make([]bool, n)
	current := 0
	visited[current] = true
	order = append(order, 0)
	for step := 1; step < n; step++ {
		best := -1
		var bestDist int64
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			dx := points[current].x - points[i].x
			dy := points[current].y - points[i].y
			d := dx*dx + dy*dy
			if best == -1 || d < bestDist || (d == bestDist && i < best) {
				best = i
				bestDist = d
			}
		}
		visited[best] = true
		order = append(order, best)
		current = best
	}
	return order
}

func twoOptPass(points []Point, order []int) {
	n := len(order)
	improved := true
	for improved {
		improved = false
		for i := 0; i < n-1; i++ {
			a := order[i]
			b := order[i+1]
			for j := i + 2; j < n; j++ {
				if i == 0 && j == n-1 {
					continue
				}
				c := order[j]
				d := order[(j+1)%n]
				if dist(points[a], points[c])+dist(points[b], points[d])+1e-12 < dist(points[a], points[b])+dist(points[c], points[d]) {
					for l, r := i+1, j; l < r; l, r = l+1, r-1 {
						order[l], order[r] = order[r], order[l]
					}
					improved = true
				}
			}
		}
	}
}

func cycleLength(points []Point, order []int) float64 {
	n := len(order)
	if n < 2 {
		return 0
	}
	total := 0.0
	for i := 0; i < n; i++ {
		total += dist(points[order[i]], points[order[(i+1)%n]])
	}
	return total
}

func annealRoute(points []Point) []int {
	/*
		TASK: Построить короткий маршрут курьера, который посещает все деревни и возвращается в офис.
		Деревни заданы точками на плоскости, расстояние между ними равно расстоянию по прямой.
		Курьер стартует в деревне 1, должен посетить каждую деревню 2..N ровно один раз и вернуться в деревню 1.
		Нужно вернуть последовательность номеров деревень, которая начинается и заканчивается числом 1.
		Пример: точки (0,0), (1,1), (0,1), (1,0) можно обойти как 1 3 2 4 1.
		Пример: N = 1 даёт маршрут 1 1.
		Плохой пример: маршрут 1 2 1 3 1, потому что деревня 1 встречается в середине.

		METHOD: Simulated annealing (имитация отжига)
		Сначала строим начальный порядок жадным алгоритмом и улучшаем его проходом 2-opt.
		Порядок order хранит индексы деревень по циклу обхода.
		Температура T - это число, которое управляет вероятностью принять ухудшение.
		На каждой итерации выбираем два индекса l и r и пробуем перевернуть участок order[l..r].
		Если длина цикла уменьшилась, изменение принимаем.
		Если длина увеличилась на delta, принимаем изменение с вероятностью exp(-delta/T).
		Температура постепенно уменьшается умножением на cooling, поэтому случайных ухудшений становится меньше.
		Лучший найденный порядок сохраняем отдельно в best.
		Генератор случайных чисел имеет фиксированное начальное значение, поэтому результат воспроизводим.

		TIME COMPLEXITY: O(N^2 + I)
		- Начальные greedy и 2-opt занимают до O(N^2), далее выполняется I итераций с операциями O(1) или O(N) в худшем для переворота.

		SPACE COMPLEXITY: O(N)
		- Храним текущий и лучший порядок обхода.
	*/
	n := len(points)
	order := greedyOrder(points)
	if n <= 2 {
		return orderToRoute(order)
	}
	twoOptPass(points, order)
	best := append([]int(nil), order...)
	bestLen := cycleLength(points, best)
	curLen := bestLen

	rng := rand.New(rand.NewSource(1))
	iterations := 20000 * 2
	if n < 200 {
		iterations = 200000
	}
	T := math.Max(bestLen/10.0, 1.0)
	cooling := 0.9995

	for it := 0; it < iterations; it++ {
		l := rng.Intn(n)
		r := rng.Intn(n)
		if l == r {
			continue
		}
		if l > r {
			l, r = r, l
		}
		if l == 0 && r == n-1 {
			continue
		}
		for i, j := l, r; i < j; i, j = i+1, j-1 {
			order[i], order[j] = order[j], order[i]
		}
		newLen := cycleLength(points, order)
		delta := newLen - curLen
		if delta < 0 || rng.Float64() < math.Exp(-delta/T) {
			curLen = newLen
			if curLen+1e-12 < bestLen {
				bestLen = curLen
				copy(best, order)
			}
		} else {
			for i, j := l, r; i < j; i, j = i+1, j-1 {
				order[i], order[j] = order[j], order[i]
			}
		}
		T *= cooling
		if T < 1e-9 {
			T = 1e-9
		}
	}
	twoOptPass(points, best)
	return orderToRoute(best)
}

func orderToRoute(order []int) []int {
	route := make([]int, 0, len(order)+1)
	for _, v := range order {
		route = append(route, v+1)
	}
	if len(order) > 0 {
		route = append(route, order[0]+1)
	}
	return route
}

func routeLength(points []Point, route []int) float64 {
	if len(route) < 2 {
		return 0
	}
	total := 0.0
	for i := 0; i+1 < len(route); i++ {
		total += dist(points[route[i]-1], points[route[i+1]-1])
	}
	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(reader, &points[i].x, &points[i].y); err != nil {
			return
		}
	}

	route := annealRoute(points)
	for i, v := range route {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
