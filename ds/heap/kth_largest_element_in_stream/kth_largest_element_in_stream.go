// Package kth_largest_element_in_stream
/*
703_kth_largest_element_in_a_stream

В общем, если на вход дается массив чисел и надо отобрать k-ый наибольший элемент, то:
1. Метод с heap.
Помещаем элементы в кучу и потом достаем k элементов.
Соответственно, k-ый элемент и будет ответом.
*/
package kth_largest_element_in_stream

import "container/heap"

type KthLargest struct {
	k    int
	heap intHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	//Convert to heap
	heap.Init(&h)

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)

	//Pop minimum element until len(h) = k and kthlargest = h[0]
	for len(kl.heap) > kl.k {
		heap.Pop(&kl.heap)
	}

	return kl.heap[0]
}
