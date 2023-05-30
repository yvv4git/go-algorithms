package ex01

import (
	"testing"
)

// Количество перестановок некоторого множества равно факториалу от количества элементов этого множества.
func TestTransposeCnt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "case-1",
			args: args{
				n: 1,
			},
			want: 1,
			desc: "1! = 1",
		},
		{
			name: "case-2",
			args: args{
				n: 2,
			},
			want: 2,
			desc: "2! = 1*2 = 2",
		},
		{
			name: "case-3",
			args: args{
				n: 3,
			},
			want: 6,
			desc: "3! = 1*2*3 = 2*3 = 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransposeCnt(tt.args.n); got != tt.want {
				t.Errorf("TransposeCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerm(t *testing.T) {
	type args struct {
		a []rune
		f func([]rune)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case-1",
			args: args{
				a: []rune{'a', 'b', 'c'},
				f: func(a []rune) {
					t.Log(string(a))
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Perm(tt.args.a, tt.args.f)
		})
	}
}
