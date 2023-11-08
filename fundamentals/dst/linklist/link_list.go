/*
Package linklist 封装了基础的单向链表及相关api
主要数据结构
* Node 链表节点
*/
package linklist

import (
	"errors"
	"sync"
)

var (
	//ErrEmpty return this error when remote item from linklist if it is empty
	ErrEmpty = errors.New("Empty")
)

//Node link list node
type Node struct {
	Val  interface{}
	Next *Node
}

//NewNode create a link list node
func NewNode(v interface{}) *Node {
	return &Node{Val: v}
}

//LinkList link list
type LinkList struct {
	lock   *sync.Mutex
	Head   *Node
	End    *Node
	Length int
}

//NewLinkList create an empty link list
func NewLinkList() *LinkList {
	return &LinkList{lock: &sync.Mutex{}, Head: nil, End: nil, Length: 0}
}

//AddToHead add a node to head
func (l *LinkList) AddToHead(v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	node := NewNode(v)
	if l.Length == 0 {
		l.Head = node
		l.End = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.Length = l.Length + 1
}

//AddToEnd add a node to end
func (l *LinkList) AddToEnd(v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	node := NewNode(v)
	if l.Length == 0 {
		l.Head = node
		l.End = node
	} else {
		l.End.Next = node
		l.End = node
	}
	l.Length = l.Length + 1
}

//RemoveFromHead remove a node from head
func (l *LinkList) RemoveFromHead() (interface{}, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	var v interface{}
	if l.Length <= 0 {
		return v, ErrEmpty
	} else if l.Length == 1 {
		v = l.Head.Val
		l.Head = nil
		l.End = nil
		l.Length = 0
	} else {
		v = l.Head.Val
		l.Head = l.Head.Next
		l.Length = l.Length - 1
	}
	return v, nil
}

//RemoveFromEnd remove a node from end
func (l *LinkList) RemoveFromEnd() (interface{}, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	var v interface{}
	if l.Length <= 0 {
		return v, ErrEmpty
	} else if l.Length == 1 {
		v = l.Head.Val
		l.Head = nil
		l.End = nil
		l.Length = 0
	} else {
		perNode := l.Head
		for i := 1; i < l.Length-1; i++ {
			perNode = perNode.Next
		}
		v = l.End.Val
		l.End = perNode
		l.End.Next = nil
		l.Length = l.Length - 1
	}
	return v, nil
}

//IsEmpty return true if link list is empty
func (l *LinkList) IsEmpty() bool {
	return l.Length == 0
}

//Size return number of items in the link list
func (l *LinkList) Size() int {
	return l.Length
}
