package max_consecutive_one

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}

	testCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{1, 1, 0, 1, 1, 1},
			},
			want: 3,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 0, 1, 1, 0, 1},
			},
			want: 2,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "CASE-4",
			args: args{
				nums: []int{0, 0, 0, 0, 0},
			},
			want: 0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if result := findMaxConsecutiveOnes(tt.args.nums); result != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", result, tt.want)
			}
		})
	}
}
