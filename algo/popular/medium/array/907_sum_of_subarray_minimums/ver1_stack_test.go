package _07_sum_of_subarray_minimums

import "testing"

func Test_sumSubarrayMinsV1(t *testing.T) {
	type args struct {
		A []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				A: []int{3, 1, 2, 4},
			},
			want: 17,
		},
		{
			name: "Test Case 2",
			args: args{
				A: []int{1, 2, 3, 4},
			},
			want: 20,
		},
		{
			name: "Test Case 3",
			args: args{
				A: []int{4, 3, 2, 1},
			},
			want: 20,
		},
		{
			name: "Test Case 4",
			args: args{
				A: []int{1, 1, 1, 1},
			},
			want: 10,
		},
		//{
		//	name: "Test Case 5",
		//	args: args{
		//		A: []int{3, 1, 2, 4, 1},
		//	},
		//	want: 25,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMinsV1(tt.args.A); got != tt.want {
				t.Errorf("sumSubarrayMinsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
