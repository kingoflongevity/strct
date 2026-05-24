package stack

import (
	"fmt"
	"relax/Stack/model"
)

type stackService struct{}

func (s *stackService) NewStack() *model.Stack {
	return model.NewStack()
}

func (s *stackService) Push(stack *model.Stack, val int) {
	stack.Data = append(stack.Data, val)
	stack.Top++
}

func (s *stackService) IsEmpty(stack *model.Stack) bool {
	return stack.Top == 0
}

func (s *stackService) LenStack(stack *model.Stack) int {
	return stack.Top
}

func (s *stackService) Peek(stack *model.Stack) any {
	if s.IsEmpty(stack) {
		return nil
	}
	return stack.Data[stack.Top-1]
}

func (s *stackService) Pop(stack *model.Stack) any {
	if s.IsEmpty(stack) {
		return nil
	}
	stack.Top--
	val := stack.Data[stack.Top]
	stack.Data = stack.Data[:stack.Top]
	return val
}

func (s *stackService) PrintStack(stack *model.Stack) {
	for i := 0; i < stack.Top; i++ {
		fmt.Printf("[%d]", stack.Data[i])
		if i != stack.Top-1 {
			fmt.Print(" → ")
		} else {
			fmt.Println(" → nil")
		}
	}
}
func NewStackService() Stack {
	return &stackService{}
}
