package hash

import "relax/Hash/model"

type HashTable interface {
	NewHashTable(cap int) *model.HashTable
	Insert(table *model.HashTable, key, value int)
	Get(table *model.HashTable, key int) *model.Pair
	Delete(table *model.HashTable, key int)
	Contains(table *model.HashTable, key int) bool
	Size(table *model.HashTable) int
	Cap(table *model.HashTable) int
	Keys(table *model.HashTable) []int
	Values(table *model.HashTable) []int
	Pairs(table *model.HashTable) []*model.Pair
	PrintTable(table *model.HashTable)
}

type OpenAddressing interface {
	NewOpenAddrTable(cap int) *model.OpenAddrTable
	Insert(table *model.OpenAddrTable, key, value int)
	Get(table *model.OpenAddrTable, key int) *model.Pair
	Delete(table *model.OpenAddrTable, key int)
	Contains(table *model.OpenAddrTable, key int) bool
	Size(table *model.OpenAddrTable) int
	Keys(table *model.OpenAddrTable) []int
	Values(table *model.OpenAddrTable) []int
	Pairs(table *model.OpenAddrTable) []*model.Pair
	PrintTable(table *model.OpenAddrTable)
}
