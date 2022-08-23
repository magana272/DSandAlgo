package main

import "fmt"

type iStack interface {
	pop(x int)
	push(x int)
}

type Node struct {
	value    int
	nextNode *Node
}

type Stack struct {
	size     int
	headNode *Node
}

func (s *Stack) Push(x int) {
	var node = Node{value: x}
	if s.headNode != nil {
		node.nextNode = s.headNode
	}
	s.headNode = &node
}
func (s *Stack) Pop() int {
	var pop int
	if s.headNode != nil {
		pop = s.headNode.value
		s.headNode = s.headNode.nextNode
		return pop
	} else {
		fmt.Println("No Elements")
		panic("No Values")
	}
}

func main() {
	var newstack = Stack{}
	newstack.Push(2)
	newstack.Push(1)
	newstack.Push(1)
	newstack.Push(1)
	fmt.Println(newstack.Pop())
	fmt.Println(newstack.Pop())
	fmt.Println(newstack.Pop())
	fmt.Println(newstack.Pop())

}
