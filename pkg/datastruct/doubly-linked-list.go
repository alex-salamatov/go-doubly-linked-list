package datastruct

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
