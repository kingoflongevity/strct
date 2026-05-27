package tree

import "relax/Tree/model"

type LinksTreeService interface {
	NewTree() *model.TreeNode
	LevelOrder(root *model.TreeNode) []int
	PrintLevelOrder(root *model.TreeNode)
	PreOrder(root *model.TreeNode) []int
	InOrder(root *model.TreeNode) []int
	PostOrder(root *model.TreeNode) []int
}
type ArrayTreeService interface {
	NewTree() *model.ArrayTree
	PreOrder(arr *model.ArrayTree) []any
	InOrder(arr *model.ArrayTree) []any
	PostOrder(arr *model.ArrayTree) []any
	LevelOrder(arr *model.ArrayTree) []int
}
