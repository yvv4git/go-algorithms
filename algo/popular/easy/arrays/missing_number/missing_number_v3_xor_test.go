package missingnumber

import (
	"fmt"
	"testing"
)

func Test_missingNumberV3(t *testing.T) {
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
			if got := missingNumberV3(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------------------------
func TestResearch04(t *testing.T) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	a, b := 0, 0

	for i := 0; i <= len(nums); i++ {
		a = a ^ i
		t.Logf("a: %d  a_b: %08b", a, a)
	}

	for _, v := range nums {
		b = b ^ v
		t.Logf("b: %d  b_b: %08b", b, b)
	}

	t.Logf(">a: %d  a_b: %08b", a, a)
	t.Logf(">b: %d  b_b: %08b", b, b)
	result := a ^ b
	t.Logf("Result: %#v  result_b: %08b", result, result) // 8
}

func TestResearch05(t *testing.T) {
	a := 1
	b := 1
	r := a ^ b
	t.Logf("Result: %#v  result_b: %08b", r, r)
}

func TestResearch06(t *testing.T) {
	//for i := 0; i < 10; i++ {
	//	t.Logf("i=%d i_b=%08b", i, i)
	//}

	a := []int{1, 2, 3, 4, 5}
	for i, v := range a {
		r := i ^ v
		fmt.Printf("i=%d(%08b) ^ v=%d(%08b) = r=%d(%08b) \n", i, i, v, v, r, r)
	}
	/*
		i=0(00000000) ^ v=1(00000001) = r=1(00000001)
		i=1(00000001) ^ v=2(00000010) = r=3(00000011)
		i=2(00000010) ^ v=3(00000011) = r=1(00000001)
		i=3(00000011) ^ v=4(00000100) = r=7(00000111)
		i=4(00000100) ^ v=5(00000101) = r=1(00000001)
	*/
}

func TestResearch07(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := 0
	for i, v := range a {
		bOld := b
		b = b ^ v
		fmt.Printf("[%d] b=%d(%08b) ^ v=%d(%08b) = b=%d(%08b) \n", i, bOld, bOld, v, v, b, b)
	}
	/*
		i=0(00000000) ^ v=1(00000001) = r=1(00000001)
		i=1(00000001) ^ v=2(00000010) = r=3(00000011)
		i=2(00000010) ^ v=3(00000011) = r=1(00000001)
		i=3(00000011) ^ v=4(00000100) = r=7(00000111)
		i=4(00000100) ^ v=5(00000101) = r=1(00000001)
	*/
}
