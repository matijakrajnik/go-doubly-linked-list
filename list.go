// Double linked list.

package godll

import (
	"errors"
	"fmt"
)

var (
	IndexOutOfRangeError = errors.New("Index is out of range!")
)

type List[T any] struct {
	head   *Node[T] // Pointer to head (first node in list).
	tail   *Node[T] // Pointer to tail (last node in list).
	length int      // Number of nodes in list.
}

// Head returns first node in list.
func (l *List[T]) Head() *Node[T] {
	return l.head
}

// Tail returns last node in list.
func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

// Length returns number of List[T] length field.
func (l *List[T]) Length() int {
	return l.length
}

// Append adds node to the end of the List.
func (l *List[T]) Append(node *Node[T]) {
	// Set node as a head and tail if list is empty.
	if l.length == 0 {
		l.head = node
		l.tail = node
		l.length++
		return
	}

	// If list not empty, find last node and set passed node as next node for current last one.
	last := l.head
	for i := 1; i < l.length; i++ {
		last = last.next
	}
	last.next = node
	node.previous = last
	l.tail = node
	l.length++
}

// Prepend adds node to the beggining of the List.
func (l *List[T]) Prepend(node *Node[T]) {
	// Set node as a head and tail if list is empty.
	if l.length == 0 {
		l.head = node
		l.tail = node
		l.length++
		return
	}

	// Add current head to new variable oldHead. Set passed node as a new head. In new head set next node to be old head.
	oldHead := l.head
	l.head = node
	node.next = oldHead
	oldHead.previous = node
	l.length++
}

func (l *List[T]) InsertAt(index int, node *Node[T]) error {
	// Return error if index is larger than legth of list.
	if index > l.length {
		return IndexOutOfRangeError
	}

	// If index is 0 set node as new head of list and connect it to neighbours with new next and previous links.
	// If length is 0, set node as tail also.
	if index == 0 {
		node.next = l.head
		if l.length != 0 {
			l.head.previous = node
		}
		l.head = node
		if l.length == 0 {
			l.tail = node
		}
		l.length++
		return nil
	}

	// If index is equal to length of list set node as tail of a list and connect it to neighbours with new next and previous links.
	if index == l.length {
		node.previous = l.tail
		l.tail.next = node
		l.tail = node
		l.length++
		return nil
	}

	// Traverse through list to find index, insert new node and connect it to neighbours with new next and previous links.
	current := l.head
	for i := 1; i < index; i++ {
		current = current.next
	}
	next := current.next
	current.next = node
	node.next = next
	node.previous = current
	next.previous = node
	l.length++
	return nil
}

// Print prints all elements in a List.
func (l *List[T]) Print() {
	if l.length == 0 {
		return
	}
	current := l.head
	for i := 1; i <= l.length; i++ {
		fmt.Printf("%v ", current.Value)
		current = current.next
	}
	fmt.Printf("\n")
}
