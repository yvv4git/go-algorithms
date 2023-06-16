package binary_watch

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readBinaryWatchV2(t *testing.T) {
	type args struct {
		turnedOn int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "CASE-1",
			args: args{
				turnedOn: 1,
			},
			want: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			name: "CASE-2",
			args: args{
				turnedOn: 9,
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatchV2(tt.args.turnedOn); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("%#v \n", got)
				t.Errorf("readBinaryWatchV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
