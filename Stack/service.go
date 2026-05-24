package stack

import "relax/Stack/model"

type Stack interface {
	NewStack() *model.Stack
	Push(stack *model.Stack, val int)
	Pop(stack *model.Stack) any
	Peek(stack *model.Stack) any
	IsEmpty(stack *model.Stack) bool
	LenStack(stack *model.Stack) int
	PrintStack(stack *model.Stack)
}
