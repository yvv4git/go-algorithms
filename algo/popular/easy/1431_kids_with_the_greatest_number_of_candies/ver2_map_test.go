package main

import (
	"reflect"
	"testing"
)

func Test_kidsWithCandiesMapV2(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "CASE-1",
			args: args{
				candies:      []int{2, 3, 5, 1, 3},
				extraCandies: 3,
			},
			want: []bool{true, true, true, false, true},
		},
		{
			name: "CASE-2",
			args: args{
				candies:      []int{4, 2, 1, 1, 2},
				extraCandies: 1,
			},
			want: []bool{true, false, false, false, false},
		},
		{
			name: "CASE-3",
			args: args{
				candies:      []int{12, 1, 12},
				extraCandies: 10,
			},
			want: []bool{true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandiesMapV2(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kidsWithCandiesMapV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
