package queue

import (
	"errors"
	"fmt"
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

func (q *queueService) PrintQueue(queue *model.Queue) {
	// Data row
	fmt.Printf("Index: ")
	for i := 0; i < queue.Cap; i++ {
		fmt.Printf("  %d  ", i)
	}
	fmt.Println()

	fmt.Printf("Data:  ")
	for i := 0; i < queue.Cap; i++ {
		occupied := false
		for j := 0; j < queue.Size; j++ {
			if (queue.Front+j)%queue.Cap == i {
				occupied = true
				break
			}
		}
		if occupied {
			fmt.Printf(" %2d  ", queue.Data[i])
		} else {
			fmt.Printf("  ·  ")
		}
	}
	fmt.Println()

	// Front arrow
	fmt.Printf("Front: ")
	for i := 0; i < queue.Cap; i++ {
		if i == queue.Front && queue.Size > 0 {
			fmt.Printf("  ↑  ")
		} else {
			fmt.Printf("     ")
		}
	}
	fmt.Println()

	// Rear arrow
	fmt.Printf("Rear:  ")
	for i := 0; i < queue.Cap; i++ {
		if i == (queue.Rear+queue.Cap-1)%queue.Cap && queue.Size > 0 {
			fmt.Printf("  ↑  ")
		} else {
			fmt.Printf("     ")
		}
	}
	fmt.Println()

	fmt.Printf("Size=%d  Cap=%d  Front=%d  Rear=%d\n", queue.Size, queue.Cap, queue.Front, queue.Rear)
}
func NewQueueService() Queue {
	return &queueService{}
}
