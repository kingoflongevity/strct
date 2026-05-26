package tree

import "relax/Tree/model"

type TreeService interface {
	NewTree() *model.TreeNode
	LevelOrder(root *model.TreeNode) []int
	PrintLevelOrder(root *model.TreeNode)
	PreOrder(root *model.TreeNode) []int
	InOrder(root *model.TreeNode) []int
	PostOrder(root *model.TreeNode) []int
}
