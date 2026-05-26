package model

type EntryStatus int

const (
	Empty   EntryStatus = iota
	Occupied
	Deleted
)

type OpenAddrEntry struct {
	Pair   *Pair
	Status EntryStatus
}

type OpenAddrTable struct {
	Data []*OpenAddrEntry
	Size int
	Cap  int
}

func NewOpenAddrTable(cap int) *OpenAddrTable {
	data := make([]*OpenAddrEntry, cap)
	for i := 0; i < cap; i++ {
		data[i] = &OpenAddrEntry{Status: Empty}
	}
	return &OpenAddrTable{
		Data: data,
		Size: 0,
		Cap:  cap,
	}
}
