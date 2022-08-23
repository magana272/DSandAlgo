package main

import "fmt"

func main() {

	var testTwoDArr [][]int
	testTwoDArr = make([][]int, 3)
	var col int
	col = 1
	for i := range testTwoDArr {
		testTwoDArr[i] = make([]int, col)
	}
	fmt.Println(testTwoDArr)

}
