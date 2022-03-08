package fibonacci

import "testing"

func Test_fibonacciStruct(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "CASE-2",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "CASE-3",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "CASE-4",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "CASE-5",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciByStruct(tt.args.n); got != tt.want {
				t.Errorf("fibonacciStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
