package top_k_frequently_elements

import (
	"reflect"
	"testing"
)

func TestCalcTopFreqK(t *testing.T) {
	type args struct {
		a []int
		k int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				a: []int{1, 1, 1, 1, 2, 3, 3, 3, 4, 4, 1},
				k: 3,
			},
			want: []int{1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequentV1(tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequentV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
