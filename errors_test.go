package godll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOutOfRangeError(t *testing.T) {
	err := &IndexOutOfRangeError{Index: 123}
	assert.Equal(t, "Index 123 is out of range!\n", err.Error())
}

func TestNegativeIndexError(t *testing.T) {
	err := &NegativeIndexError{Index: -123}
	assert.Equal(t, "Index -123 is a negative number!\n", err.Error())
}

func TestNodeNotFoundError(t *testing.T) {
	err := &NodeNotFoundError[int]{Node: NewNode(123)}
	assert.Equal(t, "Node not found: &{Value:123 next:<nil> previous:<nil>}\n", err.Error())
}
