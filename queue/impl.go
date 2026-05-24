package queue

import (
	"errors"
	"relax/queue/model"
)

type queueService struct{}

func (q *queueService) NewQueue(cap int) *model.Queue {
	return model.NewQueue(cap)
}

func (q *queueService) Enqueue(queue *model.Queue, val int) {
	if queue.Size == queue.Cap {
		return
	}
	queue.Data[queue.Rear] = val
	queue.Rear = (queue.Rear + 1) % queue.Cap
	queue.Size++
}

func (q *queueService) Dequeue(queue *model.Queue) (any, error) {
	if queue.Size == 0 {
		return nil, errors.New("dequeue from empty queue")
	}
	val := queue.Data[queue.Front]
	queue.Front = (queue.Front + 1) % queue.Cap
	queue.Size--
	return val, nil
}

func (q *queueService) Peek(queue *model.Queue) (any, error) {
	if queue.Size == 0 {
		return nil, errors.New("peek from empty queue")
	}
	return queue.Data[queue.Front], nil
}

func (q *queueService) IsEmpty(queue *model.Queue) bool {
	return queue.Size == 0
}

func (q *queueService) IsFull(queue *model.Queue) bool {
	return queue.Size == queue.Cap
}

func (q *queueService) Len(queue *model.Queue) int {
	return queue.Size
}

func NewQueueService() Queue {
	return &queueService{}
}
