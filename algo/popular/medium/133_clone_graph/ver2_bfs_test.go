package main

import (
	"reflect"
	"testing"
)

func TestCloneGraphV2(t *testing.T) {
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
			if got := CloneGraphV2(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloneGraphV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
