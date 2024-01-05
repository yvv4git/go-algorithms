package _31_palindrome_partitioning

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test case 1",
			args: args{s: "aab"},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "Test case 2",
			args: args{s: "a"},
			want: [][]string{{"a"}},
		},
		{
			name: "Test case 5",
			args: args{s: "bb"},
			want: [][]string{{"b", "b"}, {"bb"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionV1(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
