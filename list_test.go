package godll

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		list := &List[int]{}
		assert.Nil(t, list.Head())
	})

	t.Run("After Append/Prepend", func(t *testing.T) {
		list, nodes := &List[int]{}, testNodesInt(3)

		list.Append(nodes[0])
		assert.Equal(t, nodes[0], list.Head())
		list.Prepend(nodes[1])
		assert.Equal(t, nodes[1], list.Head())
		list.Append(nodes[2])
		assert.Equal(t, nodes[1], list.Head())
	})

	t.Run("After Swap", func(t *testing.T) {
		list, nodes := testListInt(4)

		t.Run("Head changed", func(t *testing.T) {
			err := list.Swap(0, 2)
			assert.Nil(t, err)
			assert.Equal(t, nodes[2], list.Head())
		})

		t.Run("Head unchanged", func(t *testing.T) {
			err := list.Swap(1, 3)
			assert.Nil(t, err)
			assert.Equal(t, nodes[2], list.Head())
		})
	})

	t.Run("After Insert", func(t *testing.T) {
		list, _ := testListInt(4)
		node1 := NewNode(123)
		node2 := NewNode(456)

		t.Run("Head changed", func(t *testing.T) {
			err := list.InsertAt(0, node1)
			assert.Nil(t, err)
			assert.Equal(t, node1, list.Head())
		})

		t.Run("Head unchanged", func(t *testing.T) {
			err := list.InsertAt(1, node2)
			assert.Nil(t, err)
			assert.Equal(t, node1, list.Head())
		})
	})

	t.Run("After Delete", func(t *testing.T) {
		t.Run("DeleteAt head changed", func(t *testing.T) {
			list, nodes := testListInt(4)
			err := list.DeleteAt(0)
			assert.Nil(t, err)
			assert.Equal(t, nodes[1], list.Head())
		})

		t.Run("DeleteAt head unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)
			err := list.DeleteAt(1)
			assert.Nil(t, err)
			assert.Equal(t, nodes[0], list.Head())
		})

		t.Run("DeleteNode head changed", func(t *testing.T) {
			list, nodes := testListInt(4)

			err := list.DeleteNode(nodes[0])
			assert.Nil(t, err)
			assert.Equal(t, nodes[1], list.Head())
		})

		t.Run("DeleteNode head unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)

			err := list.DeleteNode(nodes[1])
			assert.Nil(t, err)
			assert.Equal(t, nodes[0], list.Head())
		})

		t.Run("DeleteValues head changed", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Prepend(NewNode(nodes[0].Value))

			deleted := list.DeleteValues(nodes[0].Value, nil)
			assert.Equal(t, 2, deleted)
			assert.Equal(t, nodes[1], list.Head())
		})

		t.Run("DeleteValues head unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Append(NewNode(nodes[1].Value))

			deleted := list.DeleteValues(nodes[1].Value, nil)
			assert.Equal(t, 2, deleted)
			assert.Equal(t, nodes[0], list.Head())
		})
	})
	t.Run("After Sort", func(t *testing.T) {
		t.Run("Head changed", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Sort(func(v1, v2 int) bool { return v1 > v2 })
			assert.Equal(t, nodes[3], list.Head())
		})

		t.Run("Head unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Sort(func(v1, v2 int) bool { return v1 < v2 })
			assert.Equal(t, nodes[0], list.Head())
		})
	})
}

