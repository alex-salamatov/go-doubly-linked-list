package datastruct

import (
	"fmt"
	"math/rand/v2"
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

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{0, nil, nil}
}

func NewDoublyLinkedListFromSlice(a *[]int) *DoublyLinkedList {
	l := DoublyLinkedList{0, nil, nil}
	for _, val := range *a {
		l.Append(val)
	}
	return &l
}

func NewDnode(val int) *Dnode {
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

func (dll *DoublyLinkedList) swap(i, j int) {
	if !dll.IsValidIndex(i) || !dll.IsValidIndex(j) {
		panic("some index is out of range")
	}

	ni := dll.getNodeAt(i)
	nj := dll.getNodeAt(j)

	tmp := ni.value
	ni.value = nj.value
	nj.value = tmp
}

func (dll *DoublyLinkedList) Len() int {
	return dll.len
}

func (dll *DoublyLinkedList) Append(val int) {
	node := NewDnode(val)

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

func (dll *DoublyLinkedList) At(i int) int {
	if !dll.IsValidIndex(i) {
		panic("index is out of range")
	}

	if dll.Len() == 0 {
		panic("cannot return an element: the list is empty")
	}

	return dll.getNodeAt(i).value
}

func (dll *DoublyLinkedList) RemoveAt(i int) {
	if !dll.IsValidIndex(i) {
		panic("index is out of range")
	}

	if i == 0 {
		dll.first.next.prev = nil
		dll.first = dll.first.next
		dll.len--
		return
	}

	if i == dll.Len()-1 {
		dll.last.prev.next = nil
		dll.last = dll.last.prev
		dll.len--
		return
	}

	node := dll.getNodeAt(i)
	node.prev.next = node.next
	node.next.prev = node.prev
	dll.len--
}

// i == -1 means insert as front element
func (dll *DoublyLinkedList) InsertAfter(i, value int) {
	if !dll.IsValidIndex(i) && i != -1 {
		panic("index is out of range")
	}

	node := NewDnode(value)

	if (i == -1 && dll.Len() == 0) || (i == dll.Len()-1) {
		dll.Append(value)
		return
	}

	if i == -1 { //insert in front of list ...
		node.next = dll.first
		node.next.prev = node
		dll.first = node
		dll.len++
		return
	}

	afterNode := dll.getNodeAt(i)
	node.next = afterNode.next
	afterNode.next = node
	node.prev = afterNode
	node.next.prev = node
	dll.len++
}

func (dll *DoublyLinkedList) Print() {
	fmt.Printf("DList: Len = %d, Elements: ", dll.Len())
	for i := 0; i < dll.Len(); i++ {
		val := dll.At(i)
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

func (dll *DoublyLinkedList) QuickSort() {
	if dll.Len() <= 1 {
		return
	}

	pi := rand.IntN(dll.Len()) //random index of element which we take as pivot = pivot index
	pivot := dll.At(pi)

	//	dll.Print()
	//	fmt.Printf(" original pivot index = %d, pivot value = %d ", pi, pivot)

	var spi int //sorted pivot index

	node := dll.getNodeAt(0)

	for i := 0; i < dll.Len(); i++ {
		if node.value < pivot {
			spi++
		}
		node = node.next
	}

	if pi != spi {
		dll.swap(pi, spi) //pivot is placed to its place
	}

	//	fmt.Printf("  sortedPivotIndex = %d, N LeftElements = %d N RightElements = %d \n", spi, spi, dll.Len()-spi-1)

	left := NewDoublyLinkedList()
	right := NewDoublyLinkedList()
	pivotNode := dll.getNodeAt(spi)

	if spi > 0 {
		left.first = dll.first     //elements on the left of pivot
		left.last = pivotNode.prev //elements on the left of pivot
		left.len = spi
	}

	if spi < dll.Len()-1 {
		right.first = pivotNode.next //elements on the right of pivot
		right.last = dll.last
		right.len = dll.Len() - spi - 1
	}

	if left.Len() > 0 && right.Len() > 0 {
		//start moving elements from left to right parts if needed
		for i, jr := 0, 0; i < spi; i++ {

			if left.At(i) < pivot {
				continue //left[i] can stay in its position on the left of pivot
			}

			for k := jr; k < right.Len(); k++ {

				if right.At(k) < pivot {
					dll.swap(i, k+spi+1)
					jr = k + 1 //index of element of the right from pivot slice, starting from which there still might be elements to be moved to the left of pivot
					break
				}
			}
		}
	}

	//	fmt.Printf("Output: Pivot = %d\n", pivot)
	//left.Print()
	//right.Print()

	left.QuickSort()  //apply quick sort to the left  from pivot partition
	right.QuickSort() //apply quick sort to the right from pivot partition
}

func (dll *DoublyLinkedList) isSameAs(a *[]int) bool {
	if len(*a) == 0 && dll.Len() == 0 {
		return true
	}

	if len(*a) != dll.Len() {
		return false
	}

	node := dll.first

	for _, val := range *a {
		if val != node.value {
			return false
		}
		node = node.next
	}

	return true
}

func (dll *DoublyLinkedList) GetSliceValues() *[]int {
	a := make([]int, dll.Len())

	node := dll.first

	for i := range a {
		a[i] = node.value
		node = node.next
	}

	return &a
}
