package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(10)
	intList.PushBack(20)
	intList.PushBack(30)
	intList.PushFront("Front")
	for el := intList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
