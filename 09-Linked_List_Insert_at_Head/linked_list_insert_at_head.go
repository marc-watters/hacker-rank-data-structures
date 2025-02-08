package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	/* UNUSUED */
	//tail *SinglyLinkedListNode
}

func insertNodeAtHead(llist *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	node := &SinglyLinkedListNode{data, nil}
	if llist == nil {
		llist = node
		return llist
	} else {
		node.next = llist
		llist = node
	}
	return llist
}

func printSinglyLinkedList(node *SinglyLinkedListNode) {
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func main() {
	var inputs int32
	if _, err := fmt.Scan(&inputs); err != nil {
		panic(err)
	}

	llist := &SinglyLinkedList{}
	for range inputs {
		var llistItem int32
		if _, err := fmt.Scan(&llistItem); err != nil {
			panic(err)
		}
		llistHead := insertNodeAtHead(llist.head, llistItem)
		llist.head = llistHead
	}

	printSinglyLinkedList(llist.head)
}
