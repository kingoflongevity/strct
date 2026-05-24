package model

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewLinkNode(val int) *LinkNode {
	return &LinkNode{
		Val:  val,
		Next: nil,
	}
}
