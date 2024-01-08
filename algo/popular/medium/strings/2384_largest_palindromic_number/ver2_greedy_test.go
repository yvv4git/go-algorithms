package main

import "testing"

func Test_largestPalindromeV2(t *testing.T) {
	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				num: "444947137",
			},
			want: "7449447",
		},
		{
			name: "Test Case 2",
			args: args{
				num: "00009",
			},
			want: "9",
		},
		{
			name: "Test Case 3",
			args: args{
				num: "0000",
			},
			want: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPalindromeV2(tt.args.num); got != tt.want {
				t.Errorf("largestPalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
