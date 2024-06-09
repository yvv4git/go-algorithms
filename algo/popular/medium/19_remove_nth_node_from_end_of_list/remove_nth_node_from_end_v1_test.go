package main

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEndV1(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
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
			if got := removeNthFromEndV1(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEndV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
