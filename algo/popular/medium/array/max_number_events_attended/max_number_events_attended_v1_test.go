package max_number_events_attended

import "testing"

func Test_maxEventsV1(t *testing.T) {
	type args struct {
		events [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				events: [][]int{{1, 2}, {2, 3}, {3, 4}},
			},
			want: 3,
		},
		{
			name: "CASE-2",
			args: args{
				events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEventsV1(tt.args.events); got != tt.want {
				t.Errorf("maxEventsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
