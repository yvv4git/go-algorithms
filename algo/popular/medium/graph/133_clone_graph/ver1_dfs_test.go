package main

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	type args struct {
		node *Node
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
