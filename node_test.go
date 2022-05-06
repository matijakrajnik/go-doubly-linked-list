package godll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		node := NewNode(8)
		assert.Equal(t, 8, node.Value)
		assert.Nil(t, node.Next())
	})

	t.Run("Float64", func(t *testing.T) {
		node := NewNode(4.2)
		assert.Equal(t, 4.2, node.Value)
		assert.Nil(t, node.Next())
	})

	t.Run("String", func(t *testing.T) {
		node := NewNode("Bruce Wayne")
		assert.Equal(t, "Bruce Wayne", node.Value)
		assert.Nil(t, node.Next())
	})

	t.Run("Struct", func(t *testing.T) {
		person := PersonTest{FirstName: "Bruce", LastName: "Wayne"}
		node := NewNode(person)
		assert.Equal(t, person, node.Value)
		assert.Equal(t, person.FirstName, node.Value.FirstName)
		assert.Equal(t, person.LastName, node.Value.LastName)
		assert.Nil(t, node.Next())
	})
}

func TestNext(t *testing.T) {
	_, nodes := testListInt(3)

	assert.Equal(t, nodes[1], nodes[0].Next())
	assert.Equal(t, nodes[2], nodes[1].Next())
	assert.Nil(t, nodes[2].Next())
}

func TestPrevious(t *testing.T) {
	_, nodes := testListInt(3)

	assert.Nil(t, nodes[0].Previous())
	assert.Equal(t, nodes[0], nodes[1].Previous())
	assert.Equal(t, nodes[1], nodes[2].Previous())
}
