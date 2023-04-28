package add_digits

import "testing"

func Test_addDigitsV2(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				num: 38,
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				num: 0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := addDigitsV2(tt.args.num); result != tt.want {
				t.Errorf("addDigitsV2() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_addDigitsV3_Research(t *testing.T) {
	for i := 0; i <= 20; i++ {
		remainder := i % 10
		t.Logf("i=%v remainder=%v", i, remainder)
	}
}
