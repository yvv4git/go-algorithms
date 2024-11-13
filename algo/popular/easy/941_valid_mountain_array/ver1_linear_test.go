package main

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{2, 1},
			},
			want: false,
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				arr: []int{0, 3, 2, 1},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
