package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// MerkleTree представляет структуру дерева Меркла.
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode представляет узел дерева Меркла.
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree создает новое дерево Меркла из списка данных.
// Временная сложность: O(n log n), где n — количество элементов в данных.
// Пространственная сложность: O(n), где n — количество элементов в данных.
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	// Если количество данных нечетное, дублируем последний элемент.
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	// Создаем листья дерева.
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	// Строим дерево Меркла, комбинируя узлы.
	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}

	// Корневой узел дерева Меркла.
	mTree := MerkleTree{&nodes[0]}

	return &mTree
}

// NewMerkleNode создает новый узел дерева Меркла.
// Временная сложность: O(1).
// Пространственная сложность: O(1).
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}

	if left == nil && right == nil {
		// Если это лист, хешируем данные.
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		// Если это не лист, комбинируем хеши от дочерних узлов.
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}

// String возвращает строковое представление хеша узла.
// Временная сложность: O(1).
// Пространственная сложность: O(1).
func (n *MerkleNode) String() string {
	return hex.EncodeToString(n.Data)
}

// AddNode добавляет новый узел в дерево Меркла.
// Временная сложность: O(n log n), где n — количество элементов в дереве.
// Пространственная сложность: O(n).
func (mt *MerkleTree) AddNode(data []byte) {
	newNode := NewMerkleNode(nil, nil, data)
	mt.RootNode = mt.rebuildTree(mt.RootNode, newNode)
}

// rebuildTree перестраивает дерево Меркла после добавления нового узла.
// Временная сложность: O(n log n), где n — количество элементов в дереве.
// Пространственная сложность: O(n).
func (mt *MerkleTree) rebuildTree(currentNode, newNode *MerkleNode) *MerkleNode {
	if currentNode == nil {
		return newNode
	}

	// Создаем новый узел, комбинируя текущий узел и новый узел.
	newRoot := NewMerkleNode(currentNode, newNode, nil)

	return newRoot
}

// RemoveNode удаляет узел из дерева Меркла.
// Временная сложность: O(n log n), где n — количество элементов в дереве.
// Пространственная сложность: O(n).
func (mt *MerkleTree) RemoveNode(data []byte) {
	mt.RootNode = mt.removeNode(mt.RootNode, data)
}

// removeNode рекурсивно удаляет узел из дерева Меркла.
// Временная сложность: O(n log n), где n — количество элементов в дереве.
// Пространственная сложность: O(n).
func (mt *MerkleTree) removeNode(currentNode *MerkleNode, data []byte) *MerkleNode {
	if currentNode == nil {
		return nil
	}

	// Если текущий узел — лист и его данные совпадают с удаляемыми данными.
	if currentNode.Left == nil && currentNode.Right == nil && string(currentNode.Data) == string(data) {
		return nil
	}

	// Рекурсивно удаляем узел из левого поддерева.
	currentNode.Left = mt.removeNode(currentNode.Left, data)

	// Рекурсивно удаляем узел из правого поддерева.
	currentNode.Right = mt.removeNode(currentNode.Right, data)

	// Если оба дочерних узла удалены, удаляем текущий узел.
	if currentNode.Left == nil && currentNode.Right == nil {
		return nil
	}

	// Пересчитываем хеш текущего узла.
	currentNode.Data = mt.calculateHash(currentNode.Left, currentNode.Right)

	return currentNode
}

// calculateHash вычисляет хеш для узла на основе его дочерних узлов.
// Временная сложность: O(1).
// Пространственная сложность: O(1).
func (mt *MerkleTree) calculateHash(left, right *MerkleNode) []byte {
	var prevHashes []byte
	if left != nil {
		prevHashes = append(prevHashes, left.Data...)
	}
	if right != nil {
		prevHashes = append(prevHashes, right.Data...)
	}
	hash := sha256.Sum256(prevHashes)
	return hash[:]
}

func main() {
	// Пример данных для дерева Меркла.
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
		[]byte("data4"),
	}

	// Создаем дерево Меркла.
	mTree := NewMerkleTree(data)

	// Выводим корень дерева Меркла.
	fmt.Printf("Merkle Root: %s\n", mTree.RootNode.String())

	// Добавляем новый узел.
	mTree.AddNode([]byte("data5"))
	fmt.Printf("Merkle Root after adding 'data5': %s\n", mTree.RootNode.String())

	// Удаляем узел.
	mTree.RemoveNode([]byte("data3"))
	fmt.Printf("Merkle Root after removing 'data3': %s\n", mTree.RootNode.String()) // Падает с паникой
}
