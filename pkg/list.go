// Double linked list.

package godll

import "fmt"

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
