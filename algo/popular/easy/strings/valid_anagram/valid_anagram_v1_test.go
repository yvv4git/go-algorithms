package valid_anagram

import "testing"

func Test_isAnagramV1(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagramV1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
