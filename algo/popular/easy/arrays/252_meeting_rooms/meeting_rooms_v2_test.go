package _252_meeting_rooms

import "testing"

func Test_canAttendMeetingsV2(t *testing.T) {
	type args struct {
		intervals []Interval
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				intervals: []Interval{{0, 30}, {5, 10}, {15, 20}},
			},
			want: false,
		},
		{
			name: "CASE-2",
			args: args{
				intervals: []Interval{{7, 10}, {2, 4}},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetingsV2(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendMeetingsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
