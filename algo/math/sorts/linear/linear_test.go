package linear

import (
	"reflect"
	"testing"
)

func TestLinearSort(t *testing.T) {
	type args struct {
		payload []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				payload: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LinearSortBigOrder(tt.args.payload)
			if !reflect.DeepEqual(tt.args.payload, tt.want) {
				t.Errorf("LinearSort() = %v, want %v", tt.args.payload, tt.want)
			}
		})
	}
}

func TestLinearSortLittleOrder(t *testing.T) {
	type args struct {
		list []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				list: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			LinearSortLittleOrder(tt.args.list)
			if !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("LinearSort() = %v, want %v", tt.args.list, tt.want)
			}
		})
	}
}
