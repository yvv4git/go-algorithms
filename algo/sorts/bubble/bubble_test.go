package bubble

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
			BubbleSort(tt.args.ar)
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
			BubbleSort2(tt.args.ar)
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
			BubbleSort3(tt.args.ar)
		})
	}
}
