package tree

import (
	"fmt"
	"relax/Tree/model"
	"relax/queue"
	queuemodel "relax/queue/model"
	"strings"
)

type treeService struct{}

func (t *treeService) NewTree() *model.TreeNode {
	return model.NewTreeNode(0)
}

type arrayTreeService struct{}

func (a *arrayTreeService) NewTree() *model.ArrayTree {
	return model.NewArrayTree()
}

func nodeLabel(node *model.TreeNode, isCurrent bool, isProcessed bool) string {
	if node == nil {
		return " · "
	}
	switch {
	case isCurrent:
		return fmt.Sprintf(">%d<", node.Val)
	case isProcessed:
		return fmt.Sprintf("[%d]", node.Val)
	default:
		return fmt.Sprintf(" %d ", node.Val)
	}
}

func printTreeLevels(levels [][]*model.TreeNode, highlight int, processedSet map[int]bool) {
	depth := len(levels)
	if depth == 0 {
		return
	}

	labelW := 3
	leafCount := 1 << (depth - 1)
	cellW := labelW * 2
	totalW := leafCount * cellW

	pos := make([][]int, depth)
	pos[depth-1] = make([]int, len(levels[depth-1]))
	for i := range pos[depth-1] {
		pos[depth-1][i] = i*cellW + cellW/2
	}
	for lev := depth - 2; lev >= 0; lev-- {
		pos[lev] = make([]int, len(levels[lev]))
		for i := range pos[lev] {
			li, ri := i*2, i*2+1
			l, r := 0, 0
			n := 0
			if li < len(pos[lev+1]) {
				l += pos[lev+1][li]
				n++
			}
			if ri < len(pos[lev+1]) {
				r += pos[lev+1][ri]
				n++
			}
			if n == 2 {
				pos[lev][i] = (l + r) / 2
			} else if n == 1 {
				pos[lev][i] = l + r
			}
		}
	}

	for lev := 0; lev < depth; lev++ {
		line := []byte(strings.Repeat(" ", totalW))
		for i, node := range levels[lev] {
			if node == nil {
				continue
			}
			_, isProc := processedSet[node.Val]
			label := nodeLabel(node, node.Val == highlight, isProc)
			start := pos[lev][i] - 1
			if start < 0 {
				start = 0
			}
			if start+3 > totalW {
				start = totalW - 3
			}
			copy(line[start:], label)
		}
		fmt.Println(string(line))

		if lev < depth-1 {
			branchLine := []byte(strings.Repeat(" ", totalW))
			for i, node := range levels[lev+1] {
				if node == nil {
					continue
				}
				p := pos[lev+1][i]
				if i%2 == 0 {
					if p-1 >= 0 {
						branchLine[p-1] = '/'
					}
				} else {
					if p+1 < totalW {
						branchLine[p+1] = '\\'
					} else if p < totalW {
						branchLine[p] = '\\'
					}
				}
			}
			fmt.Println(string(branchLine))
		}
	}
}

func collectLevels(root *model.TreeNode) [][]*model.TreeNode {
	var levels [][]*model.TreeNode
	cur := []*model.TreeNode{root}
	for len(cur) > 0 {
		levels = append(levels, cur)
		var next []*model.TreeNode
		for _, node := range cur {
			if node != nil {
				next = append(next, node.Left, node.Right)
			}
		}
		allNil := true
		for _, n := range next {
			if n != nil {
				allNil = false
				break
			}
		}
		if allNil {
			break
		}
		cur = next
	}
	return levels
}

func printQueueLine(qu *queuemodel.Queue) string {
	if qu.Size == 0 {
		return "[empty]"
	}
	var vals []string
	for j := 0; j < qu.Size; j++ {
		idx := (qu.Front + j) % qu.Cap
		node := qu.Data[idx].(*model.TreeNode)
		vals = append(vals, fmt.Sprintf("%d", node.Val))
	}
	return "Front→ [" + strings.Join(vals, ", ") + "] ←Rear"
}

func printStep(qu *queuemodel.Queue, levels [][]*model.TreeNode, processedSet map[int]bool, step int, op string, val int, result []int) {
	fmt.Printf("═══ [%d] %s: %d ═══\n", step, op, val)
	printTreeLevels(levels, val, processedSet)
	fmt.Println("Queue:", printQueueLine(qu))
	fmt.Printf("Size=%d  Front=%d  Rear=%d\n", qu.Size, qu.Front, qu.Rear)
	fmt.Println("Result:", result)
	fmt.Println()
}

func (t *treeService) LevelOrder(root *model.TreeNode) []int {
	qs := queue.NewQueueService()
	qu := qs.NewQueue(100)

	levels := collectLevels(root)
	processedSet := make(map[int]bool)

	var result []int

	qs.Enqueue(qu, root)
	printStep(qu, levels, processedSet, 0, "Enqueue root", root.Val, result)

	step := 1
	for !qs.IsEmpty(qu) {
		nodeAny, err := qs.Dequeue(qu)
		if err != nil {
			break
		}
		node := nodeAny.(*model.TreeNode)

		printStep(qu, levels, processedSet, step, "Dequeue", node.Val, result)
		step++

		result = append(result, node.Val)
		processedSet[node.Val] = true

		if node.Left != nil {
			qs.Enqueue(qu, node.Left)
			printStep(qu, levels, processedSet, step, "Enqueue left", node.Left.Val, result)
			step++
		}
		if node.Right != nil {
			qs.Enqueue(qu, node.Right)
			printStep(qu, levels, processedSet, step, "Enqueue right", node.Right.Val, result)
			step++
		}
	}
	return result
}

