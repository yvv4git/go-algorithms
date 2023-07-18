package range_sub_query

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		left  int
		right int
	}

	tests := []struct {
		arr  []int
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			arr:  []int{-2, 0, 3, -5, 2, -1},
			args: args{
				left:  0,
				right: 2,
			},
			want: 1,
		},
		{
			name: "CASE-2",
			arr:  []int{-2, 0, 3, -5, 2, -1},
			args: args{
				left:  2,
				right: 5,
			},
			want: -1,
		},
		{
			name: "CASE-3",
			arr:  []int{-2, 0, 3, -5, 2, -1},
			args: args{
				left:  0,
				right: 5,
			},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//this := &NumArray{
			//	data: tt.fields.data,
			//}
			//if got := this.SumRange(tt.args.left, tt.args.right); got != tt.want {
			//	t.Errorf("SumRange() = %v, want %v", got, tt.want)
			//}
		})
	}
}
