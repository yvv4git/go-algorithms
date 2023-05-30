package ex01_count

import (
	"testing"
)

func TestCombinationsCnt(t *testing.T) {
	type args struct {
		n int
		m int
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
				n: 3,
				m: 2,
			},
			want: 3,
		},
		{
			name: "case-2",
			args: args{
				n: 3,
				m: 3,
			},
			want: 1,
			desc: "Обрати внимание, когда n = m, количество сочетаний будет равно 1",
		},
		{
			name: "case-3",
			args: args{
				n: 4,
				m: 2,
			},
			want: 6,
			desc: "Допустим, на вечеринке 4 человека. Из них можно образовать 6 разных пар",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationsCnt(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("CombinationsCnt() = %v, want %v \n", got, tt.want)
			}
		})
	}
}

func TestTransposeSliceRunes(t *testing.T) {
	type args struct {
		a     []rune
		width int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case-1",
			args: args{
				a:     []rune("abc"),
				width: 1,
			},
		},
		{
			name: "case-2",
			args: args{
				a:     []rune("abc"),
				width: 2,
			},
		},
		{
			name: "case-3",
			args: args{
				a:     []rune("abc"),
				width: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TransposeSliceRunes(tt.args.a, tt.args.width)
		})
	}
}
