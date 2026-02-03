package main

import (
	"fmt"

	"github.com/alex-salamatov/doubly_linked_list/pkg/datastruct"
)

func main() {
	dll := datastruct.MakeDoublyLinkedList()
	dll.Append(99)
	elem, _ := dll.At(0)
	fmt.Printf("Doubly Linked List length = %d, value at position %d is %d", dll.Len(), 0, elem)
}
