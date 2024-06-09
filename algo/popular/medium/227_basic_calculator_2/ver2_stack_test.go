package main

import "testing"

func Test_calculateV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: Addition",
			args: args{s: "1+2+3"},
			want: 6,
		},
		{
			name: "Test case 2: Subtraction",
			args: args{s: "5-3-1"},
			want: 1,
		},
		{
			name: "Test case 3: Multiplication",
			args: args{s: "2*3*4"},
			want: 24,
		},
		{
			name: "Test case 4: Division",
			args: args{s: "10/2/2"},
			want: 2,
		},
		{
			name: "Test case 5: Mix of operations",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: "Test case 6: Division and subtraction",
			args: args{s: "10/2-1"},
			want: 4,
		},
		{
			name: "Test case 7: Addition with negative numbers",
			args: args{s: "-1+2+3"},
			want: 4,
		},
		{
			name: "Test case 9: Multiplication with zero",
			args: args{s: "0*10"},
			want: 0,
		},
		{
			name: "Test case 10: Division by one",
			args: args{s: "100/1"},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateV2(tt.args.s); got != tt.want {
				t.Errorf("calculateV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
