package main

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				name:  "alex",
				typed: "aaleex",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				name:  "saeed",
				typed: "ssaaedd",
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				name:  "leelee",
				typed: "lleeelee",
			},
			want: true,
		},
		{
			name: "CASE-4",
			args: args{
				name:  "laiden",
				typed: "laiden",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
