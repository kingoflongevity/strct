package links

import (
	"fmt"
	"relax/Links/model"
)

type linkService struct{}

func (l *linkService) NewLinkNode(val int) *model.LinkNode {
	return model.NewLinkNode(val)
}

func (l *linkService) AddLinkNode(node *model.LinkNode, val int) *model.LinkNode {
	node1 := model.NewLinkNode(val)
	node.Next = node1
	return node1
}

func (l *linkService) InsertNode(node *model.LinkNode, val int) *model.LinkNode {
	if node == nil {
		return nil
	}
	newNode := l.NewLinkNode(val)
	newNode.Next = node.Next
	node.Next = newNode

	return node
}

func (l *linkService) NewHeadNode() *model.LinkNode {
	return model.NewLinkNode(0)
}

func (l *linkService) AccessNode(head *model.LinkNode, post int) *model.LinkNode {
	curr := head
	if head == nil {
		return nil
	}
	for i := 1; i < post; i++ {
		curr = curr.Next
	}
	return curr

}

func (l *linkService) PrintList(head *model.LinkNode) {
	curr := head
	for curr != nil {
		fmt.Printf("[%d]", curr.Val)
		if curr.Next != nil {
			fmt.Print(" → ")
		} else {
			fmt.Println(" → nil")
		}
		curr = curr.Next
	}
}
func NewLinkService() Links {
	return &linkService{}
}
