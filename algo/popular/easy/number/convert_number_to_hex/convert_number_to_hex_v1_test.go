package convert_number_to_hex

import "testing"

func Test_toHexV1(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				num: 26,
			},
			want: "1a",
		},
		{
			name: "CASE-2",
			args: args{
				num: -1,
			},
			want: "ffffffff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHexV1(tt.args.num); got != tt.want {
				t.Errorf("toHexV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
