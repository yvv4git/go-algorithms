package main

import (
	"reflect"
	"testing"
)

func Test_weirdAlgorithm(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Тест 1: n = 1",
			args: args{n: 1},
			want: []int{1},
		},
		{
			name: "Тест 2: n = 2",
			args: args{n: 2},
			want: []int{2, 1},
		},
		{
			name: "Тест 3: n = 3",
			args: args{n: 3},
			want: []int{3, 10, 5, 16, 8, 4, 2, 1},
		},
		{
			name: "Тест 4: n = 4",
			args: args{n: 4},
			want: []int{4, 2, 1},
		},
		{
			name: "Тест 5: n = 5",
			args: args{n: 5},
			want: []int{5, 16, 8, 4, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := weirdAlgorithm(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("weirdAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
