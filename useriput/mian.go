package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Print("Enter you name: ")
	name, err := r.ReadString('\n')
	if err != nil {
		fmt.Errorf("Error Occured")
	}
	w.WriteString(name)
	fmt.Print("Enter a number: ")
	var i int
	val, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Errorf("Error Occured")
	}
	fmt.Print("Your name is: " + name)
	if err != nil {
		fmt.Errorf("Error Occured")
	}
	if val > 0 && val < 10 {
		println("Is in the range 0, 10")
		return

	}
	fmt.Printf("%b", val)

}
