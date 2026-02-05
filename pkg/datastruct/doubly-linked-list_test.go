package datastruct

import (
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()

	if l.len != 0 || l.first != nil || l.last != nil {
		t.Errorf("TestNewDoublyLinkedList: expected len = 0, first = nil, last = nil, got %d %v %v", l.len, l.first, l.last)
	}
}

func TestNewDnode(t *testing.T) {
	n := NewDnode(25)

	if n.value != 25 || n.next != nil || n.prev != nil {
		t.Errorf("TestNewDnode: expected value = 25, next = nil, prev = nil, got %d %v %v", n.value, n.next, n.prev)
	}
}

func TestAppend(t *testing.T) {
	l := NewDoublyLinkedList()

	l.Append(10)
	l.Append(20)

	if l.Len() != 2 {
		t.Errorf("TestAppend: expected size 2, got %d", l.Len())
	}

	if l.first.value != 10 || l.first.next.value != 20 {
		t.Errorf("TestAppend: values are not appended as expected. Expected list[0] = 10 and list[1] = 20, got %d and %d", l.first.value, l.first.next.value)
	}
}

func TestIsValidIndex(t *testing.T) {
	l := NewDoublyLinkedList()

	if l.IsValidIndex(0) || l.IsValidIndex(1) {
		t.Errorf("TestIsValidIndex: expected false for empty list, got true")
	}

	l.Append(1)
	l.Append(2)

	if l.IsValidIndex(-1) || l.IsValidIndex(2) {
		t.Errorf("TestIsValidIndex: expected false for list with two elements, got true")
	}
}

func TestQuickSort(t *testing.T) {
	l0 := NewDoublyLinkedList()
	l0.QuickSort()

	l0.Append(1)
	l0.QuickSort()

	l0.Append(0)
	l0.QuickSort()

	a0_result := []int{0, 1}
	if !l0.isSameAs(&a0_result) {
		t.Errorf("TestQuickSort: expected list %v, got %v", a0_result, *l0.GetSliceValues())
	}

	a1_input := []int{-56, -66, 134, 99, 153, 50, -16, 88, 60, -151, 91, -136, 6, 160, 140, -110, 33, 159, -168, 93, 132, -187, 96, 19, 34, 120, 96, -170, -70, 72}
	a1_result := []int{-187, -170, -168, -151, -136, -110, -70, -66, -56, -16, 6, 19, 33, 34, 50, 60, 72, 88, 91, 93, 96, 96, 99, 120, 132, 134, 140, 153, 159, 160}
	l1 := NewDoublyLinkedListFromSlice(&a1_input)
	l1.QuickSort()
	if !l1.isSameAs(&a1_result) {
		t.Errorf("TestQuickSort: expected\n list %v\n, got %v", a1_result, *l1.GetSliceValues())
	}
}
