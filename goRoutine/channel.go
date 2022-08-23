package main

import (
	"fmt"
	"time"
)

func sendValues(c chan string) {
	fmt.Println("Execute go rutine")
	c <- "Hello String"
	fmt.Println("Fin")
}
func main() {
	ch := make(chan string, 1000)
	defer close(ch)
	go sendValues(ch)
	go sendValues(ch)
	value := <-ch
	fmt.Println(value)
	time.Sleep(1 * time.Second)
}
