package _07_design_linked_list

// Node represents a node of linked list
type Node struct {
	Val  int
	Next *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *Node
	Size int
}

func Constructor() LinkedList {
	return LinkedList{&Node{}, 0}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
// TIME COMPLEXITY: O(n)
func (this *LinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	curr := this.Head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

// AddAtHead - добавить ноду с головы
// TIME COMPLEXITY: O(1)
func (this *LinkedList) AddAtHead(val int) {
	node := &Node{val, this.Head.Next}
	this.Head.Next = node
	this.Size++
}

// AddAtTail - добавить ноду с хвоста
// TIME COMPLEXITY: O(n)
func (this *LinkedList) AddAtTail(val int) {
	node := &Node{val, nil}
	curr := this.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	this.Size++
}

// AddAtIndex - добавить по индексу
// TIME COMPLEXITY: O(n)
func (this *LinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	node := &Node{val, nil}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	node.Next = curr.Next
	curr.Next = node
	this.Size++
}

// DeleteAtIndex - удалить по индексу
// TIME COMPLEXITY: O(n)
func (this *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	this.Size--
}
