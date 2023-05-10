package sort_colors

func sortColorsV5(nums []int) {
	/*
		Method: a+b+c = n
		Time: O(n) + O(n) = O(2n) = O(n)
		Space: O(1)
	*/
	// Time: O(n)
	var nbZeroes, nbOnes, nbTwos int = 0, 0, 0
	for _, num := range nums {
		if num == 0 {
			nbZeroes++
		} else if num == 1 {
			nbOnes++
		} else /* if num == 2 */ {
			nbTwos++
		}
	}

	// Time: O(n)
	for idx, _ := range nums {
		if idx >= nbZeroes+nbOnes {
			nums[idx] = 2
		} else if idx >= nbZeroes {
			nums[idx] = 1
		} else /* idx > 0 */ {
			nums[idx] = 0
		}
	}

}
