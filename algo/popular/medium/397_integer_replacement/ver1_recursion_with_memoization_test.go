package main

import "testing"

func Test_integerReplacement(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: n = 1",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "Test case 2: n = 2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "Test case 3: n = 3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "Test case 4: n = 4",
			args: args{n: 4},
			want: 2,
		},
		{
			name: "Test case 5: n = 8",
			args: args{n: 8},
			want: 3,
		},
		{
			name: "Test case 6: n = 7",
			args: args{n: 7},
			want: 4,
		},
		{
			name: "Test case 7: n = 15",
			args: args{n: 15},
			want: 5,
		},
		{
			name: "Test case 8: n = 65535",
			args: args{n: 65535},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
