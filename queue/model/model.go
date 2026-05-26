package model

type Queue struct {
	Data  []any
	Front int
	Rear  int
	Size  int
	Cap   int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		Data:  make([]any, cap),
		Front: 0,
		Rear:  0,
		Size:  0,
		Cap:   cap,
	}
}
