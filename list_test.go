package godll

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	list := &List[int]{}
	nodes := testNodesInt(4)

	assert.Nil(t, list.Head())

	list.Append(nodes[0])
	assert.Equal(t, nodes[0], list.Head())

	list.Prepend(nodes[1])
	assert.Equal(t, nodes[1], list.Head())

	list.Append(nodes[2])
	assert.Equal(t, nodes[1], list.Head())

	err := list.Swap(0, 2)
	assert.Nil(t, err)
	assert.Equal(t, nodes[2], list.Head())

	assert.Nil(t, err)
	err = list.InsertAt(0, nodes[3])
	assert.Nil(t, err)
	assert.Equal(t, nodes[3], list.Head())
}

func TestTail(t *testing.T) {
	list := &List[int]{}
	nodes := testNodesInt(4)

	assert.Nil(t, list.Tail())

	list.Append(nodes[0])
	assert.Equal(t, nodes[0], list.Tail())

	list.Prepend(nodes[1])
	assert.Equal(t, nodes[0], list.Tail())

	list.Append(nodes[2])
	assert.Equal(t, nodes[2], list.Tail())

	err := list.Swap(0, 2)
	assert.Nil(t, err)
	assert.Equal(t, nodes[1], list.Tail())

	assert.Nil(t, err)
	err = list.InsertAt(3, nodes[3])
	assert.Nil(t, err)
	assert.Equal(t, nodes[3], list.Tail())
}

func TestLength(t *testing.T) {
	list := &List[int]{}
	nodes := testNodesInt(6)

	assert.Equal(t, 0, list.Length())

	list.Append(nodes[0])
	assert.Equal(t, 1, list.Length())

	list.Prepend(nodes[1])
	assert.Equal(t, 2, list.Length())

	list.Append(nodes[2])
	assert.Equal(t, 3, list.Length())

	list.InsertAt(0, nodes[3])
	assert.Equal(t, 4, list.Length())

	list.InsertAt(2, nodes[4])
	assert.Equal(t, 5, list.Length())

	list.InsertAt(5, nodes[5])
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
	nodes := testNodesInt(5)

	for _, node := range nodes {
		list.Append(node)
		assert.Equal(t, node, list.tail)
		assert.Equal(t, node.Value, list.tail.Value)
	}
}

func TestAppendFloat64(t *testing.T) {
	list := &List[float64]{}
	nodes := testNodesFloat64(5)

	for _, node := range nodes {
		list.Append(node)
		assert.Equal(t, node, list.tail)
		assert.Equal(t, node.Value, list.tail.Value)
	}
}

func TestAppendString(t *testing.T) {
	list := &List[string]{}
	nodes := testNodesString(5)

	for _, node := range nodes {
		list.Append(node)
		assert.Equal(t, node, list.tail)
		assert.Equal(t, node.Value, list.tail.Value)
	}
}

func TestAppendStruct(t *testing.T) {
	list := &List[PersonTest]{}
	nodes := testNodesStruct(5)

	for _, node := range nodes {
		list.Append(node)
		assert.Equal(t, node, list.tail)
		assert.Equal(t, node.Value, list.tail.Value)
		assert.Equal(t, node.Value.FirstName, list.tail.Value.FirstName)
		assert.Equal(t, node.Value.LastName, list.tail.Value.LastName)
	}
}

func TestPrependInt(t *testing.T) {
	list := &List[int]{}
	nodes := testNodesInt(5)

	for _, node := range nodes {
		list.Prepend(node)
		assert.Equal(t, node, list.head)
		assert.Equal(t, node.Value, list.head.Value)
	}
}

func TestPrependFloat64(t *testing.T) {
	list := &List[float64]{}
	nodes := testNodesFloat64(5)

	for _, node := range nodes {
		list.Prepend(node)
		assert.Equal(t, node, list.head)
		assert.Equal(t, node.Value, list.head.Value)
	}
}

