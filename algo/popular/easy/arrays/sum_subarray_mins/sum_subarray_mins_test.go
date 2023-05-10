package sum_subarray_mins

import "testing"

func Test_sumSubarrayMins(t *testing.T) {
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
				arr: []int{3, 1, 2, 4},
			},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(tt.args.arr); got != tt.want {
				t.Errorf("sumSubarrayMins() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------------------------------------------------------

func TestResearchSubSubarrayMins01(t *testing.T) {
	// Символ % позволяет получить остаток от деления.
	t.Logf("17 %% 1000000007 = %v", 17%1000000007) // 17

	t.Logf("17 %% 1 = %v", 17%1) // 0
	t.Logf("17 %% 2 = %v", 17%2) // 1
}
