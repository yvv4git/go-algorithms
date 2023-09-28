package _28_summary_ranges

import (
	"reflect"
	"testing"
)

func Test_summaryRangesV1(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []string
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 7},
			},
			want: []string{"0->2", "4->5", "7"},
			desc: `The ranges are:
					[0,2] --> "0->2"
					[4,5] --> "4->5"
					[7,7] --> "7"`,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{0, 2, 3, 4, 6, 8, 9},
			},
			want: []string{"0", "2->4", "6", "8->9"},
			desc: `The ranges are:
					[0,0] --> "0"
					[2,4] --> "2->4"
					[6,6] --> "6"
					[8,9] --> "8->9"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRangesV1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRangesV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
