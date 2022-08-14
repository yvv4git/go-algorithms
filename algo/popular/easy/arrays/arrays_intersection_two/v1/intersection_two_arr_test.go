package v1

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			want: []int{},
		},
		{
			name: "CASE-2",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2, 2},
		},
		{
			name: "CASE-3",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{4, 9},
		},
		{
			name: "CASE-4",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6},
			},
			want: []int{},
		},
		{
			name: "CASE-5",
			args: args{
				nums1: []int{1, 2, 3, 4, 5},
				nums2: []int{4, 5, 6},
			},
			want: []int{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpSlice(t *testing.T) {
	s := []int{4, 5, 3, 4, 5}
	t.Logf("%v", s[:2]) // до idx=2 исключительно, т.е. без элемента с idx2
}
