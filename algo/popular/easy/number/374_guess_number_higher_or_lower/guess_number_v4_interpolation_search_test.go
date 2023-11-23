package _74_guess_number_higher_or_lower

import "testing"

func Test_interpolationSearch(t *testing.T) {
	type args struct {
		key   int
		array []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				key:   6,
				array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 5, // Видимо индекс элемента в массиве
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interpolationSearch(tt.args.key, tt.args.array); got != tt.want {
				t.Errorf("interpolationSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
