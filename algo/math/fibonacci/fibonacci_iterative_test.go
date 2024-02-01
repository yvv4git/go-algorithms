package fibonacci

import "testing"

func Test_fibonacciV4(t *testing.T) {
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
		{
			name: "Test Case 6: n = 30",
			args: args{n: 30},
			want: 832040,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciIterative(tt.args.n); got != tt.want {
				t.Errorf("fibonacciIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
