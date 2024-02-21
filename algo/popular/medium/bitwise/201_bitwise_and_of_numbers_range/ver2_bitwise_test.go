package main

import "testing"

func Test_rangeBitwiseAndV2(t *testing.T) {
	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAndV2(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAndV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
