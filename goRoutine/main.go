package main

import (
	"fmt"
	"time"
)

func compute(val int) {
	for i := 0; i < val; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func main() {
	fmt.Println("Concurrent in GO")
	go compute(10)
	go compute(10)
	fmt.Scanln()
}
