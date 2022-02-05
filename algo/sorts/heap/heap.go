package heap

func HeapSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	heapIt(ar)

	ar[0], ar[len(ar)-1] = ar[len(ar)-1], ar[0]

	HeapSort(ar[:len(ar)-1])
}

func heapIt(ar []int) {
	// left = 2*i + 1
	// right = 2*i + 2
	// i = 0
	// мы каждый раз рассматриваем 3 ноды - рут и 2 листа
	// и по кажлому листу повторяем рекурсивно heapIt(ar[1:]), heapIt(ar[2:])
	if len(ar) < 2 {
		return
	}

	if len(ar) == 2 {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
		return
	}

	if len(ar) > 3 {
		heapIt(ar[1:])
		heapIt(ar[2:])
	}

	if ar[1] > ar[2] {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
	} else {
		if ar[0] < ar[2] {
			ar[0], ar[2] = ar[2], ar[0]
		}
	}
}
