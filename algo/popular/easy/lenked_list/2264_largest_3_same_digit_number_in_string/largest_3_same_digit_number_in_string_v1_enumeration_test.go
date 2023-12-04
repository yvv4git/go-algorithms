package _264_largest_3_same_digit_number_in_string

import "testing"

func Test_largestGoodIntegerV1(t *testing.T) {
	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Тест 1",
			args: args{num: "6777133339"},
			want: "777",
		},
		{
			name: "Тест 2",
			args: args{num: "2300019"},
			want: "000",
		},
		{
			name: "Тест 3",
			args: args{num: "42352338"},
			want: "",
		},
		{
			name: "Тест 4",
			args: args{num: "111222333444555"},
			want: "555",
		},
		{
			name: "Тест 5",
			args: args{num: "999888777666555"},
			want: "999",
		},
		{
			name: "Тест 6",
			args: args{num: "1234567890"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodIntegerV1(tt.args.num); got != tt.want {
				t.Errorf("largestGoodIntegerV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
