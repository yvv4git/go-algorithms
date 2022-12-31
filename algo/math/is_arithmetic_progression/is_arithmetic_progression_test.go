package is_arithmetic_progression

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	type args struct {
		arr []int
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{3, 5, 1},
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int{1, 2, 4},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanMakeArithmeticProgression(tt.args.arr); got != tt.want {
				t.Errorf("CanMakeArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}
