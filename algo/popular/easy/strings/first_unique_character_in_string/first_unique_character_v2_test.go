package first_unique_character_in_string

import "testing"

func Test_firstUniqCharV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "CASE-2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "CASE-3",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqCharV2(tt.args.s); got != tt.want {
				t.Errorf("firstUniqCharV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
