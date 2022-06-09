// Have the function LRUCache(strArr) take the array of characters stored in strArr,
// which will contain characters ranging from A to Z in some arbitrary order,
// and determine what elements still remain in a virtual cache that can hold
// up to 5 elements with an LRU cache algorithm implemented.

package main

import (
	"fmt"
)

func main() {
	readLine := []string{"A", "B", "C", "D", "A", "E", "D", "Z"} // expected result: C-A-E-D-Z
	// readLine := []string{"A", "B", "A", "C", "A", "B"} // expected result: C-A-B
	// readLine := []string{"A", "X", "C", "E", "D", "A", "M", "E", "K"} // expected result: D-A-M-E-K

	nodes := NewNode(5)
	for _, str := range readLine {
		nodes.Push(str)
	}

	fmt.Println(nodes.Display())
}

// Node struct
type Node struct {
	capacity int // store max capacity of Node
	caches   []string
}

// Initiate Node with given capacity
func NewNode(cap int) Node {
	return Node{
		capacity: cap,
		caches:   []string{},
	}
}

// Print all value of node.caches
func (node *Node) Display() (display string) {
	for _, str := range node.caches {
		if display == "" {
			display += fmt.Sprintf("%s", str)
		} else {
			display += fmt.Sprintf("-%s", str)
		}
	}

	return
}

// Push string to a node.caches
func (node *Node) Push(str string) {
	index, exists := node.Contains(str)
	if len(node.caches) < node.capacity {
		if !exists {
			node.caches = append(node.caches, str)
			return
		}
	}

	node.Pop(index)
	node.Push(str)
}

// Pop specific index from a node.caches
func (node *Node) Pop(index int) {
	node.caches = append(node.caches[:index], node.caches[index+1:]...)
}

// Contains will check on node.caches with specific given string
func (node *Node) Contains(str string) (key int, exists bool) {
	for i := range node.caches {
		if node.caches[i] == str {
			return i, true
		}
	}

	return 0, false
}
