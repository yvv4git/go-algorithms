package basis

import (
	"testing"
)

func Test_reverseInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "CASE-2",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "CASE-3",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "CASE-4",
			args: args{
				x: 123456789,
			},
			want: 987654321,
		},
		{
			name: "NEGATIVE-1",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "x^31 - int32",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInt(tt.args.x); got != tt.want {
				t.Errorf("reverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExploreReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "CASE-1",
			x:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := tt.x   // число, которое надо развернуть
			result := 0 // здесь будет результат

			for x > 0 {
				result = result*10 + x%10
				t.Logf("%d*10 + %d%%10", x, x)
				t.Logf("%d + %d", x/10, x%10)
				x = x / 10
			}

			t.Logf("Result: %d", result)
		})
	}
}
