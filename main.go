package main

import (
	"fmt"
	links "relax/Links"
	stack "relax/Stack"
	"relax/queue"
)

func main() {
	fmt.Println("=== Linked List ===")
	testLinks()

	fmt.Println("\n=== Stack ===")
	testStack()

	fmt.Println("\n=== Circular Queue ===")
	testQueue()
}

func testLinks() {
	l := links.NewLinkService()
	head := l.NewHeadNode()
	l.PrintList(head)
}

func testStack() {
	s := stack.NewStackService()
	st := s.NewStack()
	s.Push(st, 1)
	s.Push(st, 2)
	s.Push(st, 3)
	s.Pop(st)
	s.PrintStack(st)
}

func testQueue() {
	q := queue.NewQueueService()
	qu := q.NewQueue(5)

	// Enqueue 3 elements
	q.Enqueue(qu, 10)
	q.Enqueue(qu, 20)
	q.Enqueue(qu, 30)
	fmt.Println("After enqueue 10, 20, 30:")
	q.PrintQueue(qu)
	fmt.Println()

	// Dequeue 2 elements
	q.Dequeue(qu)
	q.Dequeue(qu)
	fmt.Println("After dequeue 2 times:")
	q.PrintQueue(qu)
	fmt.Println()

	// Enqueue 3 more — will wrap around
	q.Enqueue(qu, 40)
	q.Enqueue(qu, 50)
	q.Enqueue(qu, 60)
	fmt.Println("After enqueue 40, 50, 60 (circular wrap):")
	q.PrintQueue(qu)
}
