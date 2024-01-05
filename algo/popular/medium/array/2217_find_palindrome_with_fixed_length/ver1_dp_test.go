package _217_find_palindrome_with_fixed_length

import (
	"reflect"
	"testing"
)

func Test_kthPalindromeV1(t *testing.T) {
	type args struct {
		queries   []int
		intLength int
	}

	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Test case 1",
			args: args{
				queries:   []int{1, 2, 3, 4, 5},
				intLength: 3,
			},
			want: []int64{101, 111, 121, 131, 141},
		},
		{
			name: "Test case 2",
			args: args{
				queries:   []int{1, 2, 3, 4, 5},
				intLength: 4,
			},
			want: []int64{1001, 1111, 1221, 1331, 1441},
		},
		{
			name: "Test case 3",
			args: args{
				queries:   []int{1, 2, 3, 4, 5},
				intLength: 5,
			},
			want: []int64{10001, 10101, 10201, 10301, 10401},
		},
		{
			name: "Test case 4",
			args: args{
				queries:   []int{1, 2, 3, 4, 5},
				intLength: 6,
			},
			want: []int64{100001, 101101, 102201, 103301, 104401},
		},
		{
			name: "Test case 5",
			args: args{
				queries:   []int{1, 2, 3, 4, 5, 90},
				intLength: 3,
			},
			want: []int64{101, 111, 121, 131, 141, 999},
		},
		{
			name: "Test case 6",
			args: args{
				queries:   []int{2, 4, 6},
				intLength: 4,
			},
			want: []int64{1111, 1331, 1551},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthPalindromeV1(tt.args.queries, tt.args.intLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthPalindromeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
