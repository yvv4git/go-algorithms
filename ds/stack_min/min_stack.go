package stackmin

type MinStack struct {
	head *StackNode
}

func Constructor() MinStack {
	return MinStack{
		head: nil,
	}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = NewStackNode(val, val, nil)
	} else {
		min := val
		if this.head.min < min { // ключевой момент, когда определяем минимальную ноду
			min = this.head.min // Каждая нода хранит минимальное значение, которое было до нее вместе с ней
		}
		this.head = NewStackNode(
			val,       // текущее значение
			min,       // минимальное значение
			this.head, // this.head - верхушка стека, теперь на нее будет указывать новая верхушка стека
		)
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

// StackNode - used as implementations of stack node
type StackNode struct {
	val  int
	min  int
	next *StackNode
}

// NewStackNode - simple factory for create instance of StackNode
func NewStackNode(val, min int, next *StackNode) *StackNode {
	return &StackNode{
		val:  val,
		min:  min,
		next: next,
	}
}
