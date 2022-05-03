# Golang Doubly Linked List

Module `godll` is Golang implementation of doubly linked list with nodes containing comparable generic values.

## Installation

Module requires at least Golang 1.18 version. Install it with:

```bash
go get github.com/matijakrajnik/godll
```

## Usage

### Creating new list

```go
package main

import (
 "fmt"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 fmt.Printf("%+v\n", l)
 // Output:
 // &{head:<nil> tail:<nil> length:0}
}
```

### Append and prepend new node

Node can be prepended (added at the beginning of the list) or appended (added at the end of the list).

```go
package main

import (
 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 l.Append(godll.NewNode(5))
 l.Prepend(godll.NewNode(8))
}
```

### Printing list

Node values can be printed to passed io.Writer.

```go
package main

import (
 "os"

 "github.com/matijakrajnik/godll"
)

type Person struct {
 First string
 Last  string
}

func main() {
 intL := &godll.List[int]{}
 intL.Prepend(godll.NewNode(6))
 intL.Append(godll.NewNode(5))
 intL.Append(godll.NewNode(2))
 intL.Prepend(godll.NewNode(9))
 intL.Print(os.Stdout)
 // Output:
 // 9 6 5 2

 personL := &godll.List[Person]{}
 personL.Append(godll.NewNode(Person{First: "Bruce", Last: "Wayne"}))
 personL.Prepend(godll.NewNode(Person{First: "Clark", Last: "Kent"}))
 personL.Print(os.Stdout)
 // Output:
 // {First:Clark Last:Kent} {First:Bruce Last:Wayne}
}
```

### Retrieving values

Node values are compared using "==" by default. Pass nil as compare function for default behaviour:

```go
package main

import (
 "fmt"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 l.Append(godll.NewNode(6))
 l.Append(godll.NewNode(5))
 l.Append(godll.NewNode(5))
 l.Append(godll.NewNode(2))
 l.Append(godll.NewNode(9))
 l.Append(godll.NewNode(2))

 fmt.Printf("Head: %+v\n", l.Head())
 fmt.Printf("Tail: %+v\n", l.Tail())
 // Output:
 // Head: &{Value:6 next:0xc0000ac030 previous:<nil>}
 // Tail: &{Value:2 next:<nil> previous:0xc0000ac078}

 node, _ := l.GetByIndex(3)
 fmt.Printf("Node at index 3: %+v\n", node)
 // Output:
 // Node at index 3: &{Value:2 next:0xc0000ac078 previous:0xc0000ac048}

 index, node := l.GetByValue(9, nil)
 fmt.Printf("Value 9 is at index: %v. Node value is: %+v\n", index, node)
 // Output:
 // Value 9 is at index: 4. Node value is: &{Value:9 next:0xc00000c0a8 previous:0xc00000c078}

 all := l.GetAllValues(5, nil)
 fmt.Printf("All nodes with value 5 found at: %+v\n", all)
 // Output:
 // All nodes with value 5 found at: map[1:0xc00000c048 2:0xc00000c060]
}
```

Custom compare function can be defined like this:

```go
package main

import (
 "fmt"

 "github.com/matijakrajnik/godll"
)

type Person struct {
 ID    int
 First string
 Last  string
}

func main() {
 l := &godll.List[Person]{}
 l.Append(godll.NewNode(Person{ID: 1, First: "Bruce", Last: "Wayne"}))
 l.Append(godll.NewNode(Person{ID: 2, First: "Clark", Last: "Kent"}))
 p := Person{ID: 2}
 index, node := l.GetByValue(p, func(v1, v2 Person) bool { return v1.ID == v2.ID })
 fmt.Printf("Person with ID=2 is at index: %v. Node value is: %+v\n", index, node)
 // Output:
 // Person with ID=2 is at index: 1. Node value is: &{Value:{ID:2 First:Clark Last:Kent} next:<nil> previous:0xc00007e040}
}
```