func TestPrependString(t *testing.T) {
	list := &List[string]{}
	nodes := testNodesString(5)

	for _, node := range nodes {
		list.Prepend(node)
		assert.Equal(t, node, list.head)
		assert.Equal(t, node.Value, list.head.Value)
	}
}

func TestPrependStruct(t *testing.T) {
	list := &List[PersonTest]{}
	nodes := testNodesStruct(5)

	for _, node := range nodes {
		list.Prepend(node)
		assert.Equal(t, node, list.head)
		assert.Equal(t, node.Value, list.head.Value)
		assert.Equal(t, node.Value.FirstName, list.head.Value.FirstName)
		assert.Equal(t, node.Value.LastName, list.head.Value.LastName)
	}
}

func TestInsertAt(t *testing.T) {
	list, nodes := testListInt(3)

	newNode := NewNode(12)
	err := list.InsertAt(1, newNode)

	assert.Nil(t, err)
	assert.Equal(t, nodes[0], newNode.previous)
	assert.Equal(t, nodes[1], newNode.next)
}

func TestInsertAtBeginning(t *testing.T) {
	list, nodes := testListInt(3)

	newNode := NewNode(12)
	err := list.InsertAt(0, newNode)

	assert.Nil(t, err)
	assert.Equal(t, newNode, list.head)
	assert.Nil(t, newNode.previous)
	assert.Equal(t, nodes[0], newNode.next)
}

func TestInsertAtEnd(t *testing.T) {
	list, nodes := testListInt(3)

	newNode := NewNode(12)
	err := list.InsertAt(3, newNode)

	assert.Nil(t, err)
	assert.Equal(t, newNode, list.tail)
	assert.Equal(t, nodes[2], newNode.previous)
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
	assert.Equal(t, &IndexOutOfRangeError{Index: 1}, err)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
	assert.Equal(t, 0, list.length)

	list, nodes := testListInt(3)

	newNode2 := NewNode(12)
	err = list.InsertAt(4, newNode2)
	assert.Equal(t, &IndexOutOfRangeError{Index: 4}, err)
	assert.Equal(t, nodes[0], list.head)
	assert.Equal(t, nodes[2], list.tail)
	assert.Equal(t, 3, list.length)
}

func TestInsertNegativeIndex(t *testing.T) {
	list, _ := testListInt(3)

	err := list.InsertAt(-1, NewNode(5))
	assert.Equal(t, &NegativeIndexError{Index: -1}, err)
	assert.Equal(t, 3, list.length)
}

func TestGetByIndex(t *testing.T) {
	list, nodes := testListInt(5)

	for i, node := range nodes {
		retrieved, err := list.GetByIndex(i)
		assert.Nil(t, err)
		assert.Equal(t, node, retrieved)
	}
}

func TestGetByIndexOutRange(t *testing.T) {
	list, _ := testListInt(0)

	retrieved, err := list.GetByIndex(0)
	assert.Equal(t, &IndexOutOfRangeError{Index: 0}, err)
	assert.Nil(t, retrieved)

	list, _ = testListInt(3)

	retrieved, err = list.GetByIndex(3)
	assert.Equal(t, &IndexOutOfRangeError{Index: 3}, err)
	assert.Nil(t, retrieved)
}

func TestGetByNegativeIndex(t *testing.T) {
	list, _ := testListInt(3)

	retrieved, err := list.GetByIndex(-1)
	assert.Equal(t, &NegativeIndexError{Index: -1}, err)
	assert.Nil(t, retrieved)
}

func TestGetByValueInt(t *testing.T) {
	list, nodes := testListInt(5)

	for i, node := range nodes {
		index := list.GetByValue(node.Value)
		assert.Equal(t, i, index)
	}
}

func TestGetByValueNotfoundInt(t *testing.T) {
	list := &List[int]{}
	index := list.GetByValue(123)
	assert.Equal(t, -1, index)

	list, _ = testListInt(5)
	index = list.GetByValue(123)
	assert.Equal(t, -1, index)
}

