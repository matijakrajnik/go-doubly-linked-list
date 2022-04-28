package godll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNodeInt(t *testing.T) {
	node := NewNode(8)
	assert.Equal(t, 8, node.Value)
	assert.Nil(t, node.Next())
}

func TestNewNodeFloat64(t *testing.T) {
	node := NewNode(4.2)
	assert.Equal(t, 4.2, node.Value)
	assert.Nil(t, node.Next())
}

func TestNewNodeString(t *testing.T) {
	node := NewNode("Bruce Wayne")
	assert.Equal(t, "Bruce Wayne", node.Value)
	assert.Nil(t, node.Next())
}

func TestNewNodeStruct(t *testing.T) {
	person := PersonTest{FirstName: "Bruce", LastName: "Wayne"}
	node := NewNode(person)
	assert.Equal(t, person, node.Value)
	assert.Equal(t, person.FirstName, node.Value.FirstName)
	assert.Equal(t, person.LastName, node.Value.LastName)
	assert.Nil(t, node.Next())
}

func TestNext(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)
	assert.Equal(t, node2, node1.Next())
	assert.Equal(t, node3, node2.Next())
	assert.Nil(t, node3.Next())
}

func TestPrevious(t *testing.T) {
	list := &List[int]{}
	node1 := NewNode(8)
	node2 := NewNode(2)
	node3 := NewNode(5)
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)
	assert.Nil(t, node1.Previous())
	assert.Equal(t, node1, node2.Previous())
	assert.Equal(t, node2, node3.Previous())
}
