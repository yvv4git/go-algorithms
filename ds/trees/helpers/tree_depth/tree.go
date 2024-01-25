package tree

import (
	"container/list"
	"fmt"
)

type Node struct {
	ID       int64
	Name     string
	ParentID int64
}

type Tree struct {
	nodes map[int64]*Node
	root  *Node
}

func New(nodes []Node) *Tree {
	nodeMap := make(map[int64]*Node)
	var root *Node

	// First pass: create nodes and map them by their ID.
	for _, node := range nodes {
		newNode := &Node{
			ID:       node.ID,
			Name:     node.Name,
			ParentID: node.ParentID,
		}
		nodeMap[node.ID] = newNode

		// If the node has no parent, it's the root.
		if node.ParentID == 0 {
			root = newNode
		}
	}

	return &Tree{
		nodes: nodeMap,
		root:  root,
	}
}

// CalculateHeight calculates the height of the tree using iterative DFS.
func (t *Tree) CalculateHeight() int {
	if t.root == nil {
		return 0
	}

	type stackItem struct {
		node  *Node
		depth int
	}

	stack := list.New()
	stack.PushBack(stackItem{node: t.root, depth: 1})
	maxDepth := 0

	for stack.Len() > 0 {
		element := stack.Back()
		item := element.Value.(stackItem)
		stack.Remove(element)

		if item.depth > maxDepth {
			maxDepth = item.depth
		}

		// Add all children of the current node to the stack.
		for _, childNode := range t.nodes {
			if childNode.ParentID == item.node.ID {
				stack.PushBack(stackItem{node: childNode, depth: item.depth + 1})
			}
		}
	}

	return maxDepth
}

// Print prints the tree structure to the console with dashes indicating hierarchy.
func (t *Tree) Print() {
	if t.root == nil {
		fmt.Println("The tree is empty.")
		return
	}

	// Helper function to perform DFS and print the tree with dashes.
	var printTree func(node *Node, indent string, isLast bool)
	printTree = func(node *Node, indent string, isLast bool) {
		// Print the current node with the appropriate indentation and prefix.
		fmt.Print(indent)
		if isLast {
			fmt.Print("\\-")
			indent += "  "
		} else {
			fmt.Print("|-")
			indent += "| "
		}
		fmt.Println(node.Name)

		// Find the children of the current node.
		var children []*Node
		for _, childNode := range t.nodes {
			if childNode.ParentID == node.ID {
				children = append(children, childNode)
			}
		}

		// Recursively print the children with the updated indentation.
		for i, child := range children {
			printTree(child, indent, i == len(children)-1)
		}
	}

	// Start the DFS from the root node with no indentation and no special treatment for the last child.
	printTree(t.root, "", false)
}
