package _464_max_product_of_two_elem_of_arr

func maxProductV2(nums []int) int {
	max := 0
	afterMax := 0

	for _, num := range nums {
		if num > max {
			afterMax = max
			max = num
		} else {
			if num > afterMax {
				afterMax = num
			}
		}
	}

	return (max - 1) * (afterMax - 1)
}
