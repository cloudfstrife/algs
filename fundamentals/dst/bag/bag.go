package bag

import (
	"github.com/cloudfstrife/algs/fundamentals/dst/linklist"
)

//Bag bag
type Bag struct {
	link *linklist.LinkList
}

//NewBag  create an empty bag
func NewBag() *Bag {
	return &Bag{link: linklist.NewLinkList()}
}

//Add  add an item
func (bag *Bag) Add(item interface{}) {
	bag.link.AddToEnd(item)
}

//IsEmpty return true if bag is empty
func (bag *Bag) IsEmpty() bool {
	return bag.link.IsEmpty()
}

//Size return the number of items in the bag
func (bag *Bag) Size() int {
	return bag.link.Size()
}
