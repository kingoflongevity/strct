package queue

import "relax/queue/model"

type Queue interface {
	NewQueue(cap int) *model.Queue
	Enqueue(queue *model.Queue, val int)
	Dequeue(queue *model.Queue) (any, error)
	Peek(queue *model.Queue) (any, error)
	IsEmpty(queue *model.Queue) bool
	IsFull(queue *model.Queue) bool
	Len(queue *model.Queue) int
}
