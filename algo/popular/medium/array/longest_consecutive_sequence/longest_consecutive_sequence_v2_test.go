package longest_consecutive_sequence

import "testing"

func Test_longestConsecutiveV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutiveV2(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutiveV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
