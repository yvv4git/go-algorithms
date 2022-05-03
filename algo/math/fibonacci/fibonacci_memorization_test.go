package fibonacci

import (
	"reflect"
	"testing"
)

func Test_fibSeriesMemoization(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 5,
			},
			want: []int{1, 1, 2, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := fibSeriesMemoization(tt.args.n); !reflect.DeepEqual(result, tt.want) {
				t.Errorf("fibSeriesMemoization() = %v, want %v", result, tt.want)
			}
		})
	}
}
