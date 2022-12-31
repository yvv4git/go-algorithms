package is_prefix_array

import "testing"

func TestIsPrefixArray(t *testing.T) {
	type args struct {
		s   string
		arr []string
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				s:   "iloveleetcode",
				arr: []string{"i", "love", "leetcode", "apples"},
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				s:   "iloveleetcode",
				arr: []string{"apples", "i", "love", "leetcode"},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrefixArray(tt.args.s, tt.args.arr); got != tt.want {
				t.Errorf("IsPrefixArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
