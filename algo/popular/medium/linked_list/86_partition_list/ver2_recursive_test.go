package main

import (
	"reflect"
	"testing"
)

func Test_partitionRecursive(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionRecursive(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
