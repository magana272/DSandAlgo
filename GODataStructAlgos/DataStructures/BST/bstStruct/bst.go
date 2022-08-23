package main

import "fmt"

type BST interface {
	find()
	insert()
	delete()
}

type Tree struct {
	root *Node
}
type Node struct {
	key   byte
	value string
	left  *Node
	right *Node
}

// Insert

// tree Insert
func (t *Tree) insert(data byte, str string) {
	if t.root == nil {
		t.root = &Node{key: data, value: str}
	} else {
		t.root.insert(data, str)
	}
}

//Node inset

func (n *Node) insert(data byte, str string) *Node {
	if n == nil {
		return &Node{key: data, value: str}
	}
	if data > n.key {
		n.right = n.right.insert(data, str)
	}
	if data < n.key {
		n.left = n.left.insert(data, str)
	}
	return n
}

// / Print
// Tree
func (t *Tree) printVal() {
	if t.root == nil {
		return
	} else {
		t.root.printVal()
	}
}

// Node
func (n *Node) printVal() {
	if n == nil {
		return
	} else {
		fmt.Println(n.key, n.value)
		fmt.Println(" ")
		fmt.Println("left")
		n.left.printVal()
		fmt.Println("right")
		n.right.printVal()
	}
}

func (t *Tree) find(i byte) {
	if t.root == nil {
		fmt.Println("Tree is empty")
	} else {
		t.root.find(i)
	}
}
func (n *Node) find(i byte) {
	if n == nil {
		fmt.Println("Key not found")
		return
	}
	if n.key == i {
		fmt.Println("Found the key:")
		fmt.Println("Key Value:", n.value)
	}
	if i > n.key {
		n.right.find(i)
	}
	if i < n.key {
		n.left.find(i)
	}
}
