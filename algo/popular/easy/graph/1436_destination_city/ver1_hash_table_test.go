package _436_destination_city

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			},
			want: "Sao Paulo",
		},
		{
			name: "Test Case 2",
			args: args{
				paths: [][]string{{"Boston", "Texas"}, {"Texas", "Missouri"}, {"Missouri", "NewYork"}, {"NewYork", "Chicago"}},
			},
			want: "Chicago",
		},
		{
			name: "Test Case 3",
			args: args{
				paths: [][]string{{"A", "B"}, {"B", "C"}, {"C", "A"}},
			},
			want: "",
		},
		{
			name: "Test case 4",
			args: args{
				paths: [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			},
			want: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCityV1(tt.args.paths); got != tt.want {
				t.Errorf("destCityV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
