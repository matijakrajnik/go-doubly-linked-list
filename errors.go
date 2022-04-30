package godll

import "fmt"

type IndexOutOfRangeError struct {
	Index int
}

func (e *IndexOutOfRangeError) Error() string {
	return fmt.Sprintf("Index %v is out of range!\n", e.Index)
}

type NegativeIndexError struct {
	Index int
}

func (e *NegativeIndexError) Error() string {
	return fmt.Sprintf("Index %v is a negative number!", e.Index)
}
