package src

import (
	"errors"
	"fmt"
	"sync"
)

type Node struct {
	Data any
	Next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	mu   sync.RWMutex // Lock for thread safety
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Insert inserts a new node with the given data at the end of the list
func (ll *LinkedList) Insert(data any) {
	ll.mu.Lock()         // Acquire exclusive lock for writing
	defer ll.mu.Unlock() // Ensure unlocking at the end of the function

	newNode := &Node{Data: data}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.Next = newNode
		ll.tail = newNode
	}
}

// Delete removes the first node with the given data from the list
func (ll *LinkedList) Delete(data any) error {
	ll.mu.Lock()         // Acquire exclusive lock for writing
	defer ll.mu.Unlock() // Ensure unlocking at the end of the function

	if ll.head == nil {
		return errors.New("list is empty")
	}

	// If the head is the node to be deleted
	if ll.head.Data == data {
		ll.head = ll.head.Next
		if ll.head == nil {
			ll.tail = nil // If the list becomes empty, set tail to nil
		}
		return nil
	}

	// Traverse the list to find the node to delete
	temp := ll.head
	for temp.Next != nil {
		if temp.Next.Data == data {
			// Remove the node
			temp.Next = temp.Next.Next
			if temp.Next == nil {
				ll.tail = temp // If the deleted node was the tail, update tail
			}
			return nil
		}
		temp = temp.Next
	}

	return errors.New("node not found") // If data not found
}

// Display prints all the elements of the linked list
func (ll *LinkedList) Display() {
	ll.mu.RLock()         // Acquire read lock for reading
	defer ll.mu.RUnlock() // Ensure unlocking at the end of the function

	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	temp := ll.head
	for temp != nil {
		fmt.Printf("%v -> ", temp.Data)
		temp = temp.Next
	}
	fmt.Println("NULL")
}

func LinkedListRun() {
	// Create a new LinkedList
	ll := NewLinkedList()

	// Insert data into the list
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	// Display the list
	ll.Display()

	// Delete an item
	if err := ll.Delete(20); err != nil {
		fmt.Println("Error:", err)
	} else {
		ll.Display()
	}

	// Delete the head node
	if err := ll.Delete(10); err != nil {
		fmt.Println("Error:", err)
	} else {
		ll.Display()
	}

	// Try to delete a non-existent node
	if err := ll.Delete(100); err != nil {
		fmt.Println("Error:", err)
	} else {
		ll.Display()
	}
}
