package stack

import (
	"github.com/cloudfstrife/algs/fundamentals/dst/linklist"
)

//Stack Pushdown (LIFO) stack
type Stack struct {
	link *linklist.LinkList
}

//NewStack create an empty stack
func NewStack() *Stack {
	return &Stack{link: linklist.NewLinkList()}
}

//Push add an item
func (stack *Stack) Push(v interface{}) {
	stack.link.AddToEnd(v)
}

//Pop remove the most recently added item
func (stack *Stack) Pop() (interface{}, error) {
	return stack.link.RemoveFromEnd()
}

//IsEmpty  return true if the stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.link.IsEmpty()
}

//Size return the  number of items in the stack
func (stack *Stack) Size() int {
	return stack.link.Size()
}