func TestTail(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		list := &List[int]{}
		assert.Nil(t, list.Tail())
	})

	t.Run("After Append/Prepend", func(t *testing.T) {
		list, nodes := &List[int]{}, testNodesInt(3)

		list.Append(nodes[0])
		assert.Equal(t, nodes[0], list.Tail())

		list.Prepend(nodes[1])
		assert.Equal(t, nodes[0], list.Tail())

		list.Append(nodes[2])
		assert.Equal(t, nodes[2], list.Tail())
	})

	t.Run("After Swap", func(t *testing.T) {
		list, nodes := testListInt(4)

		t.Run("Tail changed", func(t *testing.T) {
			err := list.Swap(1, 3)
			assert.Nil(t, err)
			assert.Equal(t, nodes[1], list.Tail())
		})

		t.Run("Tail unchanged", func(t *testing.T) {
			err := list.Swap(0, 2)
			assert.Nil(t, err)
			assert.Equal(t, nodes[1], list.Tail())
		})
	})

	t.Run("After Insert", func(t *testing.T) {
		list, _ := testListInt(4)
		node1 := NewNode(123)
		node2 := NewNode(456)

		t.Run("Tail changed", func(t *testing.T) {
			err := list.InsertAt(4, node1)
			assert.Nil(t, err)
			assert.Equal(t, node1, list.Tail())
		})

		t.Run("Tail unchanged", func(t *testing.T) {
			err := list.InsertAt(4, node2)
			assert.Nil(t, err)
			assert.Equal(t, node1, list.Tail())
		})
	})

	t.Run("After Delete", func(t *testing.T) {
		t.Run("DeleteAt tail changed", func(t *testing.T) {
			list, nodes := testListInt(4)
			err := list.DeleteAt(3)
			assert.Nil(t, err)
			assert.Equal(t, nodes[2], list.Tail())
		})

		t.Run("DeleteAt tail unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)
			err := list.DeleteAt(2)
			assert.Nil(t, err)
			assert.Equal(t, nodes[3], list.Tail())
		})

		t.Run("DeleteNode tail changed", func(t *testing.T) {
			list, nodes := testListInt(4)

			err := list.DeleteNode(nodes[3])
			assert.Nil(t, err)
			assert.Equal(t, nodes[2], list.Tail())
		})

		t.Run("DeleteNode tail unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)

			err := list.DeleteNode(nodes[2])
			assert.Nil(t, err)
			assert.Equal(t, nodes[3], list.Tail())
		})

		t.Run("DeleteValues tail changed", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Append(NewNode(nodes[3].Value))

			deleted := list.DeleteValues(nodes[3].Value, nil)
			assert.Equal(t, 2, deleted)
			assert.Equal(t, nodes[2], list.Tail())
		})

		t.Run("DeleteValues tail unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Prepend(NewNode(nodes[2].Value))

			deleted := list.DeleteValues(nodes[2].Value, nil)
			assert.Equal(t, 2, deleted)
			assert.Equal(t, nodes[3], list.Tail())
		})
	})
	t.Run("After Sort", func(t *testing.T) {
		t.Run("Tail changed", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Sort(func(v1, v2 int) bool { return v1 > v2 })
			assert.Equal(t, nodes[0], list.Tail())
		})

		t.Run("Tail unchanged", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Sort(func(v1, v2 int) bool { return v1 < v2 })
			assert.Equal(t, nodes[3], list.Tail())
		})
	})
}

func TestLength(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		list := &List[int]{}
		assert.Equal(t, 0, list.Length())
	})

	t.Run("After Append/Prepend", func(t *testing.T) {
		list, nodes := &List[int]{}, testNodesInt(3)

		list.Append(nodes[0])
		assert.Equal(t, 1, list.Length())

		list.Prepend(nodes[1])
		assert.Equal(t, 2, list.Length())
	})

	t.Run("After Swap", func(t *testing.T) {
		list, _ := testListInt(4)

		err := list.Swap(1, 3)
		assert.Nil(t, err)
		assert.Equal(t, 4, list.Length())
	})

	t.Run("After Insert", func(t *testing.T) {
		list, _ := testListInt(4)

		err := list.InsertAt(2, NewNode(123))
		assert.Nil(t, err)
		assert.Equal(t, 5, list.Length())
	})

	t.Run("After Delete", func(t *testing.T) {
		t.Run("DeleteAt", func(t *testing.T) {
			list, _ := testListInt(4)
			err := list.DeleteAt(1)
			assert.Nil(t, err)
			assert.Equal(t, 3, list.Length())
		})

		t.Run("DeleteNode", func(t *testing.T) {
			list, nodes := testListInt(4)

			err := list.DeleteNode(nodes[0])
			assert.Nil(t, err)
			assert.Equal(t, 3, list.Length())
		})

		t.Run("DeleteValues", func(t *testing.T) {
			list, nodes := testListInt(4)
			list.Append(NewNode(nodes[3].Value))

			deleted := list.DeleteValues(nodes[3].Value, nil)
			assert.Equal(t, 2, deleted)
			assert.Equal(t, 3, list.Length())
		})
	})
	t.Run("After Sort", func(t *testing.T) {
		t.Run("Tail changed", func(t *testing.T) {
			list, _ := testListInt(4)
			list.Sort(func(v1, v2 int) bool { return v1 > v2 })
			assert.Equal(t, 4, list.Length())
		})
	})
}

