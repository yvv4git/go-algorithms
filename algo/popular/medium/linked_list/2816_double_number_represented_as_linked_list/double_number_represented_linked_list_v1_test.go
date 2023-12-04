package _816_double_number_represented_as_linked_list

import (
	"reflect"
	"testing"
)

func Test_doubleItV1(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test case 1",
			args: args{
				head: createLinkedList([]int{1, 8, 9}),
			},
			want: createLinkedList([]int{3, 7, 8}),
		},
		{
			name: "Test case 2",
			args: args{
				head: createLinkedList([]int{5, 6, 7}),
			},
			want: createLinkedList([]int{1, 1, 1, 1}),
		},
		{
			name: "Test case 3",
			args: args{
				head: createLinkedList([]int{0, 1, 2, 3, 4}),
			},
			want: createLinkedList([]int{0, 2, 4, 6, 8}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleItV1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Logf("got=%#v want=%#v", *got, *tt.want)
				t.Errorf("doubleItV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
