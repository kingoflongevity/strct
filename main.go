package main

import (
	links "relax/Links"
	stack "relax/Stack"
)

func main() {
	testLinks()
	testStack()
}

func testLinks() {
	l := links.NewLinkService()
	l.PrintList(l.NewHeadNode())
}

func testStack() {
	s := stack.NewStackService()
	stack := s.NewStack()
	s.Push(stack, 1)
	s.Push(stack, 2)
	s.Push(stack, 3)
	s.Pop(stack)
	s.PrintStack(stack)
}
