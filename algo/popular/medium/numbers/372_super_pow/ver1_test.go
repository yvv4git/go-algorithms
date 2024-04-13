package main

import (
	"fmt"
	"testing"
)

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "Test Case 2",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "Test Case 3",
			args: args{
				a: 1,
				b: []int{4, 3, 2, 1},
			},
			want: 1,
		},
		{
			name: "Test Case 4",
			args: args{
				a: 2147483647,
				b: []int{2, 0, 0},
			},
			want: 1198,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResearch01(t *testing.T) {
	a := 1
	mod := 1337
	a %= mod
	fmt.Println("a", a) // 1

	a = 5
	a %= mod
	fmt.Println("a", a) // 5
}
