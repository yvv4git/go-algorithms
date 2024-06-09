package main

import (
	"reflect"
	"testing"
)

func Test_deserialize(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want *NestedInteger
	}{
		{
			name: "Тест с простым целым числом",
			args: args{s: "324"},
			want: &NestedInteger{
				isInteger: true,
				value:     324,
				list:      nil,
			},
		},
		{
			name: "Тест с вложенным списком с одним уровнем вложенности",
			args: args{s: "[123,[456]]"},
			want: &NestedInteger{
				isInteger: false,
				value:     0,
				list: []*NestedInteger{
					{
						isInteger: true,
						value:     123,
						list:      nil,
					},
					{
						isInteger: false,
						value:     0,
						list: []*NestedInteger{
							{
								isInteger: true,
								value:     456,
								list:      nil,
							},
						},
					},
				},
			},
		},
		{
			name: "Тест с вложенным списком с несколькими уровнями вложенности",
			args: args{s: "[123,[456,[789]]]"},
			want: &NestedInteger{
				isInteger: false,
				value:     0,
				list: []*NestedInteger{
					{
						isInteger: true,
						value:     123,
						list:      nil,
					},
					{
						isInteger: false,
						value:     0,
						list: []*NestedInteger{
							{
								isInteger: true,
								value:     456,
								list:      nil,
							},
							{
								isInteger: false,
								value:     0,
								list: []*NestedInteger{
									{
										isInteger: true,
										value:     789,
										list:      nil,
									},
								},
							},
						},
					},
				},
			},
		},
		// Добавьте здесь дополнительные тесты, если необходимо
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
