package queue

import (
	"github.com/cloudfstrife/algs/fundamentals/dst/linklist"
)

//Queue FIFO queue
type Queue struct {
	link *linklist.LinkList
}

//NewQueue  create an empty queue
func NewQueue() *Queue {
	return &Queue{link: linklist.NewLinkList()}
}

//EnQueue  add an item
func (queue *Queue) EnQueue(v interface{}) {
	queue.link.AddToEnd(v)
}

//DeQueue  remove the least recently added item
func (queue *Queue) DeQueue() (interface{}, error) {
	return queue.link.RemoveFromHead()
}

//IsEmpty return true if the queue is empty
func (queue *Queue) IsEmpty() bool {
	return queue.link.IsEmpty()
}

//Size return  the number of items in the queue
func (queue *Queue) Size() int {
	return queue.link.Size()
}
