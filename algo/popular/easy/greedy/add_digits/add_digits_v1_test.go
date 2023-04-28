package add_digits

import "testing"

func Test_addDigitsV1(t *testing.T) {
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
			if result := addDigitsV1(tt.args.num); result != tt.want {
				t.Errorf("addDigitsV1() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_addDigits_Research(t *testing.T) {
	for i := 10; i <= 20; i++ {
		r := i % 9 // Операция получения остатка от деления
		t.Logf("Чтобы разложить i=%d на сумму чисел, надо получить остаток от деления на 9 r=%v  ", i, r)
	}
}

func Test_addDigits_Research2(t *testing.T) {
	for i := 0; i <= 20; i++ {
		r := i % 9 // Операция получения остатка от деления
		t.Logf("Чтобы разложить i=%d на сумму чисел, надо получить остаток от деления на 9 r=%v  ", i, r)
	}
}
