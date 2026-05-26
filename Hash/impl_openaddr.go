package hash

import (
	"fmt"
	"relax/Hash/model"
)

type openAddrService struct{}

func (o *openAddrService) NewOpenAddrTable(cap int) *model.OpenAddrTable {
	return model.NewOpenAddrTable(cap)
}

func (o *openAddrService) hashMod(key, cap int) int {
	return key % cap
}

func (o *openAddrService) probe(table *model.OpenAddrTable, key int) int {
	index := o.hashMod(key, table.Cap)
	for range table.Cap {
		entry := table.Data[index]
		if entry.Status == model.Empty {
			return -1
		}
		if entry.Status == model.Occupied && entry.Pair.Key == key {
			return index
		}
		index = (index + 1) % table.Cap
	}
	return -1
}

func (o *openAddrService) Insert(table *model.OpenAddrTable, key, value int) {
	if table.Size >= table.Cap {
		return
	}
	index := o.hashMod(key, table.Cap)
	for range table.Cap {
		entry := table.Data[index]
		if entry.Status != model.Occupied {
			entry.Pair = &model.Pair{Key: key, Value: value}
			entry.Status = model.Occupied
			table.Size++
			return
		}
		if entry.Pair.Key == key {
			entry.Pair.Value = value
			return
		}
		index = (index + 1) % table.Cap
	}
}

func (o *openAddrService) Get(table *model.OpenAddrTable, key int) *model.Pair {
	index := o.probe(table, key)
	if index == -1 {
		return nil
	}
	return table.Data[index].Pair
}

func (o *openAddrService) Contains(table *model.OpenAddrTable, key int) bool {
	return o.Get(table, key) != nil
}

func (o *openAddrService) Delete(table *model.OpenAddrTable, key int) {
	index := o.probe(table, key)
	if index == -1 {
		return
	}
	table.Data[index].Status = model.Deleted
	table.Data[index].Pair = nil
	table.Size--
}

func (o *openAddrService) Size(table *model.OpenAddrTable) int {
	return table.Size
}

func (o *openAddrService) Keys(table *model.OpenAddrTable) []int {
	var keys []int
	for i := 0; i < table.Cap; i++ {
		if table.Data[i].Status == model.Occupied {
			keys = append(keys, table.Data[i].Pair.Key)
		}
	}
	return keys
}

func (o *openAddrService) Values(table *model.OpenAddrTable) []int {
	var values []int
	for i := 0; i < table.Cap; i++ {
		if table.Data[i].Status == model.Occupied {
			values = append(values, table.Data[i].Pair.Value)
		}
	}
	return values
}

func (o *openAddrService) Pairs(table *model.OpenAddrTable) []*model.Pair {
	var pairs []*model.Pair
	for i := 0; i < table.Cap; i++ {
		if table.Data[i].Status == model.Occupied {
			pairs = append(pairs, table.Data[i].Pair)
		}
	}
	return pairs
}

func (o *openAddrService) PrintTable(table *model.OpenAddrTable) {
	fmt.Printf("OpenAddrTable (Size=%d, Cap=%d):\n", table.Size, table.Cap)
	for i := 0; i < table.Cap; i++ {
		entry := table.Data[i]
		switch entry.Status {
		case model.Empty:
			fmt.Printf("  [%d]: ·\n", i)
		case model.Deleted:
			fmt.Printf("  [%d]: ✗\n", i)
		case model.Occupied:
			fmt.Printf("  [%d]: (%d:%d)\n", i, entry.Pair.Key, entry.Pair.Value)
		}
	}
}

func NewOpenAddrService() OpenAddressing {
	return &openAddrService{}
}
