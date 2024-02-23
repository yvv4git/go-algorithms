package main

import (
	"reflect"
	"testing"
)

func Test_combinationSumV2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSumV2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
