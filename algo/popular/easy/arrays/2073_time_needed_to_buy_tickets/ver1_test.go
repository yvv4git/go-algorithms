package _073_time_needed_to_buy_tickets

import "testing"

func Test_timeRequiredToBuyV1(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "Test Case 1",
			args: args{
				tickets: []int{2, 3, 2},
				k:       2,
			},
			want: 6,
			desc: `
			На первом проходе каждый покупатель купит по билету и в массиве будет: [1, 2, 1], res = 3
  			На втором проходе каждый покупатель купит по билету и в массиве будет: [0, 1, 0], покупатель tickets[2] = 0, значит k-ый покупатель все купил, res=3
			res = 3 + 3 = 6
			`,
		},
		{
			name: "Test Case 2",
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       0,
			},
			want: 8,
			desc: `
			После первого прохода массив станет таким: [4, 0, 0, 0], res = 4, так как каждый покупатель(их 4) купил 1 билет.
			Далее покупатель tickets[0] будет покупать билеты 4 раза, т.е. это займет 4 прохода.
			`,
		},
		{
			name: "Test Case 3",
			args: args{
				tickets: []int{1, 1, 1, 1},
				k:       0,
			},
			want: 1,
			desc: `
				На первом проходе tickets[k=0] покупает 1 билет и на этом задача выполнена. 
				В итоге, res = 0
			`,
		},
		{
			name: "Test Case 4",
			args: args{
				tickets: []int{1, 2, 3, 4},
				k:       3,
			},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuyV1(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuyV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
