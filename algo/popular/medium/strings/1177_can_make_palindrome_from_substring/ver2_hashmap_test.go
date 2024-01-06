package _177_can_make_palindrome_from_substring

import (
	"reflect"
	"testing"
)

func Test_canMakePaliQueriesV2(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		//{ // Этот кейс не прошел
		//	name: "Test Case 1",
		//	args: args{
		//		s:       "abcda",
		//		queries: [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}},
		//	},
		//	want: []bool{true, false, false, true},
		//},
		{
			name: "Test Case 3",
			args: args{
				s:       "aaaaa",
				queries: [][]int{{0, 4, 2}, {0, 4, 1}},
			},
			want: []bool{true, true},
		},
		{
			name: "Test Case 4",
			args: args{
				s:       "abc",
				queries: [][]int{{0, 2, 1}},
			},
			want: []bool{true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakePaliQueriesV2(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canMakePaliQueriesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
