package range_sub_query

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{data: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0

	for _, value := range this.data[left : right+1] {
		sum += value
	}

	return sum
}
