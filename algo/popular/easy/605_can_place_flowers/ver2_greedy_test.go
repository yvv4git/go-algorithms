package main

import "testing"

func Test_canPlaceFlowersV2(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// Тестовые кейсы
		{
			name: "Посадка одного цветка в пустую клумбу",
			args: args{
				flowerbed: []int{0, 0, 0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Посадка одного цветка в клумбу с одним цветком",
			args: args{
				flowerbed: []int{1, 0, 0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Посадка одного цветка в клумбу с двумя цветами",
			args: args{
				flowerbed: []int{1, 0, 1},
				n:         1,
			},
			want: false,
		},
		{
			name: "Посадка трех цветов в клумбу с двумя цветами",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         3,
			},
			want: false,
		},
		{
			name: "Посадка двух цветов в клумбу с двумя цветами",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "Посадка одного цветка в клумбу с двумя цветами",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Посадка нуля цветов в пустую клумбу",
			args: args{
				flowerbed: []int{0, 0, 0},
				n:         0,
			},
			want: true,
		},
		{
			name: "Посадка одного цветка в клумбу с одним цветком в конце",
			args: args{
				flowerbed: []int{0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Посадка одного цветка в клумбу с одним цветком в начале",
			args: args{
				flowerbed: []int{1, 0, 0},
				n:         1,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowersV2(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowersV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
