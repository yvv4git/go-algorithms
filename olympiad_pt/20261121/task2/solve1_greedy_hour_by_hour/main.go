package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type device struct {
	a         int
	b         int
	t         int
	neighbors []int
}

type input struct {
	n       int
	p       int
	k       int
	x       int
	m       int
	devices []device
}

type dispatch struct {
	u int
	v int
	s int
}

func buildStrategy(in input) []dispatch {
	/*
		TASK: Построить стратегию распространения вируса и вернуть список успешных отправок копий.
		Граф - это набор устройств и каналов между ними, а сосед - устройство, соединённое каналом.
		Устройство i включено на отрезках [a_i + 24d, a_i + 24d + b_i] для всех целых d.
		Копия на устройстве i распаковывается t_i часов и должна успеть до выключения устройства.
		Успешная распаковка называется закреплением, после неё копия сразу рассылает новые копии не более чем k соседям.
		Функция возвращает срез dispatch, где dispatch{u, v, s} означает отправку с u на v в момент s.
		Пример: N = 4, p = 1, k = 1, x = 10, m = 13, устройства и связи как в примере из условия.
		Начальная копия на устройстве 1 активна в момент 10 и отправляет копию на устройство 2.
		Копия на устройстве 2 активируется в момент 12, после чего может отправить копию дальше.
		Случай k = 0: рассылать нельзя, остаётся только начальное закрепление, ответ содержит один элемент.
		Случай m = 1: процесс длится один час, а t_v >= 1, поэтому новых закреплений не будет.
		Граница: если устройство выключается ровно в момент завершения распаковки, копия выживает.
		Плохой пример: отправка на выключенное или занятое устройство приводит к гибели копии, такие отправки не включаются.

		METHOD: Greedy hour-by-hour simulation (жадное моделирование по часам)
		Идём по моментам времени от x до x+m-1 и обрабатываем активации копий.
		Массив freeAt[i] хранит момент, начиная с которого устройство i свободно.
		Для каждой активной копии на устройстве u перебираем соседей v и берём первых k подходящих.
		Подходящий сосед свободен в момент T, включён и успевает распаковаться до выключения.
		Условие успеха: r = (T - a_v) mod 24, r + t_v <= b_v, и T + t_v <= x + m - 1.
		Если условие выполнено, планируем активацию на v в момент T + t_v и запоминаем отправку.
		Стратегия жадная и не гарантирует максимум закреплений, но всегда строит корректный ответ.

		TIME COMPLEXITY: O(m * (N + E) + A), где E - число каналов, A - число закреплений.
		На каждом часе перебираются активные копии и их соседи.

		SPACE COMPLEXITY: O(N + E + A).
		Храним списки смежности, массив freeAt и события активации по моментам времени.
	*/
	n := in.n
	freeAt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		freeAt[i] = -1
	}

	last := in.x + in.m - 1
	events := make([][]int, last+1)
	events[in.x] = append(events[in.x], in.p)

	var result []dispatch
	for moment := in.x; moment <= last; moment++ {
		for _, u := range events[moment] {
			chosen := 0
			for _, v := range in.devices[u].neighbors {
				if chosen == in.k {
					break
				}
				if freeAt[v] > moment {
					continue
				}
				r := (moment - in.devices[v].a) % 24
				if r < 0 {
					r += 24
				}
				if r > in.devices[v].b {
					continue
				}
				if r+in.devices[v].t > in.devices[v].b {
					continue
				}
				finish := moment + in.devices[v].t
				if finish > last {
					continue
				}
				freeAt[v] = finish
				events[finish] = append(events[finish], v)
				result = append(result, dispatch{u: u, v: v, s: moment})
				chosen++
			}
		}
	}
	return result
}

func readInput(reader io.Reader) (input, bool) {
	var in input
	if _, err := fmt.Fscan(reader, &in.n, &in.p, &in.k); err != nil {
		return in, false
	}
	if _, err := fmt.Fscan(reader, &in.x, &in.m); err != nil {
		return in, false
	}
	in.devices = make([]device, in.n+1)
	for i := 1; i <= in.n; i++ {
		if _, err := fmt.Fscan(reader, &in.devices[i].a, &in.devices[i].b); err != nil {
			return in, false
		}
	}
	for i := 1; i <= in.n; i++ {
		if _, err := fmt.Fscan(reader, &in.devices[i].t); err != nil {
			return in, false
		}
	}
	for i := 1; i <= in.n; i++ {
		for {
			var v int
			if _, err := fmt.Fscan(reader, &v); err != nil {
				return in, false
			}
			if v == 0 {
				break
			}
			in.devices[i].neighbors = append(in.devices[i].neighbors, v)
		}
	}
	return in, true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	in, ok := readInput(reader)
	if !ok {
		return
	}
	result := buildStrategy(in)
	fmt.Fprintln(writer, len(result)+1)
	for _, d := range result {
		fmt.Fprintf(writer, "%d %d %d\n", d.u, d.v, d.s)
	}
}
