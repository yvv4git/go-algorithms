package main

import (
	"bufio"
	"fmt"
	"math"
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

func improveTwoOpt(points []Point, order []int) []int {
	/*
		TASK: Построить короткий маршрут курьера, который посещает все деревни и возвращается в офис.
		Деревни заданы точками на плоскости, расстояние между ними равно расстоянию по прямой.
		Курьер стартует в деревне 1, должен посетить каждую деревню 2..N ровно один раз и вернуться в деревню 1.
		Нужно вернуть последовательность номеров деревень, которая начинается и заканчивается числом 1.
		Пример: точки (0,0), (1,1), (0,1), (1,0) можно обойти как 1 3 2 4 1.
		Пример: N = 1 даёт маршрут 1 1.
		Плохой пример: маршрут 1 2 3 1 при N = 4, потому что деревня 4 не посещена.

		METHOD: 2-opt (локальное улучшение 2-opt)
		Сначала строим начальный порядок обхода жадным выбором ближайшей непосещённой деревни.
		Порядок order хранит индексы деревень в порядке посещения, начиная с 0.
		Маршрут - это цикл: после последней деревни в order снова идёт первая.
		2-opt берёт два ребра цикла (a,b) и (c,d) и заменяет их на (a,c) и (b,d), если путь становится короче.
		Такая замена равносильна перевороту участка маршрута между b и c.
		Переворот выгоден, когда dist(a,c)+dist(b,d) меньше dist(a,b)+dist(c,d).
		Алгоритм перебирает пары разрезов и повторяет проходы, пока есть улучшения.

		TIME COMPLEXITY: O(k * N^2)
		- Один проход проверяет O(N^2) пар, k - число проходов до отсутствия улучшений.

		SPACE COMPLEXITY: O(N)
		- Храним порядок обхода и координаты точек.
	*/
	n := len(order)
	if n < 4 {
		return orderToRoute(order)
	}
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
					reverse(order, i+1, j)
					improved = true
				}
			}
		}
	}
	return orderToRoute(order)
}

func reverse(order []int, left, right int) {
	for left < right {
		order[left], order[right] = order[right], order[left]
		left++
		right--
	}
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

	route := improveTwoOpt(points, greedyOrder(points))
	for i, v := range route {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
