package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		ar []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				ar: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "CASE-2",
			args: args{
				ar: []int{1, 3, 2, 5, 4},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.ar)
			if !reflect.DeepEqual(tt.want, tt.args.ar) {
				t.Errorf("We expect value %v, but we got %v", tt.want, tt.args.ar)
			}
		})
	}
}
