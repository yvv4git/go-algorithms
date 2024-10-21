package main

import "testing"

func Test_distributeCandiesV4(t *testing.T) {
	type args struct {
		candyType []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{candyType: []int{1, 1, 2, 2, 3, 3}},
			want: 3,
		},
		{
			name: "example 2",
			args: args{candyType: []int{1, 1, 2, 3}},
			want: 2,
		},
		{
			name: "all same",
			args: args{candyType: []int{1, 1, 1, 1}},
			want: 1,
		},
		{
			name: "all different",
			args: args{candyType: []int{1, 2, 3, 4}},
			want: 2,
		},
		{
			name: "empty",
			args: args{candyType: []int{}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandiesV4(tt.args.candyType); got != tt.want {
				t.Errorf("distributeCandiesV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
