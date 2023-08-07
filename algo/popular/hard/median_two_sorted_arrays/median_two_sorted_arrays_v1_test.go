package median_two_sorted_arrays

import "testing"

func Test_findMedianSortedArraysV1(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "CASE-1",
			args: args{
				a: []int{1, 3},
				b: []int{2},
			},
			want: 2,
		},
		{
			name: "CASE-1",
			args: args{
				a: []int{1, 2},
				b: []int{3, 4},
			},
			want: 2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArraysV1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findMedianSortedArraysV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
