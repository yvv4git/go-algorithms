package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomListV2(t *testing.T) {
	type args struct {
		head *Node
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Empty list",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "Single node list",
			args: args{head: &Node{Val: 1}},
			want: &Node{Val: 1},
		},
		{
			name: "Multiple node list with no random pointers",
			args: args{head: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}},
			want: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomListV2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomListV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
