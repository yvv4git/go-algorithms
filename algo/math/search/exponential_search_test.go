package search

import "testing"

func TestExponentialSearch(t *testing.T) {
	type args struct {
		list   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-2",
			args: args{
				list:   []int{0},
				target: 0,
			},
			want: 0,
		},
		{
			name: "CASE-3",
			args: args{
				list:   []int{0, 1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "CASE-4",
			args: args{
				list:   []int{0, 1, 2, 3, 4, 5},
				target: 4,
			},
			want: 4,
		},
		{
			name: "CASE-5",
			args: args{
				list:   []int{0, 1, 2, 3, 4, 5, 6},
				target: 15,
			},
			want: -1,
		},
		{
			name: "CASE-6",
			args: args{
				list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				target: 7,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exponentialSearch(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func exponentialSearch(arr []int, x int) int {
	/*
		Экспоненциальный поиск — это алгоритм, используемый для поиска в отсортированных, неограниченных/бесконечных массивах.
		Идея состоит в том, чтобы определить диапазон, в котором находится целевое значение, и выполнить бинарный поиск в пределах этого диапазона.
		Предполагая, что массив отсортирован в порядке возрастания, он ищет первый показатель, k, где значение 2k больше, чем ключ поиска.
		В настоящее время 2k а также 2k-1 становится верхней границей и нижней границей для алгоритма бинарного поиска соответственно.

		Экспоненциальный поиск занимает O(log(i))

		Он может даже превзойти бинарный поиск, когда цель находится в начале массива.
		Это связано с тем, что экспоненциальный поиск будет выполняться в O(log(i)) время, где i - это индекс элемента, который ищется в массиве,
		тогда как бинарный поиск будет выполняться в O(log(n)) время, где n это общее количество элементов в массиве.
	*/
	n := len(arr)
	if arr[0] == x {
		return 0
	}

	i := 1
	for i < n && arr[i] <= x {
		i *= 2
	}

	return binarySearch(arr, i/2, min(i, n-1), x)
}

func binarySearch(arr []int, left, right, x int) int {
	if right >= left {
		mid := left + (right-left)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return binarySearch(arr, left, mid-1, x)
		}
		return binarySearch(arr, mid+1, right, x)
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
