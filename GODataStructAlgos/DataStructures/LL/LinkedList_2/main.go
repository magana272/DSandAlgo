// Node class
package main

import (
	"fmt"
)

// Node Class
type Node struct {
	property int
	nextNode *Node
}

// Linked List Class
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}
func (LL *LinkedList) IterateList() {
	for curr := LL.headNode; curr != nil; curr = curr.nextNode {
		fmt.Println(curr.property)
	}
}

func main() {
	fmt.Println("hello")
	var LL = LinkedList{}
	LL.AddToHead(1)
	LL.AddToHead(2)
	LL.AddToHead(3)
	LL.AddToHead(4)
	LL.AddToHead(5)
	LL.IterateList()
}
