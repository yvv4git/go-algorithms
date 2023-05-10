package hard

import "testing"

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.

## Constraints:
- k == lists.length
- 0 <= k <= 104
- 0 <= lists[i].length <= 500
- -104 <= lists[i][j] <= 104
- lists[i] is sorted in ascending order.
- The sum of lists[i].length will not exceed 104.

## Example-1
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6K
]
merging them into one sorted list:
1->1->2->3->4->4->5->6


## Example-2
Input: lists = []
Output: []

## Example-3
Input: lists = [[]]
Output: []
*/

func TestMergeKSortedArrays(t *testing.T) {
	t.Log("Start")

	type args struct {
		lists []*ListNode
	}

	/*tests := []struct{
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				lists: []*ListNode{
					&ListNode{
						Val: 1,
						Next:
					},
				},
			},
		},
	}*/

	t.Log("End")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := NewHeap(3500)
	c := 0
	for _, v := range lists {
		for v != nil {
			c++
			heap.Add(v.Val)
			v = v.Next
		}
	}
	res := &ListNode{}
	head := res
	for c > 0 {
		head.Next = &ListNode{Val: heap.Get()}
		head = head.Next
		c--
	}
	return res.Next
}

type BinHeap struct {
	h []int
}

func NewHeap(size int) *BinHeap {
	return &BinHeap{make([]int, 0, size)}
}

func (b *BinHeap) siftUp(num int) {
	for num > 0 && b.h[num] < b.h[(num-1)/2] {
		b.h[num], b.h[(num-1)/2] = b.h[(num-1)/2], b.h[num]
		num = (num - 1) / 2
	}
}

func (b *BinHeap) Add(val int) {
	b.h = append(b.h, val)
	b.siftUp(len(b.h) - 1)
}

func (b *BinHeap) siftDown() {
	b.h[0] = b.h[len(b.h)-1]
	b.h = b.h[:len(b.h)-1]
	for i := 0; 2*i+1 < len(b.h); {
		j := 2*i + 1
		if 2*i+2 < len(b.h) && b.h[2*i+2] < b.h[j] {
			j = 2*i + 2
		}
		if b.h[i] < b.h[j] {
			break
		} else {
			b.h[j], b.h[i] = b.h[i], b.h[j]
		}
		i = j
	}

}
func (b *BinHeap) Get() int {
	x := b.h[0]
	b.siftDown()
	return x
}