func (t *treeService) PrintLevelOrder(root *model.TreeNode) {
	vals := t.LevelOrder(root)
	fmt.Println("Level Order:", vals)
}

func (t *treeService) PreOrder(root *model.TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, root.Val)
	t.PreOrder(root.Left)
	t.PreOrder(root.Right)
	return result
}

func (t *treeService) InOrder(root *model.TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, root.Val)
	t.InOrder(root.Left)
	t.InOrder(root.Right)
	return result
}

func (t *treeService) PostOrder(root *model.TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, root.Val)
	t.PostOrder(root.Left)
	t.PostOrder(root.Right)
	return result
}

func (a *arrayTreeService) Val(arr *model.ArrayTree, i int) any {
	if i < 0 || i >= len(arr.Tree) {
		return nil
	}
	return arr.Tree[i]
}

func (a *arrayTreeService) Left(arr *model.ArrayTree, i int) int {
	return 2*i + 1
}

func (a *arrayTreeService) Right(arr *model.ArrayTree, i int) int {
	return 2*i + 2
}
func (a *arrayTreeService) Parent(arr *model.ArrayTree, i int) int {
	return (i - 1) / 2
}

func (a *arrayTreeService) Size(arr *model.ArrayTree) int {
	return len(arr.Tree)
}

func (a *arrayTreeService) LevelOrder(arr *model.ArrayTree) []int {
	var result []int
	for i := 0; i < a.Size(arr); i++ {
		fmt.Println(a.Val(arr, i))
		result = append(result, a.Val(arr, i).(int))
	}
	return result
}

/* 深度优先遍历 */
func (a *arrayTreeService) dfs(arr *model.ArrayTree, i int, order string, res *[]any, depth int) {
	indent := strings.Repeat("  ", depth)
	val := a.Val(arr, i)

	fmt.Printf("%sdfs(i=%d, val=%v, order=%s)\n", indent, i, val, order)

	if val == nil {
		fmt.Printf("%s  → nil, return\n", indent)
		return
	}

	if order == "pre" {
		fmt.Printf("%s  → pre ✓ 记录 val=%v\n", indent, val)
		*res = append(*res, val)
	} else {
		fmt.Printf("%s  → pre ✗\n", indent)
	}

	a.dfs(arr, a.Left(arr, i), order, res, depth+1)

	if order == "in" {
		fmt.Printf("%s  → in ✓ 记录 val=%v\n", indent, val)
		*res = append(*res, val)
	} else {
		fmt.Printf("%s  → in ✗\n", indent)
	}

	a.dfs(arr, a.Right(arr, i), order, res, depth+1)

	if order == "post" {
		fmt.Printf("%s  → post ✓ 记录 val=%v\n", indent, val)
		*res = append(*res, val)
	} else {
		fmt.Printf("%s  → post ✗\n", indent)
	}

	fmt.Printf("%s← return from i=%d(val=%v)\n", indent, i, val)
}

/* 前序遍历 */
func (a *arrayTreeService) PreOrder(arr *model.ArrayTree) []any {
	var res []any
	fmt.Println("\n========== PreOrder Traversal ==========")
	a.dfs(arr, 0, "pre", &res, 0)
	fmt.Printf("Result: %v\n", res)
	return res
}

/* 中序遍历 */
func (a *arrayTreeService) InOrder(arr *model.ArrayTree) []any {
	var res []any
	fmt.Println("\n========== InOrder Traversal ==========")
	a.dfs(arr, 0, "in", &res, 0)
	fmt.Printf("Result: %v\n", res)
	return res
}

/* 后序遍历 */
func (a *arrayTreeService) PostOrder(arr *model.ArrayTree) []any {
	var res []any
	fmt.Println("\n========== PostOrder Traversal ==========")
	a.dfs(arr, 0, "post", &res, 0)
	fmt.Printf("Result: %v\n", res)
	return res
}

// 二叉搜索树搜索
func (t *treeService) Search(node *model.TreeNode, val int) *model.TreeNode {
	if node == nil {
		return nil
	}
	for node != nil {
		if node.Val == val {
			break
		}
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}

// 二叉搜索树插入
func (t *treeService) Insert(root *model.TreeNode, val int) *model.TreeNode {
	if root == nil {
		return &model.TreeNode{Val: val}
	}

	cur := root
	for {
		if val == cur.Val {
			return root
		}
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = &model.TreeNode{Val: val}
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &model.TreeNode{Val: val}
				return root
			}
			cur = cur.Right
		}
	}
}

func NewTreeService() LinksTreeService {
	return &treeService{}
}

func NewArrayTreeService() ArrayTreeService {
	return &arrayTreeService{}
}
