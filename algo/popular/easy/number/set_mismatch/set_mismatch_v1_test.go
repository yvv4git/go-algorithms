package set_mismatch

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findErrorNumsV1(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{1, 2, 2, 4},
			},
			want: []int{2, 3},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{3, 2, 2},
			},
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNumsV1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNumsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchV1Ex01(t *testing.T) {
	nums := []int{3, 2, 2}
	i := 0

	fmt.Printf("nums: %#v \n", nums)
	for i < len(nums) {
		fmt.Printf("\ti=%d \n", i)
		desiredPosition := nums[i] - 1
		fmt.Printf("\tdesiredPosition=%d \n", i)

		if nums[i]-1 != i && nums[i] != nums[desiredPosition] {
			nums[i], nums[desiredPosition] = nums[desiredPosition], nums[i]
		} else {
			i++
		}
	}
	fmt.Printf("nums: %#v \n", nums)

	var res []int

	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			res = append(res, nums[j])
			res = append(res, j+1)
			break
		}
	}

	t.Logf("Result: %#v", res)
}

func TestResearchV1Ex02(t *testing.T) {
	/*
		Непонятный алгоритм сортировки:
		- это in-place алгоритм
		- он работает только с числами, которые идут по очереди: 1,2,3,4,5. Не должно быть пропусков
		- time complexity = O(n)
	*/
	nums := []int{5, 4, 3, 2, 1}
	i := 0

	fmt.Printf("nums: %#v \n", nums)
	for i < len(nums) {
		desiredPosition := nums[i] - 1
		fmt.Printf("\ti=%d desiredPosition=%d \n", i, desiredPosition)

		if nums[i]-1 != i && nums[i] != nums[desiredPosition] {
			nums[i], nums[desiredPosition] = nums[desiredPosition], nums[i]
		} else {
			i++
		}
	}
	fmt.Printf("nums: %#v \n", nums)
}

func TestResearchV1Ex03(t *testing.T) {
	nums := []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
	i := 0

	fmt.Printf("nums: %#v \n", nums)
	for i < len(nums) {
		desiredPosition := nums[i] - 1
		fmt.Printf("\ti=%d desiredPosition=%d \n", i, desiredPosition)

		if nums[i]-1 != i && nums[i] != nums[desiredPosition] {
			nums[i], nums[desiredPosition] = nums[desiredPosition], nums[i]
		} else {
			i++
		}
	}
	fmt.Printf("nums: %#v \n", nums)
}
