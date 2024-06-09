package main

import "testing"

func Test_checkPalindromeFormationV2(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				a: "abc",
				b: "cba",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				a: "abcd",
				b: "dcba",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				a: "abcd",
				b: "abce",
			},
			want: false,
		},
		{
			name: "Test case 6",
			args: args{
				a: "x",
				b: "y",
			},
			want: true,
		},
		{
			name: "Test case 7",
			args: args{
				a: "xbdef",
				b: "xecab",
			},
			want: false,
		},
		{
			name: "Test case 8",
			args: args{
				a: "ulacfd",
				b: "jizalu",
			},
			want: true,
		},
		{
			name: "Test case 9",
			args: args{
				a: "abdef",
				b: "fecab",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindromeFormationV2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("checkPalindromeFormationV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