func TestGetByValueFloat64(t *testing.T) {
	list, nodes := testListFloat64(5)

	for i, node := range nodes {
		index := list.GetByValue(node.Value)
		assert.Equal(t, i, index)
	}
}

func TestGetByValueNotfoundFloat64(t *testing.T) {
	list := &List[float64]{}
	index := list.GetByValue(1.23)
	assert.Equal(t, -1, index)

	list, _ = testListFloat64(5)
	index = list.GetByValue(1.23)
	assert.Equal(t, -1, index)
}

func TestGetByValueString(t *testing.T) {
	list, nodes := testListString(5)

	for i, node := range nodes {
		index := list.GetByValue(node.Value)
		assert.Equal(t, i, index)
	}
}

func TestGetByValueNotfoundString(t *testing.T) {
	list := &List[string]{}
	index := list.GetByValue("NOT FOUND")
	assert.Equal(t, -1, index)

	list, _ = testListString(5)
	index = list.GetByValue("NOT FOUND")
	assert.Equal(t, -1, index)
}

func TestGetByValueStruct(t *testing.T) {
	list, nodes := testListStruct(5)

	for i, node := range nodes {
		index := list.GetByValue(node.Value)
		assert.Equal(t, i, index)
	}

	list = &List[PersonTest]{}
	p1 := PersonTest{FirstName: "Bruce", LastName: "Wayne"}
	var p2 PersonTest

	list.Append(NewNode(p1))
	i := list.GetByValue(p2)
	assert.Equal(t, -1, i)
}

func TestGetByValueNotfoundStruct(t *testing.T) {
	person := PersonTest{FirstName: "Bruce", LastName: "Wayne"}

	list := &List[PersonTest]{}
	index := list.GetByValue(person)
	assert.Equal(t, -1, index)

	list, _ = testListStruct(5)
	index = list.GetByValue(person)
	assert.Equal(t, -1, index)

	var empty PersonTest
	index = list.GetByValue(empty)
	assert.Equal(t, -1, index)
}

func TestSwap(t *testing.T) {
	list, nodes := testListInt(5)

	// Test all possible combinations for list with 5 nodes.
	indexCombinations := [][]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
		{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
		{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4},
		{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
	}

	for _, testSet := range indexCombinations {
		// Swap places once.
		err := list.Swap(testSet[0], testSet[1])
		assert.Nil(t, err)
		retrieved, err := list.GetByIndex(testSet[0])
		assert.Nil(t, err)
		assert.Equal(t, nodes[testSet[1]], retrieved)
		retrieved, err = list.GetByIndex(testSet[1])
		assert.Nil(t, err)
		assert.Equal(t, nodes[testSet[0]], retrieved)

		// Swap places of same elements again to reset to previous positions.
		err = list.Swap(testSet[1], testSet[0])
		assert.Nil(t, err)
		retrieved, err = list.GetByIndex(testSet[0])
		assert.Nil(t, err)
		assert.Equal(t, nodes[testSet[0]], retrieved)
		retrieved, err = list.GetByIndex(testSet[1])
		assert.Nil(t, err)
		assert.Equal(t, nodes[testSet[1]], retrieved)
	}
}

func TestSwapOutOfRange(t *testing.T) {
	list, _ := testListInt(5)

	err := list.Swap(1, 5)
	assert.Equal(t, &IndexOutOfRangeError{Index: 5}, err)

	err = list.Swap(5, 1)
	assert.Equal(t, &IndexOutOfRangeError{Index: 5}, err)

	err = list.Swap(5, 6)
	assert.Equal(t, &IndexOutOfRangeError{Index: 5}, err)

	err = list.Swap(6, 5)
	assert.Equal(t, &IndexOutOfRangeError{Index: 6}, err)
}

func TestSwapNegativeIndex(t *testing.T) {
	list, _ := testListInt(5)

	err := list.Swap(-1, 2)
	assert.Equal(t, &NegativeIndexError{Index: -1}, err)

	err = list.Swap(3, -1)
	assert.Equal(t, &NegativeIndexError{Index: -1}, err)

}
