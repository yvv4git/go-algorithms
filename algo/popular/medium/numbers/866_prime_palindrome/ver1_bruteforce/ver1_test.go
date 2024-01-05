package ver1_bruteforce

import "testing"

func Test_primePalindromeV1(t *testing.T) {
	type args struct {
		N int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				N: 6,
			},
			want: 7,
		},
		{
			name: "Test Case 2",
			args: args{
				N: 8,
			},
			want: 11,
		},
		{
			name: "Test Case 3",
			args: args{
				N: 13,
			},
			want: 101,
		},
		//{
		//	name: "Test Case 4",
		//	args: args{
		//		N: 100000000,
		//	},
		//	want: 1000000300000030000001,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primePalindromeV1(tt.args.N); got != tt.want {
				t.Errorf("primePalindromeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
