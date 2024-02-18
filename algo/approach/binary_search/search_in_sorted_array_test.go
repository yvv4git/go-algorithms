package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Тест 1: Элемент присутствует в массиве",
			args: args{
				arr:    []int{2, 3, 4, 10, 40},
				target: 10,
			},
			want: 3, // Индекс элемента 10 в массиве
		},
		{
			name: "Тест 2: Элемент отсутствует в массиве",
			args: args{
				arr:    []int{2, 3, 4, 10, 40},
				target: 9,
			},
			want: -1, // Элемент 9 отсутствует в массиве, поэтому ожидаем -1
		},
		{
			name: "Тест 3: Массив состоит из одного элемента",
			args: args{
				arr:    []int{10},
				target: 10,
			},
			want: 0, // Элемент 10 единственный в массиве, его индекс 0
		},
		{
			name: "Тест 4: Массив состоит из нескольких одинаковых элементов",
			args: args{
				arr:    []int{10, 10, 10, 10, 10},
				target: 10,
			},
			want: 2, // Любой из индексов, где элемент 10, например, 2
		},
		{
			name: "Тест 5: Пустой массив",
			args: args{
				arr:    []int{},
				target: 10,
			},
			want: -1, // В пустом массиве элемент 10 отсутствует, ожидаем -1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
