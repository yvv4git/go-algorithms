package missingnumber

import "testing"

func Test_missingNumberV1(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{},
			},
			want: 0,
			desc: "When we have empty slice",
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
			desc: "We have n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums",
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
			desc: "We have n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums",
		},
		{
			name: "CASE-4",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
			desc: "We have n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumberV1(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ----------------------------------------
func TestResearch01(t *testing.T) {
	/*
		Задача решается элементарно!
		Т.е. у нас есть:
		- числа от 0 до n
		- числа уникальные, т.е. 0, 1, 2, 3, 4, 5..., n

		Например, было [0, 1, 2, 3], затем 2 удалили и осталось [0, 1, 3]. Как вычислить 2?
		Можно итерируясь по индексам посчитать, какая должна быть сумма чисел: ind += i+1, 0 + 1 + 2 + 3 = 6.
		Затем посчитать текущую сумму элементов: 0 + 1 + 3 = 4.
		Следовательно, какого числа не хватает? 6-4 = 2.
	*/
	a := []int{3, 0, 1}
	t.Logf("a: %#v", a)

	var sum int
	var ind int

	for i, v := range a {
		t.Logf("i=%d v=%d", i, v)
		sum += v
		ind += i + 1
	}

	result := ind - sum
	t.Logf("result: %#v", result)
}
