package sort_colors

func sortColorsV4(nums []int) {
	/*
		METHOD: a+b+c = n
		Time: O(n) + O(a+b+c) = O(n) + O(n) = O(2n) = O(n)
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

	// Time: O(a) where `a` is the number of "0"
	for i := 0; i < nbZeroes; i++ {
		nums[i] = 0
	}

	// Time: O(b) where `b` is the number of "1"
	for i := nbZeroes; i < nbZeroes+nbOnes; i++ {
		nums[i] = 1
	}

	// Time: O(c) where `c` is the number of "2"
	for i := nbZeroes + nbOnes; i < nbZeroes+nbOnes+nbTwos; i++ {
		nums[i] = 2
	}
}
