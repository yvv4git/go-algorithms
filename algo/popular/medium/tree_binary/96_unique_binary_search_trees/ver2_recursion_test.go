package main

import "testing"

func Test_numTreesV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{n: 3},
			want: 5,
		},
		{
			name: "Test Case 4",
			args: args{n: 4},
			want: 14,
		},
		{
			name: "Test Case 5",
			args: args{n: 5},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTreesV2(tt.args.n); got != tt.want {
				t.Errorf("numTreesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
