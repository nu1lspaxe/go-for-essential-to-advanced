package algorithm

import "fmt"

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type Node_RBT struct {
	Value  int
	Color  Color
	Parent *Node_RBT
	Left   *Node_RBT
	Right  *Node_RBT
}

type RedBlackTree struct {
	Root *Node_RBT
}

func NewRBTNode(value int, color Color) *Node_RBT {
	return &Node_RBT{Value: value, Color: color}
}

// Insert inserts a value into the Red-Black Tree
func (rbt *RedBlackTree) Insert(value int) {
	newNode := NewRBTNode(value, Red)
	if rbt.Root == nil {
		newNode.Color = Black
		rbt.Root = newNode
		return
	}
	insertRBTNode(rbt.Root, newNode)
	rbt.fixInsert(newNode)
}

// insertRBTNode inserts a node in the BST order
func insertRBTNode(root, newNode *Node_RBT) {
	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
			newNode.Parent = root
		} else {
			insertRBTNode(root.Left, newNode)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode
			newNode.Parent = root
		} else {
			insertRBTNode(root.Right, newNode)
		}
	}
}

// fixInsert fixes Red-Black Tree violations after insertion
func (rbt *RedBlackTree) fixInsert(node *Node_RBT) {
	for node != rbt.Root && node.Parent.Color == Red {
		grandparent := node.Parent.Parent
		if node.Parent == grandparent.Left {
			uncle := grandparent.Right
			if uncle != nil && uncle.Color == Red {
				// Case 1: Uncle is red
				node.Parent.Color = Black
				uncle.Color = Black
				grandparent.Color = Red
				node = grandparent
			} else {
				// Case 2 & 3: Uncle is black
				if node == node.Parent.Right {
					node = node.Parent
					rbt.leftRotate(node)
				}
				node.Parent.Color = Black
				grandparent.Color = Red
				rbt.rightRotate(grandparent)
			}
		} else {
			uncle := grandparent.Left
			if uncle != nil && uncle.Color == Red {
				// Case 1: Uncle is red
				node.Parent.Color = Black
				uncle.Color = Black
				grandparent.Color = Red
				node = grandparent
			} else {
				// Case 2 & 3: Uncle is black
				if node == node.Parent.Left {
					node = node.Parent
					rbt.rightRotate(node)
				}
				node.Parent.Color = Black
				grandparent.Color = Red
				rbt.leftRotate(grandparent)
			}
		}
	}
	rbt.Root.Color = Black
}

// leftRotate performs a left rotation
func (rbt *RedBlackTree) leftRotate(node *Node_RBT) {
	rightChild := node.Right
	node.Right = rightChild.Left
	if rightChild.Left != nil {
		rightChild.Left.Parent = node
	}
	rightChild.Parent = node.Parent
	if node.Parent == nil {
		rbt.Root = rightChild
	} else if node == node.Parent.Left {
		node.Parent.Left = rightChild
	} else {
		node.Parent.Right = rightChild
	}
	rightChild.Left = node
	node.Parent = rightChild
}

// rightRotate performs a right rotation
func (rbt *RedBlackTree) rightRotate(node *Node_RBT) {
	leftChild := node.Left
	node.Left = leftChild.Right
	if leftChild.Right != nil {
		leftChild.Right.Parent = node
	}
	leftChild.Parent = node.Parent
	if node.Parent == nil {
		rbt.Root = leftChild
	} else if node == node.Parent.Right {
		node.Parent.Right = leftChild
	} else {
		node.Parent.Left = leftChild
	}
	leftChild.Right = node
	node.Parent = leftChild
}

func (rbt *RedBlackTree) InOrderTraversal(node *Node_RBT) {
	if node == nil {
		return
	}

	rbt.InOrderTraversal(node.Left)
	fmt.Printf("%d(%v) ", node.Value, colorToString(node.Color))
	rbt.InOrderTraversal(node.Right)
}

func colorToString(color Color) string {
	if color == Red {
		return "R"
	}
	return "B"
}
