package is_sorted

import "testing"

func TestIsSorted(t *testing.T) {
	type args struct {
		nums []int
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{2, 1, 3, 4},
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSorted(tt.args.nums); got != tt.want {
				t.Errorf("IsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
