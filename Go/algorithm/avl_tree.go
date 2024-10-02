package algorithm

import "fmt"

/*
AVL Tree

1. Insert
2. Rotations
3. Balance Factor
4. In-Order Traversal
*/

// Node_AVL represents a node in the AVL tree
type Node_AVL struct {
	Value  int
	Height int
	Left   *Node_AVL
	Right  *Node_AVL
}

// height returns the height of a node
func height(n *Node_AVL) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// getBalance calculates the balance factor of a node
func getBalance(n *Node_AVL) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

// rightRotate performs a right rotation
func rightRotate(y *Node_AVL) *Node_AVL {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(height(y.Left), height(y.Right))
	x.Height = max(height(x.Left), height(x.Right))

	return x
}

// leftRotate performs a left rotation
func leftRotate(x *Node_AVL) *Node_AVL {
	y := x.Right
	T2 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T2

	// Update heights
	x.Height = max(height(x.Left), height(x.Right))
	y.Height = max(height(y.Left), height(y.Right))

	return y
}

// insert inserts a new node with the given value
func insert(node *Node_AVL, value int) *Node_AVL {
	// Perform the normal BST insertion
	if node == nil {
		return &Node_AVL{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	} else {
		// Duplicate values are not allowed in BST
		return node
	}

	// Update the height of the ancestor node
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// Get the balance factor to check if this node became unbalanced
	balance := getBalance(node)

	// Left Left case (Right rotation)
	if balance > 1 && value < node.Left.Value {
		return rightRotate(node)
	}

	// Right Right case (left rotation)
	if balance < -1 && value > node.Right.Value {
		return leftRotate(node)
	}

	// Left Right case (Left-right rotation)
	if balance > 1 && value > node.Left.Value {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left case (Right-left rotation)
	if balance < -1 && value < node.Right.Value {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// inOrderTraversal prints the nodes in the tree in sorted order
func inOrderTraversal(n *Node_AVL) {
	if n == nil {
		return
	}
	inOrderTraversal(n.Left)
	fmt.Printf("%d ", n.Value)
	inOrderTraversal(n.Right)
}

func AVLTree() {
	var root *Node_AVL

	values := []int{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		root = insert(root, v)
	}

	fmt.Println("In-order traversal of the AVL Tree:")
	inOrderTraversal(root)
}
