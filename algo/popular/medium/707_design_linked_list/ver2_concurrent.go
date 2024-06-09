package _07_design_linked_list

import "sync"

type LinkedList2 struct {
	Head *Node
	Size int
	mu   sync.Mutex
}

func Constructor2() LinkedList2 {
	return LinkedList2{&Node{}, 0, sync.Mutex{}}
}

func (this *LinkedList2) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}

	this.mu.Lock()
	defer this.mu.Unlock()

	curr := this.Head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *LinkedList2) AddAtHead(val int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	node := &Node{val, this.Head.Next}
	this.Head.Next = node
	this.Size++
}

func (this *LinkedList2) AddAtTail(val int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	node := &Node{val, nil}
	curr := this.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	this.Size++
}

func (this *LinkedList2) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}

	this.mu.Lock()
	defer this.mu.Unlock()

	node := &Node{val, nil}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	node.Next = curr.Next
	curr.Next = node
	this.Size++
}

func (this *LinkedList2) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	this.mu.Lock()
	defer this.mu.Unlock()

	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	this.Size--
}
