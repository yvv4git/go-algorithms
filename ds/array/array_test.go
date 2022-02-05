package array

import (
	"reflect"
	"testing"
)

func Test_arrayGet(t *testing.T) {
	type args struct {
		arr []int
		idx int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				idx: 1,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := arrayGet(tt.args.arr, tt.args.idx); result != tt.want {
				t.Errorf("array() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_arraySearch(t *testing.T) {
	type args struct {
		arr  []int
		want int
	}

	tests := []struct {
		name   string
		args   args
		expect bool
	}{
		{
			name: "CASE-1",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				want: 3,
			},
			expect: true,
		},
		{
			name: "CASE-2",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				want: 7,
			},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraySearch(tt.args.arr, tt.args.want); got != tt.expect {
				t.Errorf("arraySearch() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func Test_arrayDelete(t *testing.T) {
	type args struct {
		arr []int
		idx int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				idx: 1,
			},
			want: []int{1, 3, 4, 5},
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				idx: 0,
			},
			want: []int{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayDelete(tt.args.arr, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrayInsert(t *testing.T) {
	type args struct {
		arr        []int
		idx        int
		newElement int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				arr:        []int{1, 2, 3, 4, 5},
				idx:        1,
				newElement: 6,
			},
			want: []int{1, 6, 2, 3, 4, 5},
		},
		{
			name: "CASE-2",
			args: args{
				arr:        []int{1, 2, 3, 4, 5},
				idx:        0,
				newElement: 7,
			},
			want: []int{7, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayInsert(tt.args.arr, tt.args.idx, tt.args.newElement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
