package binary_watch

import (
	"reflect"
	"testing"
)

func Test_readBinaryWatchV1(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "CASE-1",
			args: args{
				num: 1,
			},
			want: []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"},
		},
		{
			name: "CASE-2",
			args: args{
				num: 9,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatchV1(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatchV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
