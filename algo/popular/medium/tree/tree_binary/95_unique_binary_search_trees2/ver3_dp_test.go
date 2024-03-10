package main

import (
	"testing"
)

func Test_generateTreesV3(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int // Изменен на int, чтобы проверить количество деревьев
	}{
		{
			name: "Test Case 1",
			args: args{n: 1},
			want: 1, // Для n = 1 должно быть 1 дерево
		},
		{
			name: "Test Case 2",
			args: args{n: 2},
			want: 2, // Для n = 2 должно быть 2 дерева
		},
		{
			name: "Test Case 3",
			args: args{n: 3},
			want: 5, // Для n = 3 должно быть 5 деревьев
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTreesV3(tt.args.n); len(got) != tt.want {
				t.Errorf("generateTrees() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