### Deleting nodes

There are multiple ways supported to delete nodes. Nodes can be deleted at certain index:

```go
package main

import (
 "os"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 l.Append(godll.NewNode(6))
 l.Append(godll.NewNode(9))
 l.Append(godll.NewNode(2))

 l.Print(os.Stdout)
 l.DeleteAt(0)
 l.Print(os.Stdout)
 l.DeleteAt(1)
 l.Print(os.Stdout)
 l.DeleteAt(0)
 l.Print(os.Stdout)
 // Output:
 // 6 9 2
 // 9 2
 // 9
}
```

Pointer to specific node can be passed to delete it from the list:

```go
package main

import (
 "os"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 node1 := godll.NewNode(6)
 node2 := godll.NewNode(9)
 node3 := godll.NewNode(2)
 l.Append(node1)
 l.Append(node2)
 l.Append(node3)

 l.Print(os.Stdout)
 l.DeleteNode(node1)
 l.Print(os.Stdout)
 l.DeleteNode(node2)
 l.Print(os.Stdout)
 l.DeleteNode(node3)
 l.Print(os.Stdout)
 // Output:
 // 6 9 2
 // 9 2
 // 2
}
```

All nodes with specific values can be deleted by passing custom function. You can pass `nil` for custom function to compare node values using `==` by default.

```go
package main

import (
 "fmt"
 "os"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 l.Append(godll.NewNode(6))
 l.Append(godll.NewNode(9))
 l.Append(godll.NewNode(9))
 l.Append(godll.NewNode(9))
 l.Append(godll.NewNode(2))

 l.Print(os.Stdout)
 deleted := l.DeleteValues(9, nil)
 fmt.Printf("Deleted %v nodes\n", deleted)
 l.Print(os.Stdout)
 // Output:
 // 6 9 9 9 2
 // Deleted 3 nodes
 // 6 2
}
```

Or you can pass custom compare function:

```go
package main

import (
 "fmt"
 "os"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 l.Append(godll.NewNode(4))
 l.Append(godll.NewNode(3))
 l.Append(godll.NewNode(1))
 l.Append(godll.NewNode(2))
 l.Append(godll.NewNode(5))

 l.Print(os.Stdout)
 deleted := l.DeleteValues(3, func(v1, v2 int) bool { return v1 < v2 })
 fmt.Printf("Deleted %v nodes\n", deleted)
 l.Print(os.Stdout)
 // Output:
 // 4 3 1 2 5
 // Deleted 2 nodes
 // 4 3 5
}
```

### Swaping nodes

Nodes can be swaped by using their indexes.

```go
package main

import (
 "os"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 l.Append(godll.NewNode(1))
 l.Append(godll.NewNode(2))
 l.Append(godll.NewNode(3))
 l.Append(godll.NewNode(4))
 l.Append(godll.NewNode(5))

 l.Print(os.Stdout)
 l.Swap(1, 2)
 l.Print(os.Stdout)
 // Output:
 // 1 2 3 4 5
 // 1 3 2 4 5
}
```

### Sorting list

List can be sorted by passing sorting function. Use `<` to sort ascending or `>` to sort descending. Sorting is done using merge sort algorithm.

```go
package main

import (
 "os"

 "github.com/matijakrajnik/godll"
)

func main() {
 l := &godll.List[int]{}
 l.Append(godll.NewNode(4))
 l.Append(godll.NewNode(3))
 l.Append(godll.NewNode(1))
 l.Append(godll.NewNode(2))
 l.Append(godll.NewNode(5))

 l.Print(os.Stdout)
 l.Sort(func(v1, v2 int) bool { return v1 < v2 })
 l.Print(os.Stdout)
 l.Sort(func(v1, v2 int) bool { return v1 > v2 })
 l.Print(os.Stdout)
 // Output:
 // 4 3 1 2 5
 // 1 2 3 4 5
 // 5 4 3 2 1
}
```
