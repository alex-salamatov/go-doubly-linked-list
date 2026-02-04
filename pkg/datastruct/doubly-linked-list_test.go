package datastruct

import (
	"testing"
)

func TestAppend(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.Append(10)
	dll.Append(20)

	if dll.Len() != 2 {
		t.Errorf("Expected size 2, got %d", dll.Len())
	}

	if dll.first.value != 10 || dll.first.next.value != 20 {
		t.Errorf("Values are not appended as expected. Expected list[0] = 10 and list[1] = 20, got %d and %d", dll.first.value, dll.first.next.value)
	}
}
