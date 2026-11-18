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

func greedyRoute(points []Point) []int {
	/*
		TASK: Построить маршрут курьера, который посещает все деревни и возвращается в офис.
		Деревни заданы точками на плоскости, расстояние между ними равно расстоянию по прямой.
		Курьер стартует в деревне 1, должен посетить каждую деревню 2..N ровно один раз и вернуться в деревню 1.
		Нужно вернуть последовательность номеров деревень, которая начинается и заканчивается числом 1.
		Пример: точки (0,0), (1,1), (0,1), (1,0) дают маршрут 1 3 2 4 1.
		Пример: N = 1 даёт маршрут 1 1, потому что старт и финиш совпадают.
		Плохой пример формата: маршрут 1 2 2 1, потому что деревня 2 посещена дважды, а деревня 3 пропущена.

		METHOD: Greedy nearest neighbor (жадный ближайший сосед)
		На каждом шаге выбираем ближайшую непосещённую деревню к текущей.
		Текущая деревня сначала равна 1, а массив visited отмечает уже включённые в маршрут деревни.
		Для выбора кандидата перебираем все непосещённые деревни и сравниваем квадрат расстояния до них.
		Квадрат расстояния (dx*dx + dy*dy) сравнивать удобнее, потому что корень не меняет порядок.
		Если квадраты расстояний равны, берём деревню с меньшим номером для однозначного результата.
		После обхода всех деревень добавляем в конец маршрута деревню 1.

		TIME COMPLEXITY: O(N^2)
		- Для каждой из N позиций маршрута просматриваем до N кандидатов.

		SPACE COMPLEXITY: O(N)
		- Храним массив visited и результирующий маршрут.
	*/
	n := len(points)
	route := make([]int, 0, n+1)
	if n == 0 {
		return route
	}
	visited := make([]bool, n)
	current := 0
	visited[current] = true
	route = append(route, 1)

	for step := 1; step < n; step++ {
		best := -1
		var bestDist int64
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			dx := points[current].x - points[i].x
			dy := points[current].y - points[i].y
			dist := dx*dx + dy*dy
			if best == -1 || dist < bestDist || (dist == bestDist && i < best) {
				best = i
				bestDist = dist
			}
		}
		visited[best] = true
		route = append(route, best+1)
		current = best
	}

	route = append(route, 1)
	return route
}

func routeLength(points []Point, route []int) float64 {
	if len(route) < 2 {
		return 0
	}
	total := 0.0
	for i := 0; i+1 < len(route); i++ {
		a := points[route[i]-1]
		b := points[route[i+1]-1]
		dx := float64(a.x - b.x)
		dy := float64(a.y - b.y)
		total += math.Sqrt(dx*dx + dy*dy)
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

	route := greedyRoute(points)
	for i, v := range route {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
