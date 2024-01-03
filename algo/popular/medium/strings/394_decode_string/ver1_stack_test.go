package _94_decode_string

import "testing"

func Test_decodeStringV1(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				str: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "CASE-2",
			args: args{
				str: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "CASE-3",
			args: args{
				str: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeStringStackV1(tt.args.str); got != tt.want {
				t.Errorf("decodeStringStackV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
