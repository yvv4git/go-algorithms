package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		sw []int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				sw: []int{1, 2, 3, 4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.sw)
			t.Log(tt.args.sw)
		})
	}
}
