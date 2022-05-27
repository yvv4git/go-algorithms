package main

import "testing"

func Test_clockAfter(t *testing.T) {
	type args struct {
		nowTime     int
		afterHourse int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				nowTime:     3,
				afterHourse: 47,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clockAfter(tt.args.nowTime, tt.args.afterHourse); got != tt.want {
				t.Errorf("clockAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}
