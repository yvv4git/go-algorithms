package squares_sorted_array

import (
	"reflect"
	"testing"
)

func Test_sortedSquaresV3(t *testing.T) {
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
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquaresV3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquaresV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------

func TestResearch01(t *testing.T) {
	a := []int{-4, -1, 0, 3, 10}

	res := make([]int, len(a))
	left, right := 0, len(a)-1
	for i := 0; left <= right; i++ {
		t.Logf("l=%d r=%d i=%d", left, right, i)
		if abs(a[left]) > abs(a[right]) {
			res[len(a)-1-i] = a[left] * a[left]
			left++
		} else {
			res[len(a)-1-i] = a[right] * a[right]
			right--
		}

		if i > 100 {
			return
		}
	}

	t.Logf("--->res: %#v", res)
}

func TestResearch02(t *testing.T) {
	a := []int{-4, -1, 0, 3, 10}
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * a[i]
	}

	l, r := 0, len(a)-1
	i := 0
	for l < r {
		if (abs(a[l])) > (abs(a[r])) {
			a[l], a[r] = a[r], a[l]
			l++
		} else {
			a[l], a[r] = a[l], a[r]
			r--
		}

		i++
		if i > 100 {
			return
		}
	}

	t.Logf("a: %#v", a)
}

func TestResearch03(t *testing.T) {
	// Попробую сначала все числа возвести в квадраты.
	a := []int{-4, -1, 0, 3, 10}
	t.Logf("a: %#v", a)
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * a[i]
	}
	t.Logf("a: %#v", a)

	// Затем их отсортировать.
	/*
		Почему это будет работать?
		- сравниваем крайние элементы. Так как самое большое отрицательное число слева, может оказаться больше самого крайнего справа.
		- выбираем из них самое большое и вставляем в новый массив
	*/
	res := make([]int, len(a))
	l, r := 0, len(a)-1
	for i := 0; l <= r; i++ {
		if a[l] > a[r] {
			res[len(a)-1-i] = a[l]
			t.Logf("[1] i=%d l=%d[%d] r=%d[%d] len(a)-1-i=%d res=%#v", i, l, a[l], r, a[r], len(a)-1-i, res)
			l++
		} else {
			res[len(a)-1-i] = a[r]
			t.Logf("[2] i=%d l=%d[%d] r=%d[%d] len(a)-1-i=%d res=%#v", i, l, a[l], r, a[r], len(a)-1-i, res)
			r--
		}
	}

	t.Logf("res: %#v", res) // OK []int{0, 1, 9, 16, 100}
}

func TestResearch04(t *testing.T) {
	// Попробую сначала все числа возвести в квадраты.
	a := []int{-4, -1, 0, 3, 10}
	t.Logf("a: %#v", a)
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * a[i]
	}
	t.Logf("a: %#v", a)

	// Затем их отсортировать.
	res := make([]int, len(a))
	l, r := 0, len(a)-1
	for i := 0; l <= r; i++ {
		if a[l] < a[r] {
			res[i] = a[l]
			t.Logf("[1] i=%d l=%d[%d] r=%d[%d] res=%#v", i, l, a[l], r, a[r], res)
			r--
		} else {
			res[i] = a[r]
			t.Logf("[2] i=%d l=%d[%d] r=%d[%d] res=%#v", i, l, a[l], r, a[r], res)
			l++
		}
	}

	// Не работает...
	// Забавно, результат получился отсортирован в обратном порядке.
	t.Logf("res: %#v", res) // []int{16, 9, 1, 0, 0}
}
