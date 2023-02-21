package with_numbers

import (
	"reflect"
	"testing"
)

func Test_mergeV1(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "CASE-2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "CASE-3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeV1(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)

			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("mergeV1() = %v, but got = %v", tt.args.nums1, tt.want)
			}
		})
	}
}
