package main

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "CASE-1",
			args: args{
				num: 5,
			},
			want: 2684354560,
		},
		{
			name: "CASE-2",
			args: args{
				num: 43261596,
			},
			want: 964176192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