func TestPrint(t *testing.T) {
	list := &List[int]{}
	var output bytes.Buffer

	t.Run("Empty", func(t *testing.T) {
		list.Print(&output)
		assert.Equal(t, "", output.String())
	})

	t.Run("One node", func(t *testing.T) {
		output.Reset()
		list.Append(NewNode(4))
		list.Print(&output)
		assert.Equal(t, "4\n", output.String())
	})

	t.Run("Two nodes", func(t *testing.T) {
		output.Reset()
		list.Append(NewNode(23))
		list.Print(&output)
		assert.Equal(t, "4 23\n", output.String())
	})

	t.Run("Three nodes", func(t *testing.T) {
		output.Reset()
		list.Append(NewNode(1))
		list.Print(&output)
		assert.Equal(t, "4 23 1\n", output.String())
	})
}

func TestAppend(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		list, nodes := &List[int]{}, testNodesInt(5)
		for _, node := range nodes {
			list.Append(node)
			assert.Equal(t, node, list.tail)
			assert.Equal(t, node.Value, list.tail.Value)
		}
	})

	t.Run("Int", func(t *testing.T) {
		list, nodes := &List[float64]{}, testNodesFloat64(5)
		for _, node := range nodes {
			list.Append(node)
			assert.Equal(t, node, list.tail)
			assert.Equal(t, node.Value, list.tail.Value)
		}
	})

	t.Run("String", func(t *testing.T) {
		list, nodes := &List[string]{}, testNodesString(5)
		for _, node := range nodes {
			list.Append(node)
			assert.Equal(t, node, list.tail)
			assert.Equal(t, node.Value, list.tail.Value)
		}
	})

	t.Run("Struct", func(t *testing.T) {
		list, nodes := &List[PersonTest]{}, testNodesStruct(5)
		for _, node := range nodes {
			list.Append(node)
			assert.Equal(t, node, list.tail)
			assert.Equal(t, node.Value, list.tail.Value)
			assert.Equal(t, node.Value.FirstName, list.tail.Value.FirstName)
			assert.Equal(t, node.Value.LastName, list.tail.Value.LastName)
		}
	})
}

func TestPrepend(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		list, nodes := &List[int]{}, testNodesInt(5)
		for _, node := range nodes {
			list.Prepend(node)
			assert.Equal(t, node, list.head)
			assert.Equal(t, node.Value, list.head.Value)
		}
	})

	t.Run("Float64", func(t *testing.T) {
		list, nodes := &List[float64]{}, testNodesFloat64(5)
		for _, node := range nodes {
			list.Prepend(node)
			assert.Equal(t, node, list.head)
			assert.Equal(t, node.Value, list.head.Value)
		}
	})

	t.Run("String", func(t *testing.T) {
		list, nodes := &List[string]{}, testNodesString(5)
		for _, node := range nodes {
			list.Prepend(node)
			assert.Equal(t, node, list.head)
			assert.Equal(t, node.Value, list.head.Value)
		}
	})

	t.Run("Struct", func(t *testing.T) {
		list, nodes := &List[PersonTest]{}, testNodesStruct(5)
		for _, node := range nodes {
			list.Prepend(node)
			assert.Equal(t, node, list.head)
			assert.Equal(t, node.Value, list.head.Value)
			assert.Equal(t, node.Value.FirstName, list.head.Value.FirstName)
			assert.Equal(t, node.Value.LastName, list.head.Value.LastName)
		}
	})
}

