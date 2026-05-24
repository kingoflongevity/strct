package links

import "relax/Links/model"

type Links interface {
	NewHeadNode() *model.LinkNode
	NewLinkNode(val int) *model.LinkNode
	AddLinkNode(node *model.LinkNode, val int) *model.LinkNode
	InsertNode(node *model.LinkNode, val int) *model.LinkNode
	PrintList(head *model.LinkNode)
	AccessNode(head *model.LinkNode, post int) *model.LinkNode
}
