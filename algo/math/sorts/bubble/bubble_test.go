package bubble

import (
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	type args struct {
		ar []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				ar: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				ar: []int{2, 1, 4, 3, 6, 5},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			Sort1(tt.args.ar)
		})
	}
}

func TestBubbleSort2(t *testing.T) {
	type args struct {
		ar []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				ar: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				ar: []int{2, 1, 4, 3, 6, 5},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			Sort2(tt.args.ar)
		})
	}
}

func TestBubbleSort3(t *testing.T) {
	type args struct {
		ar []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				ar: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				ar: []int{2, 1, 4, 3, 6, 5},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			Sort3(tt.args.ar)
		})
	}
}

func TestBubbleSort4(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{2, 1},
			},
			// want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int{2, 1, 4, 3, 6, 5},
			},
			// want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort4(tt.args.arr)
			t.Logf("RESULT: %v", tt.args.arr)
		})
	}
}

func TestBubbleSort5(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{2, 1},
			},
			// want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int{2, 1, 4, 3, 6, 5},
			},
			// want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort5(tt.args.arr)
			t.Logf("RESULT: %v", tt.args.arr)
		})
	}
}
