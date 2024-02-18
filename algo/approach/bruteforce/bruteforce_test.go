package main

import "testing"

func Test_bruteForce(t *testing.T) {
	type args struct {
		target  int
		numbers []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				target:  5,
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				target:  10,
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				target:  3,
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				target:  0,
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				target:  3,
				numbers: []int{},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bruteForce(tt.args.target, tt.args.numbers); got != tt.want {
				t.Errorf("bruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
