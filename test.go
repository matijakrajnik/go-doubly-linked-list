package godll

import "fmt"

// Used only for testing.
type PersonTest struct {
	FirstName string
	LastName  string
}

// Create slice of int test nodes.
func testNodesInt(n int) []*Node[int] {
	nodes := []*Node[int]{}
	for i := 1; i <= n; i++ {
		node := NewNode(i)
		nodes = append(nodes, node)
	}
	return nodes
}

// Create int test list with test nodes.
func testListInt(n int) (*List[int], []*Node[int]) {
	list := &List[int]{}
	nodes := testNodesInt(n)
	for _, node := range nodes {
		list.Append(node)
	}
	return list, nodes
}

// Create slice of float64 test nodes.
func testNodesFloat64(n int) []*Node[float64] {
	nodes := []*Node[float64]{}
	for i := 1; i <= n; i++ {
		node := NewNode(float64(i) + 0.5)
		nodes = append(nodes, node)
	}
	return nodes
}

// Create float64 test list with test nodes.
func testListFloat64(n int) (*List[float64], []*Node[float64]) {
	list := &List[float64]{}
	nodes := testNodesFloat64(n)
	for _, node := range nodes {
		list.Append(node)
	}
	return list, nodes
}

// Create slice of string test nodes.
func testNodesString(n int) []*Node[string] {
	nodes := []*Node[string]{}
	for i := 1; i <= n; i++ {
		node := NewNode(fmt.Sprint(i))
		nodes = append(nodes, node)
	}
	return nodes
}

// Create string test list with test nodes.
func testListString(n int) (*List[string], []*Node[string]) {
	list := &List[string]{}
	nodes := testNodesString(n)
	for _, node := range nodes {
		list.Append(node)
	}
	return list, nodes
}

// Create slice of struct test nodes.
func testNodesStruct(n int) []*Node[PersonTest] {
	nodes := []*Node[PersonTest]{}
	for i := 1; i <= n; i++ {
		person := PersonTest{
			FirstName: fmt.Sprintf("Bruce%v", i),
			LastName:  fmt.Sprintf("Wayne%v", i),
		}
		node := NewNode(person)
		nodes = append(nodes, node)
	}
	return nodes
}

// Create struct test list with test nodes.
func testListStruct(n int) (*List[PersonTest], []*Node[PersonTest]) {
	list := &List[PersonTest]{}
	nodes := testNodesStruct(n)
	for _, node := range nodes {
		list.Append(node)
	}
	return list, nodes
}
