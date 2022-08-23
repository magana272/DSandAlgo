package main

import (
	"fmt"
	"time"
)

type data struct {
	length        int
	dateCollected time.Time
}

func main() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[2])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
	collected := data{1, time.Now()}
	fmt.Println(collected.length, collected.dateCollected.Format("2006-1-2 15:4:5"))

}
