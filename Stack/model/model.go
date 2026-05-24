package model

type Stack struct {
	Data []int
	Top  int
}

func NewStack() *Stack {
	return &Stack{
		make([]int, 0),
		0,
	}
}
