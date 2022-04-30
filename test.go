package godll

// Used only for testing.
type PersonTest struct {
	FirstName string
	LastName  string
}

// Create slice of test nodes.
func testNodes(n int) []*Node[int] {
	nodes := []*Node[int]{}
	for i := 1; i <= n; i++ {
		node := NewNode(i)
		nodes = append(nodes, node)
	}
	return nodes
}

// Create test list with test nodes.
func testList(n int) (*List[int], []*Node[int]) {
	list := &List[int]{}
	nodes := testNodes(n)
	for _, node := range nodes {
		list.Append(node)
	}
	return list, nodes
}
