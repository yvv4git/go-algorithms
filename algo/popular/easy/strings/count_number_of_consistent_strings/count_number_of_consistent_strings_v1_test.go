package count_number_of_consistent_strings

import "testing"

func Test_countConsistentStringsV1(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				allowed: "ab",
				words:   []string{"ad", "bd", "aaab", "baa", "badab"},
			},
			want: 2,
			desc: `Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.`,
		},
		{
			name: "CASE-2",
			args: args{
				allowed: "abc",
				words:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			},
			want: 7,
			desc: `All strings are consistent.`,
		},
		{
			name: "CASE-3",
			args: args{
				allowed: "cad",
				words:   []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			},
			want: 4,
			desc: `Explanation: Strings "cc", "acd", "ac", and "d" are consistent.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsistentStringsV1(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStringsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
