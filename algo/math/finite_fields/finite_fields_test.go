package main

import (
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		x     int
		y     int
		field int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				x:     7,
				y:     8,
				field: 19,
			},
			want: 15,
		},
		{
			name: "CASE-2",
			args: args{
				x:     11,
				y:     17,
				field: 19,
			},
			want: 9,
		},
		{
			name: "CASE-3",
			args: args{
				x:     5,
				y:     4,
				field: 19,
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.x, tt.args.y, tt.args.field); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inversion(t *testing.T) {
	type args struct {
		x     int
		field int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				x:     9,
				field: 19,
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inversion(tt.args.x, tt.args.field); got != tt.want {
				t.Errorf("inversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
