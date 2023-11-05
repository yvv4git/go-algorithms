package _05_isomorphic_strings

import "testing"

func Test_isIsomorphicV2(t *testing.T) {
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
				s: "",
				t: "",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "CASE-4",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphicV2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphicV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
