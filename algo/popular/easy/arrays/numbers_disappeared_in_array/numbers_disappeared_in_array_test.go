package numbersdisappearedinarray

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{},
			},
			want: []int{}, // числа, которые отсутствуют в последовательности
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{5, 6}, // числа, которые отсутствуют в последовательности
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{2}, // числа, которые отсутствуют в последовательности
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
