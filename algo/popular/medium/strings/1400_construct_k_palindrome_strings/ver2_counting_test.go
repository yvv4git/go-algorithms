package _400_construct_k_palindrome_strings

import "testing"

func Test_canConstructV2(t *testing.T) {
	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Тест 1",
			args: args{
				s: "annabelle",
				k: 2,
			},
			want: true,
		},
		{
			name: "Тест 2",
			args: args{
				s: "leetcode",
				k: 3,
			},
			want: false,
		},
		{
			name: "Тест 3",
			args: args{
				s: "true",
				k: 4,
			},
			want: true,
		},
		{
			name: "Тест 4",
			args: args{
				s: "yzyzyzyzyzyzy",
				k: 2,
			},
			want: true,
		},
		//{
		//	name: "Тест 5",
		//	args: args{
		//		s: "cr",
		//		k: 7,
		//	},
		//	want: false,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstructV2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("canConstructV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
