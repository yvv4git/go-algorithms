package two_numbers_matching

import (
	"reflect"
	"testing"
)

func Test_twoNumbersMatching(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				a: 564,
				b: 8954,
			},
			want: []int{5, 4},
		},
		{
			name: "CASE-2",
			args: args{
				a: 1234,
				b: 1263,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "CASE-3",
			args: args{
				a: 1234,
				b: 1763,
			},
			want: []int{1, 3},
		},
		{
			name: "CASE-4",
			args: args{
				a: 1234,
				b: 5678,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoNumbersMatching(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoNumbersMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
