package main

import (
	"fmt"

	"github.com/alex-salamatov/doubly_linked_list/pkg/datastruct"
)

func PrintDList(dll *datastruct.DoublyLinkedList) {
	fmt.Printf("DList: ")
	for i := 0; i < dll.Len(); i++ {
		val, _ := dll.At(i)
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

func main() {
	dll := datastruct.MakeDoublyLinkedList()
	dll.Append(0)
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)
	dll.Append(5)
	dll.Append(6)

	pos0 := 0
	elem0, _ := dll.At(pos0)

	pos1 := 1
	elem1, _ := dll.At(pos1)

	pos2 := 2
	elem2, _ := dll.At(pos2)

	fmt.Printf("Doubly Linked List length = %d, value at position %d is %d\n", dll.Len(), pos0, elem0)
	fmt.Printf("Doubly Linked List length = %d, value at position %d is %d\n", dll.Len(), pos1, elem1)
	fmt.Printf("Doubly Linked List length = %d, value at position %d is %d\n", dll.Len(), pos2, elem2)

	dll.RemoveAt(4)
	dll.RemoveAt(dll.Len() - 1)
	dll.RemoveAt(0)

	posPL := dll.Len() - 2
	elemPL, _ := dll.At(posPL)

	posL := dll.Len() - 1
	elemL, _ := dll.At(posL)

	fmt.Printf("Doubly Linked List length = %d, value at position %d is %d\n", dll.Len(), posPL, elemPL)
	fmt.Printf("Doubly Linked List length = %d, value at position %d is %d\n", dll.Len(), posL, elemL)

	PrintDList(dll)

	dll.InsertAfter(-1, -1)
	dll.InsertAfter(0, -2)
	dll.InsertAfter(dll.Len()-2, -3)
	dll.InsertAfter(dll.Len()-1, -4)

	PrintDList(dll)
}
