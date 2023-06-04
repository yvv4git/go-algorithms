package set_mismatch

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
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
				nums: []int{1, 2, 2, 4},
			},
			want: []int{2, 3},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{3, 2, 2},
			},
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchV3Ex01(t *testing.T) {
	x := 4
	y := 5
	r := x ^ y
	t.Logf("%d(%08b) ^ %d(%08b) = %d(%08b)", x, x, y, y, r, r) // 4(00000100) ^ 5(00000101) = 1(00000001)

	x = 3
	y = 5
	r = x ^ y
	t.Logf("%d(%08b) ^ %d(%08b) = %d(%08b)", x, x, y, y, r, r) // 3(00000011) ^ 5(00000101) = 6(00000110)

	r = 1 ^ 2 ^ 3 ^ 4 ^ 5
	t.Logf("1 ^ 2 ^ 3 ^ 4 ^ 5 = %d(%08b)", r, r) // 1 ^ 2 ^ 3 ^ 4 ^ 5 = 1(00000001)

	r = 1 ^ 2 ^ 3 ^ 4 ^ 5 ^ 6
	t.Logf("1 ^ 2 ^ 3 ^ 4 ^ 5 ^ 6 = %d(%08b)", r, r) // 1 ^ 2 ^ 3 ^ 4 ^ 5 ^ 6 = 7(00000111)

	r = 1 ^ 2 ^ 3 ^ 5
	t.Logf("1 ^ 2 ^ 3 ^ 5 = %d(%08b)", r, r) // 1 ^ 2 ^ 3 ^ 5 = 5(00000101)

	/*
		В данном примере видна закономерность - если XOR двух последовательностей, то получим отсутствующий элемент.
		В данном случае отсутствует 4.
	*/
	r1 := 1 ^ 2 ^ 3 ^ 5
	r2 := 1 ^ 2 ^ 3 ^ 4 ^ 5
	r = r1 ^ r2
	t.Logf("(1 ^ 2 ^ 3 ^ 5) ^ (1 ^ 2 ^ 3 ^ 4 ^ 5) = %d(%08b)", r, r) // (1 ^ 2 ^ 3 ^ 5) ^ (1 ^ 2 ^ 3 ^ 4 ^ 5) = 4(00000100)
}
