package main

import (
	"fmt"

	"github.com/alex-salamatov/doubly_linked_list/pkg/datastruct"
)

func main() {
	dll := datastruct.MakeDoublyLinkedList()
	fmt.Printf("Doubly Linked List length = %d", dll.Len())
}
