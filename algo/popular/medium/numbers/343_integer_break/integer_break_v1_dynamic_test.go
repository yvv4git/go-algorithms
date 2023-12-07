package _43_integer_break

import "testing"

func Test_integerBreakV1(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "n = 3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "n = 4",
			args: args{n: 4},
			want: 4,
		},
		{
			name: "n = 5",
			args: args{n: 5},
			want: 6,
		},
		{
			name: "n = 6",
			args: args{n: 6},
			want: 9,
		},
		{
			name: "n = 7",
			args: args{n: 7},
			want: 12,
		},
		{
			name: "n = 8",
			args: args{n: 8},
			want: 18,
		},
		{
			name: "n = 9",
			args: args{n: 9},
			want: 27,
		},
		{
			name: "n = 10",
			args: args{n: 10},
			want: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreakV1(tt.args.n); got != tt.want {
				t.Errorf("integerBreakV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
