package main

import "testing"

func Test_isPathCrossing(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				path: "NES",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				path: "NESWW",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				path: "NESWW",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				path: "NESWW",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				path: "NESWW",
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				path: "NESWW",
			},
			want: true,
		},
		{
			name: "7",
			args: args{
				path: "NESWW",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPathCrossing(tt.args.path); got != tt.want {
				t.Errorf("isPathCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
