package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertAtTail(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	node := &SinglyLinkedListNode{data, nil}

	if head == nil {
		return node
	}

	tail := head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = node

	return head
}

func main() {
	// Get length of inputs
	var inputs int32
	if _, err := fmt.Scan(&inputs); err != nil {
		panic(err)
	}

	// Make list
	var head *SinglyLinkedListNode
	for range inputs {
		var data int32
		if _, err := fmt.Scan(&data); err != nil {
			panic(err)
		}
		head = insertAtTail(head, data)
	}

	// Print list
	current := head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