func TestInsertAt(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		list, node := &List[int]{}, NewNode(27)
		err := list.InsertAt(0, node)
		assert.Nil(t, err)
		assert.Equal(t, node, list.head)
		assert.Equal(t, node, list.tail)
	})

	t.Run("Beginning", func(t *testing.T) {
		list, nodes := testListInt(3)
		newNode := NewNode(12)
		err := list.InsertAt(0, newNode)
		assert.Nil(t, err)
		assert.Equal(t, newNode, list.head)
		assert.Nil(t, newNode.previous)
		assert.Equal(t, nodes[0], newNode.next)
	})

	t.Run("Middle", func(t *testing.T) {
		list, nodes := testListInt(3)
		newNode := NewNode(12)
		err := list.InsertAt(1, newNode)
		assert.Nil(t, err)
		assert.Equal(t, nodes[0], newNode.previous)
		assert.Equal(t, nodes[1], newNode.next)
	})

	t.Run("End", func(t *testing.T) {
		list, nodes := testListInt(3)
		newNode := NewNode(12)
		err := list.InsertAt(3, newNode)
		assert.Nil(t, err)
		assert.Equal(t, newNode, list.tail)
		assert.Equal(t, nodes[2], newNode.previous)
		assert.Nil(t, newNode.next)
	})

	t.Run("Out of range", func(t *testing.T) {
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
	})

	t.Run("Negative index", func(t *testing.T) {
		list, _ := testListInt(3)
		err := list.InsertAt(-1, NewNode(5))
		assert.Equal(t, &NegativeIndexError{Index: -1}, err)
		assert.Equal(t, 3, list.length)
	})
}

func TestGetByIndex(t *testing.T) {
	t.Run("Existing", func(t *testing.T) {
		list, nodes := testListInt(5)
		for i, node := range nodes {
			retrieved, err := list.GetByIndex(i)
			assert.Nil(t, err)
			assert.Equal(t, node, retrieved)
		}
	})

	t.Run("Out of range", func(t *testing.T) {
		list, _ := testListInt(0)
		retrieved, err := list.GetByIndex(0)
		assert.Equal(t, &IndexOutOfRangeError{Index: 0}, err)
		assert.Nil(t, retrieved)

		list, _ = testListInt(3)
		retrieved, err = list.GetByIndex(3)
		assert.Equal(t, &IndexOutOfRangeError{Index: 3}, err)
		assert.Nil(t, retrieved)
	})

	t.Run("Negative index", func(t *testing.T) {
		list, _ := testListInt(3)
		retrieved, err := list.GetByIndex(-1)
		assert.Equal(t, &NegativeIndexError{Index: -1}, err)
		assert.Nil(t, retrieved)
	})
}

