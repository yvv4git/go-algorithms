package fibonacci

import "testing"

func Test_fibonacciV2(t *testing.T) {
	type args struct {
		n uint
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Test Case 1: n = 0",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "Test Case 2: n = 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Test Case 3: n = 2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "Test Case 4: n = 10",
			args: args{n: 10},
			want: 55,
		},
		{
			name: "Test Case 5: n = 20",
			args: args{n: 20},
			want: 6765,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciDP(tt.args.n); got != tt.want {
				t.Errorf("fibonacciDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
