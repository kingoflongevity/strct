package main

import (
	"fmt"
	tree "relax/Tree"
	"relax/Tree/model"
)

func main() {
	// ts := tree.NewTreeService()

	// root := &model.TreeNode{Val: 1}
	// root.Left = &model.TreeNode{Val: 2}
	// root.Right = &model.TreeNode{Val: 3}
	// root.Left.Left = &model.TreeNode{Val: 4}
	// root.Left.Right = &model.TreeNode{Val: 5}
	// root.Right.Left = &model.TreeNode{Val: 6}
	// root.Right.Right = &model.TreeNode{Val: 7}

	// fmt.Println("Tree structure:")
	// fmt.Println("      1")
	// fmt.Println("     / \\")
	// fmt.Println("    2   3")
	// fmt.Println("   / \\ / \\")
	// fmt.Println("  4  5 6  7")
	// fmt.Println()

	// ts.PrintLevelOrder(root)

	// fmt.Println("\n\n========================================")
	// fmt.Println("Array-based Tree DFS Traversals")
	// fmt.Println("========================================")

	ats := tree.NewArrayTreeService()
	arr := model.NewArrayTree()
	arr.Tree = []any{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("Array: [1, 2, 3, 4, 5, 6, 7]")
	fmt.Println("        1")
	fmt.Println("       / \\")
	fmt.Println("      2   3")
	fmt.Println("     / \\ / \\")
	fmt.Println("    4  5 6  7")

	ats.PreOrder(arr)
	ats.InOrder(arr)
	ats.PostOrder(arr)
}