func TestGetByValue(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		t.Run("Found", func(t *testing.T) {
			list, nodes := testListInt(5)
			for i, node := range nodes {
				index, retrieved := list.GetByValue(node.Value, nil)
				assert.Equal(t, node, retrieved)
				assert.Equal(t, i, index)
			}
		})

		t.Run("Not found", func(t *testing.T) {
			list := &List[int]{}
			index, node := list.GetByValue(123, nil)
			assert.Nil(t, node)
			assert.Equal(t, -1, index)

			list, _ = testListInt(5)
			index, node = list.GetByValue(123, nil)
			assert.Nil(t, node)
			assert.Equal(t, -1, index)
		})
	})

	t.Run("Float64", func(t *testing.T) {
		t.Run("Found", func(t *testing.T) {
			list, nodes := testListFloat64(5)
			for i, node := range nodes {
				index, retrieved := list.GetByValue(node.Value, nil)
				assert.Equal(t, node, retrieved)
				assert.Equal(t, i, index)
			}
		})

		t.Run("Not found", func(t *testing.T) {
			list := &List[float64]{}
			index, node := list.GetByValue(1.23, nil)
			assert.Nil(t, node)
			assert.Equal(t, -1, index)

			list, _ = testListFloat64(5)
			index, node = list.GetByValue(1.23, nil)
			assert.Nil(t, node)
			assert.Equal(t, -1, index)
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("Found", func(t *testing.T) {
			list, nodes := testListString(5)
			for i, node := range nodes {
				index, retrieved := list.GetByValue(node.Value, nil)
				assert.Equal(t, node, retrieved)
				assert.Equal(t, i, index)
			}
		})

		t.Run("Not found", func(t *testing.T) {
			list := &List[string]{}
			index, node := list.GetByValue("NOT FOUND", nil)
			assert.Nil(t, node)
			assert.Equal(t, -1, index)

			list, _ = testListString(5)
			index, node = list.GetByValue("NOT FOUND", nil)
			assert.Nil(t, node)
			assert.Equal(t, -1, index)
		})
	})

	t.Run("Strunct", func(t *testing.T) {
		t.Run("Found", func(t *testing.T) {
			list, nodes := testListStruct(5)
			for i, node := range nodes {
				index, retrieved := list.GetByValue(node.Value, nil)
				assert.Equal(t, node, retrieved)
				assert.Equal(t, i, index)
			}
		})

		t.Run("Custom function", func(t *testing.T) {
			list, _ := testListStruct(5)
			i := 2

			p := PersonTest{ID: 123, FirstName: "Clark", LastName: "Kent"}
			node, err := list.GetByIndex(i)
			assert.Nil(t, err)
			node.Value = p

			index, node := list.GetByValue(p, func(v1, v2 PersonTest) bool { return v1.ID == v2.ID })
			assert.Equal(t, i, index)
			assert.Equal(t, p, node.Value)
		})

		t.Run("Not found", func(t *testing.T) {
			person := PersonTest{FirstName: "Bruce", LastName: "Wayne"}

			list := &List[PersonTest]{}
			index, node := list.GetByValue(person, nil)
			assert.Equal(t, -1, index)
			assert.Nil(t, node)

			list, _ = testListStruct(5)
			index, node = list.GetByValue(person, nil)
			assert.Equal(t, -1, index)
			assert.Nil(t, node)

			var empty PersonTest
			index, node = list.GetByValue(empty, nil)
			assert.Equal(t, -1, index)
			assert.Nil(t, node)
		})
	})
}

func TestGetAllValues(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		t.Run("One", func(t *testing.T) {
			list, _ := testListInt(3)
			node := NewNode(123)
			list.Append(node)
			indexes := list.GetAllValues(node.Value, nil)
			assert.Equal(t, map[int]*Node[int]{3: node}, indexes)
		})

		t.Run("Multiple", func(t *testing.T) {
			list, nodes := &List[int]{}, testNodesInt(3)
			value := 123
			node1 := NewNode(value)
			node2 := NewNode(value)
			node3 := NewNode(value)
			list.Append(nodes[0])
			list.Append(node1)
			list.Append(nodes[1])
			list.Append(node2)
			list.Append(nodes[2])
			list.Append(node3)
			indexes := list.GetAllValues(value, nil)
			assert.Equal(t, map[int]*Node[int]{1: node1, 3: node2, 5: node3}, indexes)
		})
	})

	t.Run("Custom function", func(t *testing.T) {
		list, _ := testListStruct(5)
		i, j := 2, 3
		p := PersonTest{ID: 123, FirstName: "Clark", LastName: "Kent"}
		node1, err := list.GetByIndex(i)
		assert.Nil(t, err)
		node2, err := list.GetByIndex(j)
		assert.Nil(t, err)
		node1.Value, node2.Value = p, p
		indexes := list.GetAllValues(p, func(v1, v2 PersonTest) bool { return v1.ID == v2.ID })
		expected := map[int]*Node[PersonTest]{i: node1, j: node2}
		assert.Equal(t, expected, indexes)
	})

	t.Run("Not found", func(t *testing.T) {
		list := &List[int]{}
		nodes := testNodesInt(3)
		value := 123
		indexes := list.GetAllValues(value, nil)
		assert.Empty(t, indexes)
		for _, node := range nodes {
			list.Append(node)
		}
		indexes = list.GetAllValues(value, nil)
		assert.Empty(t, indexes)
	})
}
func TestSwap(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
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
	})

	t.Run("Out of range", func(t *testing.T) {
		list, _ := testListInt(5)
		err := list.Swap(1, 5)
		assert.Equal(t, &IndexOutOfRangeError{Index: 5}, err)
		err = list.Swap(5, 1)
		assert.Equal(t, &IndexOutOfRangeError{Index: 5}, err)
		err = list.Swap(5, 6)
		assert.Equal(t, &IndexOutOfRangeError{Index: 5}, err)
		err = list.Swap(6, 5)
		assert.Equal(t, &IndexOutOfRangeError{Index: 6}, err)
	})

	t.Run("Negative index", func(t *testing.T) {
		list, _ := testListInt(5)
		err := list.Swap(-1, 2)
		assert.Equal(t, &NegativeIndexError{Index: -1}, err)
		err = list.Swap(3, -1)
		assert.Equal(t, &NegativeIndexError{Index: -1}, err)
	})
}

