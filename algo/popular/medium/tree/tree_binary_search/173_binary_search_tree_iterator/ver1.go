package main

// BSTIterator - определяем структуру BSTIterator, которая будет хранить стек узлов.
type BSTIterator struct {
	stack []*TreeNode // Стек узлов
}

// Constructor - конструктор BSTIterator, который принимает корень BST и возвращает новый экземпляр BSTIterator.
// В конструкторе мы инициализируем стек и добавляем корень в стек.
func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{stack: make([]*TreeNode, 0)}
	iter.push(root)
	return iter
}

// Next - возвращает следующий наименьший элемент в BST.
// Метод pop извлекает верхний узел из стека, добавляет правого потомка этого узла в стек,
// и возвращает значение этого узла.
// TIME COMPLEXITY: O(1), в среднем, так как мы просто извлекаем верхний узел из стека.
// SPACE COMPLEXITY: O(1), так как мы не используем дополнительное пространство, кроме стека.
func (this *BSTIterator) Next() int {
	node := this.pop()
	this.push(node.Right)
	return node.Val
}

// HasNext - проверяет, существует ли следующий наименьший элемент.
// Метод просто проверяет, не пуст ли стек.
// TIME COMPLEXITY: O(1), так как мы просто проверяем, пуст ли стек.
// SPACE COMPLEXITY: O(1), так как мы не используем дополнительное пространство, кроме стека.
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// Push - добавляет узел в стек.
// Метод проходит от корня к самому левому листу, добавляя все пройденные узлы в стек.
// TIME COMPLEXITY: O(n), где n - количество узлов в левом поддереве. В худшем случае, когда дерево сбалансировано, мы добавляем все узлы в стек.
// SPACE COMPLEXITY: O(n), так как в худшем случае мы храним все узлы в стеке.
func (this *BSTIterator) push(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

// Pop - извлекает верхний узел из стека.
// Метод просто удаляет верхний узел из стека и возвращает его.
// TIME COMPLEXITY: O(1), так как мы просто удаляем верхний узел из стека.
// SPACE COMPLEXITY: O(1), так как мы не используем дополнительное пространство, кроме стека.
func (this *BSTIterator) pop() *TreeNode {
	n := len(this.stack)
	node := this.stack[n-1]
	this.stack = this.stack[:n-1]
	return node
}
