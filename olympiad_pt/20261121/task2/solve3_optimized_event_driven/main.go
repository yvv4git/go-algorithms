package main

import (
	"io"
	"os"
	"slices"
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

type devState struct {
	freeAt  int32
	canMask uint32
	t       uint8
}

func sortNeighbors(nb []int, state []devState, tmp []int) {
	d := len(nb)
	if d < 2 {
		return
	}

	if d <= 16 {
		for i := 1; i < d; i++ {
			v := nb[i]
			tv := state[v].t
			j := i - 1
			for j >= 0 {
				u := nb[j]
				tu := state[u].t
				if tu < tv || (tu == tv && u < v) {
					break
				}
				nb[j+1] = nb[j]
				j--
			}
			nb[j+1] = v
		}
		return
	}

	var cnt [24]int
	for _, v := range nb {
		cnt[state[v].t]++
	}

	var start [24]int
	sum := 0
	for t := 0; t < 24; t++ {
		start[t] = sum
		sum += cnt[t]
	}

	pos := start
	for _, v := range nb {
		t := state[v].t
		tmp[pos[t]] = v
		pos[t]++
	}
	copy(nb, tmp[:d])

	for t := 0; t < 24; t++ {
		if cnt[t] > 1 {
			slices.Sort(nb[start[t] : start[t]+cnt[t]])
		}
	}
}

func buildStrategyV3(in input) []dispatch {
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

		METHOD: Optimized event-driven simulation (оптимизированное событийное моделирование)
		События активации хранятся в плоских массивах head, count, eventDev и eventNext вместо [][]int.
		Это убирает до миллиона маленьких слайсов и лишние аллокации при добавлении событий.
		Соседи сортируются counting sort по t_v, потому что t_v <= 23, а внутри равных t_v - по номеру устройства.
		Маленькие списки соседей сортируются insertion sort, чтобы не платить за подсчёт по 24 группам.
		Для каждого устройства canStart хранится битовой маской на 24 часа, а freeAt - в int32.
		Состояние устройства собрано в devState, чтобы внутренний цикл читал одну компактную структуру.
		Счётчик freeCount хранит число свободных устройств; при freeCount = 0 рассылка прекращается.
		При активации на устройстве u перебираем заранее отсортированных соседей и берём первых k подходящих.
		Соседи с t_v > remaining не проверяются, потому что распаковка уже не успеет до конца процесса.
		Ввод читается целиком и разбирается по индексу, а вывод собирается в один буфер и пишется один раз.

		TIME COMPLEXITY: O(E log E + m + A * D) в худшем случае.
		Counting sort по t_v занимает O(E + 24N), но сортировка групп с равным t_v по номеру может стоить O(E log E).
		Каждая активация перебирает соседей до первых k подходящих, а ранние выходы уменьшают работу на практике.

		SPACE COMPLEXITY: O(N + E + A + m).
		Храним списки смежности, devState, плоские события, массив head и count, а также список отправок.
	*/
	if in.k == 0 || in.m <= 1 {
		return nil
	}

	n := in.n
	state := make([]devState, n+1)
	for v := 1; v <= n; v++ {
		dev := &in.devices[v]
		state[v].freeAt = -1
		state[v].t = uint8(dev.t)

		maxR := dev.b - dev.t
		var mask uint32
		for r := 0; r <= maxR; r++ {
			h := dev.a + r
			if h >= 24 {
				h -= 24
			}
			mask |= 1 << uint(h)
		}
		state[v].canMask = mask
	}

	maxDeg := 0
	for u := 1; u <= n; u++ {
		if len(in.devices[u].neighbors) > maxDeg {
			maxDeg = len(in.devices[u].neighbors)
		}
	}
	tmp := make([]int, maxDeg)
	for u := 1; u <= n; u++ {
		sortNeighbors(in.devices[u].neighbors, state, tmp)
	}

	last := in.x + in.m - 1
	freeCount := n
	m := in.m
	maxEvents := n * m

	head := make([]int32, m)
	count := make([]int32, m)
	eventDev := make([]int32, 0, maxEvents)
	eventNext := make([]int32, 0, maxEvents)

	eventDev = append(eventDev, int32(in.p))
	eventNext = append(eventNext, 0)
	head[0] = 1
	count[0] = 1

	result := make([]dispatch, 0, maxEvents)
	for idx := 0; idx < m; idx++ {
		if head[idx] == 0 {
			continue
		}
		moment := in.x + idx
		if moment >= last {
			break
		}
		remaining := last - moment
		if idx > 0 {
			freeCount += int(count[idx])
		}

		canBit := uint32(1) << uint(moment%24)
		moment32 := int32(moment)
		for e := head[idx]; e != 0; e = eventNext[e-1] {
			if freeCount == 0 {
				break
			}
			u := int(eventDev[e-1])
			chosen := 0
			for _, v := range in.devices[u].neighbors {
				if chosen == in.k {
					break
				}
				st := state[v]
				if int(st.t) > remaining {
					break
				}
				if st.freeAt > moment32 {
					continue
				}
				if st.canMask&canBit == 0 {
					continue
				}

				finish := moment + int(st.t)
				st.freeAt = int32(finish)
				state[v] = st
				freeCount--

				nextIdx := finish - in.x
				eventDev = append(eventDev, int32(v))
				eventNext = append(eventNext, head[nextIdx])
				head[nextIdx] = int32(len(eventDev))
				count[nextIdx]++

				result = append(result, dispatch{u: u, v: v, s: moment})
				chosen++
			}
		}
	}
	return result
}

type byteReader struct {
	data []byte
	pos  int
}

func (r *byteReader) readInt() (int, bool) {
	data := r.data
	i := r.pos
	for i < len(data) && (data[i] < '0' || data[i] > '9') {
		i++
	}
	if i == len(data) {
		r.pos = i
		return 0, false
	}

	val := 0
	for i < len(data) && data[i] >= '0' && data[i] <= '9' {
		val = val*10 + int(data[i]-'0')
		i++
	}
	r.pos = i
	return val, true
}

func readInput(r *byteReader) (input, bool) {
	var in input
	var ok bool
	if in.n, ok = r.readInt(); !ok {
		return in, false
	}
	if in.p, ok = r.readInt(); !ok {
		return in, false
	}
	if in.k, ok = r.readInt(); !ok {
		return in, false
	}
	if in.x, ok = r.readInt(); !ok {
		return in, false
	}
	if in.m, ok = r.readInt(); !ok {
		return in, false
	}

	in.devices = make([]device, in.n+1)
	for i := 1; i <= in.n; i++ {
		if in.devices[i].a, ok = r.readInt(); !ok {
			return in, false
		}
		if in.devices[i].b, ok = r.readInt(); !ok {
			return in, false
		}
	}
	for i := 1; i <= in.n; i++ {
		if in.devices[i].t, ok = r.readInt(); !ok {
			return in, false
		}
	}
	for i := 1; i <= in.n; i++ {
		for {
			v, ok := r.readInt()
			if !ok {
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

func appendInt(buf []byte, x int) []byte {
	if x == 0 {
		return append(buf, '0')
	}
	if x < 0 {
		buf = append(buf, '-')
		x = -x
	}
	var tmp [20]byte
	i := len(tmp)
	for x > 0 {
		i--
		tmp[i] = byte('0' + x%10)
		x /= 10
	}
	return append(buf, tmp[i:]...)
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}
	r := &byteReader{data: data}
	in, ok := readInput(r)
	if !ok {
		return
	}
	result := buildStrategyV3(in)

	out := make([]byte, 0, 2+len(result)*24)
	out = appendInt(out, len(result)+1)
	out = append(out, '\n')
	for _, d := range result {
		out = appendInt(out, d.u)
		out = append(out, ' ')
		out = appendInt(out, d.v)
		out = append(out, ' ')
		out = appendInt(out, d.s)
		out = append(out, '\n')
	}
	os.Stdout.Write(out)
}