func TestDeleteAt(t *testing.T) {
	t.Run("In the middle", func(t *testing.T) {
		list, nodes := testListInt(5)
		err := list.DeleteAt(2)
		assert.Nil(t, err)
		retrieved, err := list.GetByIndex(0)
		assert.Nil(t, err)
		assert.Equal(t, nodes[0], retrieved)
		retrieved, err = list.GetByIndex(1)
		assert.Nil(t, err)
		assert.Equal(t, nodes[1], retrieved)
		retrieved, err = list.GetByIndex(2)
		assert.Nil(t, err)
		assert.Equal(t, nodes[3], retrieved)
		retrieved, err = list.GetByIndex(3)
		assert.Nil(t, err)
		assert.Equal(t, nodes[4], retrieved)
	})

	t.Run("Next to tail", func(t *testing.T) {
		list, nodes := testListInt(4)
		err := list.DeleteAt(2)
		assert.Nil(t, err)
		retrieved, err := list.GetByIndex(0)
		assert.Nil(t, err)
		assert.Equal(t, nodes[0], retrieved)
		retrieved, err = list.GetByIndex(1)
		assert.Nil(t, err)
		assert.Equal(t, nodes[1], retrieved)
		retrieved, err = list.GetByIndex(2)
		assert.Nil(t, err)
		assert.Equal(t, nodes[3], retrieved)
	})

	t.Run("Next to head", func(t *testing.T) {
		list, nodes := testListInt(4)
		err := list.DeleteAt(1)
		assert.Nil(t, err)
		retrieved, err := list.GetByIndex(0)
		assert.Nil(t, err)
		assert.Equal(t, nodes[0], retrieved)
		retrieved, err = list.GetByIndex(1)
		assert.Nil(t, err)
		assert.Equal(t, nodes[2], retrieved)
		retrieved, err = list.GetByIndex(2)
		assert.Nil(t, err)
		assert.Equal(t, nodes[3], retrieved)
	})

	t.Run("Between head and tail", func(t *testing.T) {
		list, nodes := testListInt(3)
		err := list.DeleteAt(1)
		assert.Nil(t, err)
		retrieved, err := list.GetByIndex(0)
		assert.Nil(t, err)
		assert.Equal(t, nodes[0], retrieved)
		retrieved, err = list.GetByIndex(1)
		assert.Nil(t, err)
		assert.Equal(t, nodes[2], retrieved)
	})

	t.Run("Last node", func(t *testing.T) {
		list, _ := testListInt(1)
		err := list.DeleteAt(0)
		assert.Nil(t, err)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("Out of range", func(t *testing.T) {
		list := &List[int]{}
		err := list.DeleteAt(0)
		assert.Equal(t, &IndexOutOfRangeError{Index: 0}, err)

		list, _ = testListInt(5)
		err = list.DeleteAt(5)
		assert.Equal(t, &IndexOutOfRangeError{Index: 5}, err)
	})
}

func TestDeleteNode(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		list, node := &List[int]{}, &Node[int]{}
		err := list.DeleteNode(node)
		assert.Equal(t, &NodeNotFoundError[int]{Node: node}, err)
	})

	t.Run("Head", func(t *testing.T) {
		list, nodes := testListInt(3)
		err := list.DeleteNode(nodes[0])
		assert.Nil(t, err)
		assert.Equal(t, nodes[1], list.head)
		assert.Equal(t, nodes[2], list.head.next)
		assert.Nil(t, list.head.previous)
	})

	t.Run("Middle", func(t *testing.T) {
		list, nodes := testListInt(3)
		err := list.DeleteNode(nodes[1])
		assert.Nil(t, err)
		assert.Equal(t, 2, list.length)
		assert.Equal(t, nodes[0], list.head)
		assert.Equal(t, nodes[2], list.tail)
		assert.Equal(t, nodes[2], list.head.next)
		assert.Equal(t, nodes[0], list.tail.previous)
		assert.Nil(t, list.head.previous)
		assert.Nil(t, list.tail.next)
	})

	t.Run("Tail", func(t *testing.T) {
		list, nodes := testListInt(3)
		err := list.DeleteNode(nodes[2])
		assert.Nil(t, err)
		assert.Equal(t, nodes[1], list.tail)
		assert.Equal(t, nodes[0], list.tail.previous)
		assert.Nil(t, list.tail.next)
	})

	t.Run("Last", func(t *testing.T) {
		list, nodes := testListInt(1)
		err := list.DeleteNode(nodes[0])
		assert.Nil(t, err)
		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)
	})

	t.Run("Nil", func(t *testing.T) {
		list, _ := testListInt(3)
		err := list.DeleteNode(nil)
		assert.Nil(t, err)
	})

	t.Run("Not found", func(t *testing.T) {
		list, _ := testListInt(3)
		node := NewNode(123)
		err := list.DeleteNode(node)
		assert.Equal(t, &NodeNotFoundError[int]{Node: node}, err)
	})
}

