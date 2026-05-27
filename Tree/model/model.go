package model

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ArrayTree struct {
	Tree []any
}

func NewArrayTree() *ArrayTree {
	return &ArrayTree{
		Tree: []any{},
	}
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
