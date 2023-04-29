package count_bits

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 2,
			},
			want: []int{0, 1, 1},
		},
		{
			name: "CASE-2",
			args: args{
				n: 5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := countBitsV1(tt.args.n); !reflect.DeepEqual(result, tt.want) {
				t.Errorf("countBitsV1() = %v, want %v", result, tt.want)
			}
		})
	}
}

// ----------------------------------------
func TestResearchCountBits001(t *testing.T) {
	n := 2
	result := make([]int, n+1)

	for x := 0; x <= n; x++ {
		fmt.Printf("n=%d(%08b) x=%d\n", n, n, x)
		result[x] = result[x>>1] + x&1
	}
}

func TestResearchCountBits002(t *testing.T) {
	n := 1383
	result := make([]int, n+1)

	for x := 0; x <= n; x++ {
		result[x] = result[x>>1] + x&1
		fmt.Printf("n=%d(%08b) x=%d\n x>>1=%d x&1=%d", n, n, x, x>>1, x&1)
	}
}

func TestResearchCountBits003(t *testing.T) {
	x := 100
	t.Logf("x=%d(%08b)", x, x) // x=100(01100100)

	x2 := x >> 1 // Эта операция эквивалентна делению на 2.
	odd2 := x2 & 1
	t.Logf("x2=%d(%08b) odd=%d", x2, x2, odd2) // x2=50(00110010) odd=0

	x3 := x2 >> 1
	odd3 := x3 & 1
	t.Logf("x3=%d(%08b) odd=%d", x3, x3, odd3) // x3=25(00011001) odd=1

	x4 := x3 >> 1
	odd4 := x4 & 1
	t.Logf("x4=%d(%08b) odd=%d", x4, x4, odd4) // x4=12(00001100) odd=0
}

func TestResearchCountBits004(t *testing.T) {
	for i := 0; i < 10; i++ {
		isOdd := i & 1
		t.Logf("%d is odd: %v", i, isOdd)
	}
}
