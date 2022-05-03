// Doubly linked list.

package godll

import (
	"fmt"
	"io"
)

// Function used to compare node values.
type fun[T comparable] func(v1, v2 T) bool

type List[T comparable] struct {
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

// Print prints all elements in a List using passed io.Writer interface.
func (l *List[T]) Print(w io.Writer) {
	if l.length == 0 {
		return
	}
	current := l.head
	for current != l.tail {
		fmt.Fprintf(w, "%+v ", current.Value)
		current = current.next
	}
	fmt.Fprintf(w, "%+v", l.tail.Value)
	fmt.Fprintf(w, "\n")
}

func (l *List[T]) validateNegativeIndex(index int) error {
	// Return error if index is negative number.
	if index < 0 {
		return &NegativeIndexError{Index: index}
	}

	return nil
}

func (l *List[T]) validateExistingIndex(index int) error {
	if err := l.validateNegativeIndex(index); err != nil {
		return err
	}

	// Return error if index is larger than or equal to legth of list.
	if index >= l.length {
		return &IndexOutOfRangeError{Index: index}
	}

	return nil
}

func (l *List[T]) validateInsertableIndex(index int) error {
	if err := l.validateNegativeIndex(index); err != nil {
		return err
	}

	// Return error if index is larger than legth of list.
	if index > l.length {
		return &IndexOutOfRangeError{Index: index}
	}

	return nil
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

	// Add current tail to new variable oldTail. Set passed node as a new tail and connect it to old tail with new next and previous links.
	oldTail := l.tail
	l.tail = node
	node.previous = oldTail
	oldTail.next = node
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

	// Add current head to new variable oldHead. Set passed node as a new head and connect it to old head with new next and previous links.
	oldHead := l.head
	l.head = node
	node.next = oldHead
	oldHead.previous = node
	l.length++
}

// InsertAt inserts now node at specific position.
func (l *List[T]) InsertAt(index int, node *Node[T]) error {
	if err := l.validateInsertableIndex(index); err != nil {
		return err
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

// GetByIndex retrieves node by index. Return error if index is out of range. Index of first node is 0.
func (l *List[T]) GetByIndex(index int) (*Node[T], error) {
	if err := l.validateExistingIndex(index); err != nil {
		return nil, err
	}

	// Calculate index in the middle of the list.
	m := l.length / 2

	// If index is closer to head, start iterating through nodes from head.
	if index < m {
		current := l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current, nil
	}

	// If index is closer to tail, start iterating through nodes from tail.
	current := l.tail
	for i := 0; i < (l.length - index - 1); i++ {
		current = current.previous
	}
	return current, nil
}

// GetByValue returns index of node and node with passed value using compare function compFunc.
// If compFunc is nil, use default comparison with "==". Returns -1 and nil if there is no node with given value in List.
func (l *List[T]) GetByValue(value T, compFunc fun[T]) (int, *Node[T]) {
	if compFunc == nil {
		compFunc = func(v1, v2 T) bool { return v1 == v2 }
	}
	current := l.head
	for i := 0; i < l.length; i++ {
		if compFunc(current.Value, value) {
			return i, current
		}
		current = current.next
	}
	return -1, nil
}

// GetAllValues return map with indexes and nodes of all nodes with passed value using compare function compFunc.
// If compFunc is nil, use default comparison with "==". Returns empty map if there is no node with given value in List.
func (l *List[T]) GetAllValues(value T, compFunc fun[T]) map[int]*Node[T] {
	if compFunc == nil {
		compFunc = func(v1, v2 T) bool { return v1 == v2 }
	}
	m := make(map[int]*Node[T], 0)

	current := l.head
	for i := 0; i < l.length; i++ {
		if compFunc(current.Value, value) {
			m[i] = current
		}
		current = current.next
	}

	return m
}

// Swap changes places of nodes on passed positions.
func (l *List[T]) Swap(i, j int) error {
	if err := l.validateExistingIndex(i); err != nil {
		return err
	}

	if err := l.validateExistingIndex(j); err != nil {
		return err
	}

	// Do nothing and return if indexes are the same or length of list is 1.
	if i == j || l.length == 1 {
		return nil
	}

	// If i > j switch them so we have easier corner cases to handle.
	// Now node with index "i" can be only head, it can't be tail.
	// And node with index "j" can only be tail, it can't be head.
	if i > j {
		i, j = j, i
	}

	// Retrieve nodes with passed indexes.
	node1, err := l.GetByIndex(i)
	if err != nil {
		return err
	}
	node2, err := l.GetByIndex(j)
	if err != nil {
		return err
	}

	if j-i == 1 {
		l.swapNeighbours(node1, node2)
		return nil
	}

	l.swap(node1, node2)

	return nil
}

func (l *List[T]) swapNeighbours(node1, node2 *Node[T]) {
	// Define variables where outer neihgbours of both nodes will be saved.
	var node1Prev *Node[T]
	var node2Next *Node[T]

	if node1 != l.head {
		node1Prev = node1.previous
	} else {
		l.head = node2
	}

	if node2 != l.tail {
		node2Next = node2.next
	} else {
		l.tail = node1
	}

	// Connect nodes with new neighoubrs.
	node1.next = node2Next
	node1.previous = node2
	node2.next = node1
	node2.previous = node1Prev

	// Connect old outer neighbours with new nodes.
	if node2 != l.head {
		node1Prev.next = node2
	}

	if node1 != l.tail {
		node2Next.previous = node1
	}
}

func (l *List[T]) swap(node1, node2 *Node[T]) {
	// Define variables where neihgbours of both nodes will be saved.
	var node1Prev *Node[T]
	var node1Next *Node[T]
	var node2Next *Node[T]
	var node2Prev *Node[T]

	if node1 != l.head {
		node1Prev = node1.previous
	} else {
		l.head = node2
	}
	node1Next = node1.next

	node2Prev = node2.previous
	if node2 != l.tail {
		node2Next = node2.next
	} else {
		l.tail = node1
	}

	// Connect nodes with new neighoubrs.
	node1.next = node2Next
	node1.previous = node2Prev
	node2.next = node1Next
	node2.previous = node1Prev

	// Connect old neighbours with new nodes.
	if node2 != l.head {
		node1Prev.next = node2
	}
	node1Next.previous = node2

	if node1 != l.tail {
		node2Next.previous = node1
	}
	node2Prev.next = node1
}

// DeleteAt deletes node at given index.
func (l *List[T]) DeleteAt(index int) error {
	node, err := l.GetByIndex(index)

	if err != nil {
		return err
	}

	// Define variables where outer neihgbours of both nodes will be saved.
	var nodePrev *Node[T]
	var nodeNext *Node[T]

	if node != l.head {
		nodePrev = node.previous
		nodePrev.next = node.next
	} else {
		l.head = node.next
	}

	if node != l.tail {
		nodeNext = node.next
		nodeNext.previous = node.previous
	} else {
		l.tail = node.previous
	}

	l.length--

	return nil
}

// DeleteNode deletes passed ndoe from list.
func (l *List[T]) DeleteNode(node *Node[T]) error {
	if node == nil {
		return nil
	}
	if l.length == 0 {
		return &NodeNotFoundError[T]{Node: node}
	}

	current := l.head
	for current.next != nil || current == l.tail {
		if current == node {
			l.deleteNode(node)
			return nil
		}
		current = current.next
	}

	return &NodeNotFoundError[T]{Node: node}
}

// Delete found node.
func (l *List[T]) deleteNode(node *Node[T]) {
	defer func() { l.length-- }()
	// If node is both head and tail, it is only node in the list.
	if node == l.head && node == l.tail {
		l.head, l.tail = nil, nil
		return
	}

	// If node is head, move head to next node.
	if node == l.head {
		l.head = node.next
		l.head.previous = nil
		return
	}

	// If node is tail, move tail to previous node.
	if node == l.tail {
		l.tail = node.previous
		l.tail.next = nil
		return
	}

	node.next.previous = node.previous
	node.previous.next = node.next
}

// Sort sorts nodes in List using Merge Sort algorithm with sorting function sortFunc.
// If sortFunc is nil, use default comparison with "<".
func (l *List[T]) Sort(sortFunc fun[T]) {
	// Call recursive function to sort list using merge sort algorithm.
	l.head = sort(l.head, sortFunc)

	// Iterate through sorted list to find new tail.
	current := l.head
	for current != nil && current.next != nil {
		current = current.next
	}
	l.tail = current
}

func sort[T comparable](head *Node[T], sortFunc fun[T]) *Node[T] {
	// If list is empty or it contains only one node, return passed head.
	if head == nil || head.next == nil {
		return head
	}

	// Find middle node in list.
	middle := split(head, sortFunc)

	// Sort recursively from head and middle node.
	head = sort(head, sortFunc)
	middle = sort(middle, sortFunc)

	// Call recursive merge function to merge sorted lists.
	return merge(head, middle, sortFunc)
}

func split[T comparable](node *Node[T], sortFunc fun[T]) *Node[T] {
	// Iterate with one node twice the speed to find middle one.
	fast, slow := node, node
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	// Set next of middle node as head of second list.
	middle := slow.next
	// Set middle node as tail for first list.
	slow.next = nil

	return middle
}

func merge[T comparable](node1, node2 *Node[T], sortFunc fun[T]) *Node[T] {
	// If one list is empty, return the other one.
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	// If sorting function returns true exclude head from first list
	// and merge recursively on remainder of first list and second list.
	// Return element excluded from list.
	if sortFunc(node1.Value, node2.Value) {
		node1.next = merge(node1.next, node2, sortFunc)
		node1.next.previous = node1
		node1.previous = nil
		return node1
	}

	// If sorting function returns false exclude head from second list
	// and merge recursively on first list and remainder of second list.
	// Return element excluded from list.
	node2.next = merge(node1, node2.next, sortFunc)
	node2.next.previous = node2
	node2.previous = nil
	return node2
}