func TestDeleteValues(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		list, node := &List[int]{}, &Node[int]{}
		deleted := list.DeleteValues(node.Value, nil)
		assert.Equal(t, 0, deleted)

	})

	t.Run("Head", func(t *testing.T) {
		list, nodes := testListInt(3)
		deleted := list.DeleteValues(nodes[0].Value, nil)
		assert.Equal(t, 1, deleted)
		assert.Equal(t, nodes[1], list.head)
		assert.Equal(t, nodes[2], list.head.next)
		assert.Nil(t, list.head.previous)
	})

	t.Run("Middle", func(t *testing.T) {
		list, nodes := testListInt(3)
		list.InsertAt(2, NewNode(nodes[1].Value))
		deleted := list.DeleteValues(nodes[1].Value, nil)
		assert.Equal(t, 2, deleted)
		assert.Equal(t, 2, list.length)
		assert.Equal(t, nodes[0], list.head)
		assert.Equal(t, nodes[2], list.tail)
		assert.Equal(t, nodes[2], list.head.next)
		assert.Equal(t, nodes[0], list.tail.previous)
		assert.Nil(t, list.head.previous)
		assert.Nil(t, list.tail.next)
	})

	t.Run("Tail", func(t *testing.T) {
		list, nodes := testListInt(3)
		deleted := list.DeleteValues(nodes[2].Value, nil)
		assert.Equal(t, 1, deleted)
		assert.Equal(t, nodes[1], list.tail)
		assert.Equal(t, nodes[0], list.tail.previous)
		assert.Nil(t, list.tail.next)
	})

	t.Run("Head and tail", func(t *testing.T) {
		list, nodes := testListInt(3)
		list.Append(NewNode(list.head.Value))
		deleted := list.DeleteValues(nodes[0].Value, nil)
		assert.Equal(t, 2, deleted)
		assert.Equal(t, nodes[1], list.head)
		assert.Equal(t, nodes[2], list.head.next)
		assert.Nil(t, list.head.previous)
		assert.Equal(t, nodes[2], list.tail)
		assert.Equal(t, nodes[1], list.tail.previous)
		assert.Nil(t, list.tail.next)
	})

	t.Run("Last", func(t *testing.T) {
		list, nodes := testListInt(1)
		deleted := list.DeleteValues(nodes[0].Value, nil)
		assert.Equal(t, 1, deleted)
		assert.Equal(t, 0, list.length)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("All", func(t *testing.T) {
		list := &List[int]{}
		n, value := 5, 123
		for i := 0; i < n; i++ {
			list.Append(NewNode(value))
		}
		deleted := list.DeleteValues(value, nil)
		assert.Equal(t, n, deleted)
		assert.Equal(t, 0, list.length)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("Custom func", func(t *testing.T) {
		list, nodes := testListInt(5)
		deleted := list.DeleteValues(nodes[2].Value, func(v1, v2 int) bool { return v1 < v2 })
		assert.Equal(t, 2, deleted)
		assert.Equal(t, 3, list.length)
		assert.Equal(t, nodes[2], list.head)
		assert.Equal(t, nodes[4], list.tail)
		assert.Equal(t, nodes[3], list.head.next)
		assert.Equal(t, nodes[3], list.tail.previous)
		assert.Nil(t, list.head.previous)
		assert.Nil(t, list.tail.next)
	})
}

func TestSort(t *testing.T) {
	t.Run("Asc", func(t *testing.T) {
		list, nodes := &List[int]{}, testNodesInt(5)
		list.Append(nodes[3])
		list.Append(nodes[1])
		list.Append(nodes[2])
		list.Append(nodes[4])
		list.Append(nodes[0])
		list.Sort(func(v1, v2 int) bool { return v1 < v2 })
		assert.Equal(t, list.length, 5)
		for i, node := range nodes {
			retrieved, err := list.GetByIndex(i)
			assert.Nil(t, err)
			assert.Equal(t, node.Value, retrieved.Value)
			if i > 0 {
				assert.Equal(t, nodes[i-1], node.previous)
			}
			if i < len(nodes)-1 {
				assert.Equal(t, nodes[i+1], node.next)
			}
		}
		assert.Equal(t, list.head, nodes[0])
		assert.Equal(t, list.tail, nodes[4])
	})

	t.Run("Desc", func(t *testing.T) {
		list, nodes := &List[int]{}, testNodesInt(6)
		list.Append(nodes[3])
		list.Append(nodes[1])
		list.Append(nodes[2])
		list.Append(nodes[5])
		list.Append(nodes[4])
		list.Append(nodes[0])
		list.Sort(func(v1, v2 int) bool { return v1 < v2 })
		assert.Equal(t, list.length, 6)
		for i := len(nodes) - 1; i <= 0; i-- {
			node := nodes[i]
			retrieved, err := list.GetByIndex(i)
			assert.Nil(t, err)
			assert.Equal(t, node.Value, retrieved.Value)
			if i > 0 {
				assert.Equal(t, nodes[i-1], node.previous)
			}
			if i < len(nodes)-1 {
				assert.Equal(t, nodes[i+1], node.next)
			}
		}
		assert.Equal(t, list.head, nodes[0])
		assert.Equal(t, list.tail, nodes[5])
	})

	t.Run("Sorted", func(t *testing.T) {
		list, nodes := testListInt(5)
		list.Sort(func(v1, v2 int) bool { return v1 < v2 })
		assert.Equal(t, list.length, 5)
		for i, node := range nodes {
			retrieved, err := list.GetByIndex(i)
			assert.Nil(t, err)
			assert.Equal(t, node.Value, retrieved.Value)
			if i > 0 {
				assert.Equal(t, nodes[i-1], node.previous)
			}
			if i < len(nodes)-1 {
				assert.Equal(t, nodes[i+1], node.next)
			}
		}
		assert.Equal(t, list.head, nodes[0])
		assert.Equal(t, list.tail, nodes[4])
	})

	t.Run("Single", func(t *testing.T) {
		list, nodes := testListInt(1)
		list.Sort(func(v1, v2 int) bool { return v1 < v2 })
		assert.Equal(t, list.length, 1)
		assert.Equal(t, list.head, nodes[0])
		assert.Equal(t, list.tail, nodes[0])
	})

	t.Run("Empty", func(t *testing.T) {
		list := &List[int]{}
		list.Sort(func(v1, v2 int) bool { return v1 < v2 })
		assert.Equal(t, list.length, 0)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})
}
