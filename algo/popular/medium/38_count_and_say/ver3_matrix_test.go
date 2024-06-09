package _8_count_and_say

import "testing"

func Test_countAndSayV3(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSayV3(tt.args.n); got != tt.want {
				t.Errorf("countAndSayV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
