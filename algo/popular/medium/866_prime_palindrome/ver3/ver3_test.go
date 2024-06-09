package ver3

import "testing"

func Test_primePalindrome(t *testing.T) {
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
			args: args{
				n: 6,
			},
			want: 7,
		},
		{
			name: "Test Case 2",
			args: args{
				n: 8,
			},
			want: 11,
		},
		{
			name: "Test Case 3",
			args: args{
				n: 13,
			},
			want: 101,
		},
		{
			name: "Test case 4",
			args: args{
				n: 930,
			},
			want: 10301,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primePalindrome(tt.args.n); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
