package main

import "testing"

func Test_titleToNumberV2(t *testing.T) {
	type args struct {
		columnTitle string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Single letter column",
			args: args{columnTitle: "A"},
			want: 1,
		},
		{
			name: "Single letter column",
			args: args{columnTitle: "Z"},
			want: 26,
		},
		{
			name: "Double letter column",
			args: args{columnTitle: "AA"},
			want: 27,
		},
		{
			name: "Double letter column",
			args: args{columnTitle: "AB"},
			want: 28,
		},
		{
			name: "Triple letter column",
			args: args{columnTitle: "AAA"},
			want: 703,
		},
		{
			name: "Triple letter column",
			args: args{columnTitle: "ABC"},
			want: 731,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumberV2(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
