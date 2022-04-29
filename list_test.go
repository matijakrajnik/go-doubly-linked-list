package godll

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)

	assert.Nil(t, list.Head())

	list.Append(node1)
	assert.Equal(t, node1, list.Head())

	list.Prepend(node2)
	assert.Equal(t, node2, list.Head())

	list.Append(node3)
	assert.Equal(t, node2, list.Head())
}

func TestLength(t *testing.T) {
	list := &List[int]{}

	assert.Equal(t, 0, list.Length())

	list.Append(NewNode(8))
	assert.Equal(t, 1, list.Length())

	list.Prepend(NewNode(2))
	assert.Equal(t, 2, list.Length())

	list.Append(NewNode(5))
	assert.Equal(t, 3, list.Length())

	list.InsertAt(0, NewNode(145))
	assert.Equal(t, 4, list.Length())

	list.InsertAt(2, NewNode(65))
	assert.Equal(t, 5, list.Length())

	list.InsertAt(5, NewNode(9312))
	assert.Equal(t, 6, list.Length())
}

func TestPrint(t *testing.T) {
	list := &List[int]{}
	var output bytes.Buffer

	list.Print(&output)
	assert.Equal(t, "", output.String())

	output.Reset()
	list.Append(NewNode(4))
	list.Print(&output)
	assert.Equal(t, "4 \n", output.String())

	output.Reset()
	list.Append(NewNode(23))
	list.Print(&output)
	assert.Equal(t, "4 23 \n", output.String())

	output.Reset()
	list.Append(NewNode(1))
	list.Print(&output)
	assert.Equal(t, "4 23 1 \n", output.String())
}

func TestAppendInt(t *testing.T) {
	list := &List[int]{}
	list.Append(NewNode(8))
	list.Append(NewNode(2))
	list.Append(NewNode(5))

	assert.Equal(t, 8, list.head.Value)
	assert.Equal(t, 2, list.head.next.Value)
	assert.Equal(t, 5, list.head.next.next.Value)
}

func TestAppendFloat64(t *testing.T) {
	list := &List[float64]{}
	list.Append(NewNode(8.1))
	list.Append(NewNode(2.2))
	list.Append(NewNode(5.3))

	assert.Equal(t, 8.1, list.head.Value)
	assert.Equal(t, 2.2, list.head.next.Value)
	assert.Equal(t, 5.3, list.head.next.next.Value)
}

func TestAppendString(t *testing.T) {
	list := &List[string]{}
	list.Append(NewNode("Bruce"))
	list.Append(NewNode("Wayne"))
	list.Append(NewNode("Batman"))

	assert.Equal(t, "Bruce", list.head.Value)
	assert.Equal(t, "Wayne", list.head.next.Value)
	assert.Equal(t, "Batman", list.head.next.next.Value)
}

func TestAppendStruct(t *testing.T) {
	list := &List[PersonTest]{}
	batman := PersonTest{FirstName: "Bruce", LastName: "Wayne"}
	superman := PersonTest{FirstName: "Clark", LastName: "Kent"}
	list.Append(NewNode(batman))
	list.Append(NewNode(superman))

	assert.Equal(t, batman, list.head.Value)
	assert.Equal(t, batman.FirstName, list.head.Value.FirstName)
	assert.Equal(t, batman.LastName, list.head.Value.LastName)
	assert.Equal(t, superman, list.head.next.Value)
	assert.Equal(t, superman.FirstName, list.head.next.Value.FirstName)
	assert.Equal(t, superman.LastName, list.head.next.Value.LastName)
}

func TestPrependInt(t *testing.T) {
	list := &List[int]{}
	list.Prepend(NewNode(8))
	list.Prepend(NewNode(2))
	list.Prepend(NewNode(5))

	assert.Equal(t, 5, list.head.Value)
	assert.Equal(t, 2, list.head.next.Value)
	assert.Equal(t, 8, list.head.next.next.Value)
}

func TestPrependFloat64(t *testing.T) {
	list := &List[float64]{}
	list.Prepend(NewNode(8.1))
	list.Prepend(NewNode(2.2))
	list.Prepend(NewNode(5.3))

	assert.Equal(t, 5.3, list.head.Value)
	assert.Equal(t, 2.2, list.head.next.Value)
	assert.Equal(t, 8.1, list.head.next.next.Value)
}

