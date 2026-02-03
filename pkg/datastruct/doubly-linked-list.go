package datastruct

import (
	"errors"
)

type Dnode struct {
	value int
	next  *Dnode
	prev  *Dnode
}

type DoublyLinkedList struct {
	len   int
	first *Dnode
	last  *Dnode
}

func MakeDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{0, nil, nil}
}

func MakeDnode(val int) *Dnode {
	return &Dnode{val, nil, nil}
}

func (dll *DoublyLinkedList) Len() int {
	return dll.len
}

func (dll *DoublyLinkedList) Append(val int) {
	node := MakeDnode(val)

	if dll.Len() == 0 {
		dll.first = node
		dll.last = node
		dll.len = 1
	} else {
		dll.last.next = node
		dll.last = node
		dll.len++
	}
}

func (dll *DoublyLinkedList) getNodeAt(pos int) *Dnode {
	if dll.Len() == 0 {
		return nil
	}

	if pos < dll.Len()/2 { //travers from head to tail
		n := dll.first

		for i := 0; i <= pos; i++ {
			n = n.next
		}

		return n
	} else { //travers from tail to head
		n := dll.last

		for i := dll.len - 1; i > pos; i-- {
			n = n.prev
		}

		return n
	}
}

func (dll *DoublyLinkedList) At(i int) (int, error) {
	if i < 0 || i >= dll.Len() {
		return 0, errors.New("index is out of range")
	}

	if dll.Len() == 0 {
		return 0, errors.New("cannot return an element: the list is empty")
	}

	return dll.getNodeAt(i).value, nil
}
