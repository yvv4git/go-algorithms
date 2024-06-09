package main

import (
	"reflect"
	"testing"
)

func Test_majorityElementV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementV2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElementV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
