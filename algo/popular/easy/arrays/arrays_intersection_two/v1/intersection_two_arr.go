package v1

func intersect(nums1 []int, nums2 []int) []int {
	index := 0
	len1 := len(nums1)
	len2 := len(nums2)

	for i := 0; i < len1; i++ {
		for j := index; j < len2; j++ {

			// Здесь получается, что в начало первого слайса прописывают совпадающие значения из второго.
			if nums1[i] == nums2[j] {
				nums1[index] = nums1[i]
				nums2[index], nums2[j] = nums2[j], nums2[index]
				index++
				break
			}

		}
	}

	// На выход отдают только те значения из первого слайса, которые были прописаны.
	return nums1[:index]
}