func TestPrependString(t *testing.T) {
	list := &List[string]{}
	list.Prepend(NewNode("Bruce"))
	list.Prepend(NewNode("Wayne"))
	list.Prepend(NewNode("Batman"))

	assert.Equal(t, "Batman", list.head.Value)
	assert.Equal(t, "Wayne", list.head.next.Value)
	assert.Equal(t, "Bruce", list.head.next.next.Value)
}

func TestPrependStruct(t *testing.T) {
	list := &List[PersonTest]{}
	batman := PersonTest{FirstName: "Bruce", LastName: "Wayne"}
	superman := PersonTest{FirstName: "Clark", LastName: "Kent"}
	list.Prepend(NewNode(batman))
	list.Prepend(NewNode(superman))

	assert.Equal(t, superman, list.head.Value)
	assert.Equal(t, superman.FirstName, list.head.Value.FirstName)
	assert.Equal(t, superman.LastName, list.head.Value.LastName)

	assert.Equal(t, batman, list.head.next.Value)
	assert.Equal(t, batman.FirstName, list.head.next.Value.FirstName)
	assert.Equal(t, batman.LastName, list.head.next.Value.LastName)
}

func TestInsertAt(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	newNode := NewNode(12)
	err := list.InsertAt(1, newNode)

	assert.Nil(t, err)
	assert.Equal(t, node1, newNode.previous)
	assert.Equal(t, node2, newNode.next)
}

func TestInsertAtBeginning(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	newNode := NewNode(12)
	err := list.InsertAt(0, newNode)

	assert.Nil(t, err)
	assert.Equal(t, newNode, list.head)
	assert.Nil(t, newNode.previous)
	assert.Equal(t, node1, newNode.next)
}

func TestInsertAtEnd(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	newNode := NewNode(12)
	err := list.InsertAt(3, newNode)

	assert.Nil(t, err)
	assert.Equal(t, newNode, list.tail)
	assert.Equal(t, node3, newNode.previous)
	assert.Nil(t, newNode.next)
}

func TestInsertAtEmptyList(t *testing.T) {
	list := &List[int]{}
	node := NewNode(27)
	err := list.InsertAt(0, node)

	assert.Nil(t, err)
	assert.Equal(t, node, list.head)
	assert.Equal(t, node, list.tail)
	assert.Equal(t, 1, list.length)
}

func TestInsertAtOutOfRange(t *testing.T) {
	list := &List[int]{}

	newNode1 := NewNode(12)
	err := list.InsertAt(1, newNode1)
	assert.Equal(t, IndexOutOfRangeError, err)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
	assert.Equal(t, 0, list.length)

	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	newNode2 := NewNode(12)
	err = list.InsertAt(4, newNode2)
	assert.Equal(t, IndexOutOfRangeError, err)
	assert.Equal(t, node1, list.head)
	assert.Equal(t, node3, list.tail)
	assert.Equal(t, 3, list.length)
}

func TestGetByIndex(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)
	node4 := NewNode(4)
	node5 := NewNode(9)
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)
	list.Append(node4)
	list.Append(node5)

	retrieved, err := list.GetByIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, node1, retrieved)

	retrieved, err = list.GetByIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, node2, retrieved)

	retrieved, err = list.GetByIndex(2)
	assert.Nil(t, err)
	assert.Equal(t, node3, retrieved)

	retrieved, err = list.GetByIndex(3)
	assert.Nil(t, err)
	assert.Equal(t, node4, retrieved)

	retrieved, err = list.GetByIndex(4)
	assert.Nil(t, err)
	assert.Equal(t, node5, retrieved)
}

func TestGetByIndexOutRange(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)

	retrieved, err := list.GetByIndex(0)
	assert.Equal(t, IndexOutOfRangeError, err)
	assert.Nil(t, retrieved)

	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	retrieved, err = list.GetByIndex(3)
	assert.Equal(t, IndexOutOfRangeError, err)
	assert.Nil(t, retrieved)
}
