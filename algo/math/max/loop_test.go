package max

import "testing"

func TestMaxByLoop(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxByLoop(tt.args.arr); got != tt.want {
				t.Errorf("MaxByLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFactorialByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxByLoop([]int{1, 2, 3, 4, 5})
	}
}
