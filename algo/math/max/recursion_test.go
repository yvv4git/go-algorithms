package max

import "testing"

func TestMaxByRecursion(t *testing.T) {
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
			//t.Log(MaxByRecursion(tt.args.arr))
			if got := MaxByRecursion(tt.args.arr); got != tt.want {
				t.Errorf("MaxByRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFactorialByRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxByRecursion([]int{1, 2, 3, 4, 5})
	}
}
