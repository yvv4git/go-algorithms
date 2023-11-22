package _252_meeting_rooms

import "sort"

type Interval struct {
	Start int
	End   int
}

type IntervalSlice []Interval

func (s IntervalSlice) Len() int {
	return len(s)
}

func (s IntervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntervalSlice) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func canAttendMeetingsV1(intervals []Interval) bool {
	/*
		Method: Loop after sort.
		Time complexity: O(n log n)
		Space complexity: O(1)

		Асимптотическая сложность(Time complexity) алгоритма зависит от того, как вы реализуете сортировку.
		Если вы используете стандартную функцию сортировки в Go (sort.Sort), то временная сложность этого алгоритма будет O(n log n),
		где n - количество встреч. Это связано с тем, что сортировка в Go реализована с использованием алгоритма сортировки слиянием, который имеет сложность O(n log n).
		После сортировки мы проходим по всем встречам, поэтому временная сложность этого алгоритма будет O(n), где n - количество встреч.
		Общая асимптотическая сложность будет O(n log n) + O(n), которая, учитывая большую роль большего слагаемого, будет O(n log n).
	*/
	sort.Sort(IntervalSlice(intervals))

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}

	return true
}
