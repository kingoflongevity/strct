package hash

import (
	"fmt"
	"relax/Hash/model"
)

type hashTableService struct{}

func (h *hashTableService) NewHashTable(cap int) *model.HashTable {
	return model.NewHashTable(cap)
}

func (h *hashTableService) hashModfunc(key, cap int) int {
	return key % cap
}

func (h *hashTableService) Insert(table *model.HashTable, key, value int) {
	index := h.hashModfunc(key, table.Cap)
	curr := table.Buckets[index]
	for curr != nil {
		if curr.Pair.Key == key {
			curr.Pair.Value = value
			return
		}
		curr = curr.Next
	}
	table.Buckets[index] = &model.Node{
		Pair: &model.Pair{Key: key, Value: value},
		Next: table.Buckets[index],
	}
	table.Size++
}

func (h *hashTableService) Get(table *model.HashTable, key int) *model.Pair {
	index := h.hashModfunc(key, table.Cap)
	curr := table.Buckets[index]
	for curr != nil {
		if curr.Pair.Key == key {
			return curr.Pair
		}
		curr = curr.Next
	}
	return nil
}

func (h *hashTableService) Contains(table *model.HashTable, key int) bool {
	return h.Get(table, key) != nil
}

func (h *hashTableService) Delete(table *model.HashTable, key int) {
	index := h.hashModfunc(key, table.Cap)
	curr := table.Buckets[index]
	var prev *model.Node
	for curr != nil {
		if curr.Pair.Key == key {
			if prev == nil {
				table.Buckets[index] = curr.Next
			} else {
				prev.Next = curr.Next
			}
			table.Size--
			return
		}
		prev = curr
		curr = curr.Next
	}
}

func (h *hashTableService) Size(table *model.HashTable) int {
	return table.Size
}

func (h *hashTableService) Cap(table *model.HashTable) int {
	return table.Cap
}

func (h *hashTableService) Keys(table *model.HashTable) []int {
	var keys []int
	for i := 0; i < table.Cap; i++ {
		curr := table.Buckets[i]
		for curr != nil {
			keys = append(keys, curr.Pair.Key)
			curr = curr.Next
		}
	}
	return keys
}

func (h *hashTableService) Values(table *model.HashTable) []int {
	var values []int
	for i := 0; i < table.Cap; i++ {
		curr := table.Buckets[i]
		for curr != nil {
			values = append(values, curr.Pair.Value)
			curr = curr.Next
		}
	}
	return values
}

func (h *hashTableService) Pairs(table *model.HashTable) []*model.Pair {
	var pairs []*model.Pair
	for i := 0; i < table.Cap; i++ {
		curr := table.Buckets[i]
		for curr != nil {
			pairs = append(pairs, curr.Pair)
			curr = curr.Next
		}
	}
	return pairs
}

func (h *hashTableService) PrintTable(table *model.HashTable) {
	fmt.Printf("HashTable (Size=%d, Cap=%d):\n", table.Size, table.Cap)
	for i := 0; i < table.Cap; i++ {
		fmt.Printf("  [%d]:", i)
		curr := table.Buckets[i]
		if curr == nil {
			fmt.Println(" nil")
			continue
		}
		for curr != nil {
			fmt.Printf(" → (%d:%d)", curr.Pair.Key, curr.Pair.Value)
			curr = curr.Next
		}
		fmt.Println()
	}
}

func NewHashTableService() HashTable {
	return &hashTableService{}
}
