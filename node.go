// Node in double linked list.

package godll

// Node represent node in linked list.
type Node[T comparable] struct {
	Value    T        // Value of node.
	next     *Node[T] // Pointer to next node.
	previous *Node[T] // Pointer to previous node.
}

// Next returns pointer to next Node[T] in list.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Previous returns pointer to previous Node[T] in list.
func (n *Node[T]) Previous() *Node[T] {
	return n.previous
}

// NewNode cretes new node with passed value. Return pointer to newly created node.
func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{Value: value}
}
