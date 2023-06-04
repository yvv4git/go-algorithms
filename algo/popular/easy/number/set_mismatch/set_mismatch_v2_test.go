package set_mismatch

import (
	"reflect"
	"testing"
)

func Test_findErrorNumsV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{1, 2, 2, 4},
			},
			want: []int{2, 3},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{3, 2, 2},
			},
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNumsV2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNumsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
