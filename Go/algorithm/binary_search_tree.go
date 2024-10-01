package algorithm

import "fmt"

/*
Binary Search Tree (BST)

1. Define the Tree Node
2. Insert Operation
3. Search Operation
4. In-Order Traversal
*/

// Node represents a node in the binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a new node to the binary search tree
func (n *Node) Insert(value int) {
	if value <= n.Value {
		// Insert in the left subtree
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		// Insert in the right subtree
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search checks if a value exists in the binary search tree
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value < n.Value {
		// search in the left subtree
		return n.Left.Search(value)
	} else if value > n.Value {
		// search in the right subtree
		return n.Right.Search(value)
	}

	// If value equals n.Value
	return true
}

// InOrderTraversal prints the nodes in the binary search tree
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()  // visit the left subtree
	fmt.Printf("%d ", n.Value) // visit the root node
	n.Right.InOrderTraversal() // visit the right subtree
}

func BinarySearchTree() {
	root := &Node{Value: 10}

	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(17)

	fmt.Println("In-order traversal:")
	root.InOrderTraversal()
	fmt.Println()

	fmt.Println("Search for 7:", root.Search(7))
	fmt.Println("Search for 20:", root.Search(20))
}
