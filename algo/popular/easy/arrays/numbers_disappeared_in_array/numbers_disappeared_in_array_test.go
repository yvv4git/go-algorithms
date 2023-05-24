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

// ------------------------------

func abs(n int) int {
	if n > 0 {
		return n
	}

	return -n
}

func TestResearch01(t *testing.T) {
	// Имеем на входе неотсортированный массив целых чисел. Конкретно в этой задаче числа натуральные, т.е. нет 0.
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	t.Logf("a: %#v", a)
	for _, value := range a {
		// На месте, где должно стоять это число ставим минус
		idx := abs(value) - 1
		t.Logf("->idx(%d):%d", value, idx)
		if a[idx] > 0 {
			a[idx] = -a[idx]
		}
	}
	t.Logf("a: %#v", a)
}

func TestResearch02(t *testing.T) {
	/*
		Нельзя выходить за границы массива, т.е. обращаться к индексам за пределами массива.
	*/
	a := []int{1, 2, 3}
	t.Logf("a[1]: %v", a[1] > 0)
	t.Logf("a[3]: %v", a[3] > 0)
}
