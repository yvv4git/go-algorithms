package main

import (
	"reflect"
	"testing"
)

func Test_constructRectangleV2(t *testing.T) {
	type args struct {
		area int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1: Area is a prime number",
			args: args{area: 37},
			want: []int{37, 1},
		},
		{
			name: "Case 2: Area is a perfect square",
			args: args{area: 36},
			want: []int{6, 6},
		},
		{
			name: "Case 3: Area is a composite number",
			args: args{area: 48},
			want: []int{8, 6},
		},
		{
			name: "Case 4: Area is 1",
			args: args{area: 1},
			want: []int{1, 1},
		},
		{
			name: "Case 5: Area is a large number",
			args: args{area: 1000000},
			want: []int{1000, 1000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructRectangleV2(tt.args.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangleV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
