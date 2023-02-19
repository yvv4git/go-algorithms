package main

func majorityElement(nums []int) int {
	cnt := 0
	result := 0

	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			result = nums[i]
			cnt++
		} else if nums[i] == result {
			cnt++
		} else {
			cnt--
		}
	}

	return result
}

func majorityElementV2(nums []int) int {
	count := 0

	var candidate int
	for _, num := range nums {
		// Если на предыдущем шаге не нашли majority элемент, то проверим следующего кандидата.
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count += 1
		} else {
			count += -1
		}
	}

	return candidate
}
