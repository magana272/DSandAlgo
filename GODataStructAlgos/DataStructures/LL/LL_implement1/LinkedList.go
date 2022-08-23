package main

import (
	"fmt"
)

type LinkedList interface {
	PrintLL()
	add(byte)
	delete(byte)
}

type Node struct {
	value byte
	next  *Node
}
type NormalLL struct {
	head *Node
}

func (L *NormalLL) add(data byte) {
	if L.head == nil {
		L.head = &Node{value: data}
	} else {
		L.head.add(data)
	}
}
func (N *Node) add(data byte) {
	if N.next == nil {
		N.next = &Node{value: data}
	} else {
		N.next.add(data)
	}
}

func (L *NormalLL) PrintLL() {
	if L.head == nil {
		return
	} else {
		L.head.PrintLL()
	}
}
func (N *Node) PrintLL() {
	if N == nil {
		return
	} else {
		fmt.Println(N.value)
		N.next.PrintLL()
	}
}
func (L *NormalLL) llReverse() {
	if L.head == nil {
		return
	} else {
		curr := L.head.llReverse()
		L.head = curr

	}
}
func (N *Node) llReverse() *Node {
	curr := N
	var prev *Node
	var next *Node
	for curr.next != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	curr.next = prev
	return curr
}
