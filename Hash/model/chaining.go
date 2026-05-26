package model

type Node struct {
	Pair *Pair
	Next *Node
}

type HashTable struct {
	Buckets []*Node
	Size    int
	Cap     int
}

func NewHashTable(cap int) *HashTable {
	return &HashTable{
		Buckets: make([]*Node, cap),
		Size:    0,
		Cap:     cap,
	}
}
