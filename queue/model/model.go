package model

type Queue struct {
	Data  []int
	Front int
	Rear  int
	Size  int
	Cap   int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		Data:  make([]int, cap),
		Front: 0,
		Rear:  0,
		Size:  0,
		Cap:   cap,
	}
}
