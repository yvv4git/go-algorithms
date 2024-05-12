package main

import "testing"

func Test_maxCount(t *testing.T) {
	type args struct {
		m   int
		n   int
		ops [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "No operations",
			args: args{
				m:   3,
				n:   3,
				ops: [][]int{},
			},
			want: 9, // Все элементы массива заполнены нулями, поэтому максимальное число - это произведение m и n
		},
		{
			name: "Single operation",
			args: args{
				m:   3,
				n:   3,
				ops: [][]int{{2, 2}},
			},
			want: 4, // Одна операция добавления на диапазон 2x2, поэтому максимальное число - это произведение 2 и 2
		},
		{
			name: "Multiple operations",
			args: args{
				m:   3,
				n:   3,
				ops: [][]int{{2, 2}, {3, 3}},
			},
			want: 4, // Две операции добавления на диапазоны 2x2 и 3x3, поэтому максимальное число - это произведение 2 и 2
		},
		{
			name: "Larger array",
			args: args{
				m:   10,
				n:   10,
				ops: [][]int{{5, 5}, {7, 7}},
			},
			want: 25, // Две операции добавления на диапазоны 5x5 и 7x7, поэтому максимальное число - это произведение 5 и 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.args.m, tt.args.n, tt.args.ops); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
