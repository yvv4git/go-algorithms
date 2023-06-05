package find_difference

import "testing"

func Test_findTheDifferenceV1(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "CASE-1",
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: 'e', // 101
		},
		{
			name: "CASE-2",
			args: args{
				s: "",
				t: "y",
			},
			want: 'y', // 121
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifferenceV1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifferenceV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
