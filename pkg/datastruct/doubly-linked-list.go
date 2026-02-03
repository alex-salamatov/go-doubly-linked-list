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

func (dll *DoublyLinkedList) IsValidIndex(i int) bool {
	if i < 0 || i >= dll.Len() {
		return false
	}

	return true
}

func (dll *DoublyLinkedList) getNodeAt(pos int) *Dnode {
	if dll.Len() == 0 {
		return nil
	}

	if pos < dll.Len()/2 { //navigate from head to tail
		n := dll.first

		for i := 0; i < pos; i++ {
			n = n.next
		}

		return n
	} else { //navigate from tail to head
		n := dll.last

		for i := dll.len - 1; i > pos; i-- {
			n = n.prev
		}

		return n
	}
}

func (dll *DoublyLinkedList) swap(i, j int) error {
	if !dll.IsValidIndex(i) || !dll.IsValidIndex(j) {
		return errors.New("some index is out of range")
	}

	ni := dll.getNodeAt(i)
	nj := dll.getNodeAt(j)

	tmp := ni.value
	ni.value = nj.value
	nj.value = tmp

	return nil
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
		node.prev = dll.last
		dll.last = node
		dll.len++
	}
}

func (dll *DoublyLinkedList) At(i int) (int, error) {
	if !dll.IsValidIndex(i) {
		return 0, errors.New("index is out of range")
	}

	if dll.Len() == 0 {
		return 0, errors.New("cannot return an element: the list is empty")
	}

	return dll.getNodeAt(i).value, nil
}

func (dll *DoublyLinkedList) RemoveAt(i int) error {
	if !dll.IsValidIndex(i) {
		return errors.New("index is out of range")
	}

	if i == 0 {
		dll.first.next.prev = nil
		dll.first = dll.first.next
		dll.len--
		return nil
	}

	if i == dll.Len()-1 {
		dll.last.prev.next = nil
		dll.last = dll.last.prev
		dll.len--
		return nil
	}

	node := dll.getNodeAt(i)
	node.prev.next = node.next
	node.next.prev = node.prev
	dll.len--
	return nil
}

// i == -1 means insert as front element
func (dll *DoublyLinkedList) InsertAfter(i, value int) error {
	if !dll.IsValidIndex(i) && i != -1 {
		return errors.New("index is out of range")
	}

	node := MakeDnode(value)

	if (i == -1 && dll.Len() == 0) || (i == dll.Len()-1) {
		dll.Append(value)
		return nil
	}

	if i == -1 { //insert in front of list ...
		node.next = dll.first
		node.next.prev = node
		dll.first = node
		dll.len++
		return nil
	}

	afterNode := dll.getNodeAt(i)
	node.next = afterNode.next
	afterNode.next = node
	node.prev = afterNode
	node.next.prev = node
	dll.len++
	return nil
}
