package decode_xored_array_v1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_decodeV1(t *testing.T) {
	type args struct {
		encoded []int
		first   int
	}

	tests := []struct {
		name string
		args args
		want []int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				encoded: []int{1, 2, 3},
				first:   1,
			},
			want: []int{1, 0, 2, 1},
			desc: `Explanation: If arr = [1,0,2,1], then first = 1 and encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]`,
		},
		{
			name: "CASE-2",
			args: args{
				encoded: []int{6, 2, 7, 3},
				first:   4,
			},
			want: []int{4, 2, 0, 7, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeV1(tt.args.encoded, tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchV1Ex01(t *testing.T) {
	encoded := []int{1, 2, 3}
	first := 1

	result := make([]int, len(encoded)+1)
	result[0] = first // В исходный (decoded) массив первым числом ставим first

	for idx, xorSum := range encoded {
		/*
			Т.е. у нас было 1 первым числом исходного массива, следовательно, XOR-им ее с полученным в encoded[0] значением и получаем 0.
		*/
		result[idx+1] = xorSum ^ result[idx] // Из условия задачи - каждый элемент encoded массива = XOR c предыдущим
	}

	t.Logf("result: %v", result)
}

func TestResearchV1Ex02(t *testing.T) {
	plainArr := []int{1, 0, 2, 1}
	fmt.Printf("arr: %#v \n", plainArr)

	//fmt.Printf("r: %d \n", 0^1)
	encoded := make([]int, len(plainArr)-1)
	for i := 0; i < len(plainArr)-1; i++ {
		encoded[i] = plainArr[i] ^ plainArr[i+1]
	}

	fmt.Printf("encoded: %#v \n", encoded) // encoded: []int{1, 2, 3}
}
