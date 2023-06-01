package xor_operation_in_array

import (
	"testing"
)

func Test_xorOperationV1(t *testing.T) {
	type args struct {
		n     int
		start int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n:     5,
				start: 0,
			},
			want: 8,
		},
		{
			name: "CASE-2",
			args: args{
				n:     4,
				start: 3,
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperationV1(tt.args.n, tt.args.start); got != tt.want {
				t.Errorf("xorOperationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchM1Ex01(t *testing.T) {
	n := 5
	start := 0

	result := start

	/*
		По условию задачи, числа начинаются со nums[i] = start + 2 * i
	*/
	for i := 1; i < n; i++ {
		num := start + 2*i // Вычисляем num[i]
		result ^= num      // В задаче сказано, что это число потом надо XOR c предыдущим, т.е. с результатом
	}

	t.Logf("Result: %d", result)
}
