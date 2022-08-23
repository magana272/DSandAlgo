package main

import (
	"fmt"
)

func sendValue(ch chan int, ch2 chan int) {
	v := 1
	ch <- v
	v2 := <-ch2
	fmt.Println(v, v2)
	fmt.Println("fin")

}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go sendValue(ch1, ch2)
	var v2 int
	v := 2
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
}
