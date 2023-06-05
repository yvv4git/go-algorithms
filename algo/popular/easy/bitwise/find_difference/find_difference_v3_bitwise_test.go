package find_difference

import "testing"

func Test_findTheDifferenceV3(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "CASE-1",
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: 'e', // 101
		},
		{
			name: "CASE-2",
			args: args{
				s: "",
				t: "y",
			},
			want: 'y', // 121
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifferenceV3(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifferenceV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchV3Ex01(t *testing.T) {
	/*
		Таким образом можно находить уникальный символ, цифру и т.д,
	*/
	nums := []int{1, 2, 3, 1, 2, 3, 4}
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	t.Logf("diff: %#v", diff) // 4
}
