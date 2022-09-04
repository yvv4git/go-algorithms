package concurrent

func Merge(leftPart []int, rightPart []int) (result []int) {
	result = make([]int, len(leftPart)+len(rightPart))
	idxLeft, idxRight := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case idxLeft >= len(leftPart):
			result[i] = rightPart[idxRight]
			idxRight++
		case idxRight >= len(rightPart):
			result[i] = leftPart[idxLeft]
			idxLeft++
		case leftPart[idxLeft] < rightPart[idxRight]:
			result[i] = leftPart[idxLeft]
			idxLeft++
		default:
			result[i] = rightPart[idxRight]
			idxRight++
		}
	}

	return
}

func MergeSort(data []int, r chan []int) {
	if len(data) == 1 {
		r <- data
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(data) / 2

	go MergeSort(data[:middle], leftChan)
	go MergeSort(data[middle:], rightChan)

	leftPart := <-leftChan
	rightPart := <-rightChan

	close(leftChan)
	close(rightChan)

	r <- Merge(leftPart, rightPart)
	return
}
