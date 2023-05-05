package monotonic_array

import "testing"

func Test_isMonotonic(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{1, 2, 2, 3},
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{6, 5, 4, 4},
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{1, 3, 2},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.nums); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
