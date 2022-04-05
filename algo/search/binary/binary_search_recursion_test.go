package binary

import "testing"

func Test_binarySearchRecursion(t *testing.T) {
	type args struct {
		list   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				list:   []int{},
				target: 100500,
			},
			want: -1,
		},
		{
			name: "CASE-2",
			args: args{
				list:   []int{0},
				target: 0,
			},
			want: 0,
		},
		{
			name: "CASE-3",
			args: args{
				list:   []int{0, 1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "CASE-4",
			args: args{
				list:   []int{0, 1, 2, 3, 4, 5},
				target: 4,
			},
			want: 4,
		},
		{
			name: "CASE-5",
			args: args{
				list:   []int{0, 1, 2, 3, 4, 5, 6},
				target: 15,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchRecursion(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("binarySearchRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
