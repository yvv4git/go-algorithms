package power_of_two

import "testing"

func Test_isPowerOfTwoV3(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				n: 1,
			},
			want: true,
			desc: "2^0 = 1",
		},
		{
			name: "CASE-2",
			args: args{
				n: 16,
			},
			want: true,
			desc: "2^4 = 16",
		},
		{
			name: "CASE-3",
			args: args{
				n: 3,
			},
			want: false,
			desc: "No found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwoV3(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwoV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------
func TestResearchBitwise001(t *testing.T) {
	n := 8 // 00001000

	// Step-1: вычитаем 1 из числа.
	x := n - 1 // 00000111
	t.Logf("n = %d(%08b) n - x=%d(%08b)", n, n, x, x)

	// Step-2: умножаем на n. (Bitwise AND)
	// Где в обоих операндах 1, на выходе тоже будет 1.
	// В остальных случаях 0.
	y := n & x
	t.Logf("n = %d(%08b) x=%d(%08b) n & x=%d(%08b)", n, n, x, x, y, y)
}
