package main

import (
	"reflect"
	"testing"
)

func Test_generateMatrixV2(t *testing.T) {
	type args struct {
		n int
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
			if got := generateMatrixV2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrixV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
