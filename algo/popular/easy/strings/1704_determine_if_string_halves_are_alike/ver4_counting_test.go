package main

import "testing"

func Test_halvesAreAlikeV4(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{s: "book"},
			want: true,
		},
		{
			name: "Test 2",
			args: args{s: "textbook"},
			want: false,
		},
		{
			name: "Test 3",
			args: args{s: "MerryChristmas"},
			want: false,
		},
		{
			name: "Test 4",
			args: args{s: "AbCdEfGh"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halvesAreAlikeV4(tt.args.s); got != tt.want {
				t.Errorf("halvesAreAlikeV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
