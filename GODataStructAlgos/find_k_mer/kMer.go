package main

import (
	"fmt"
	"os"
	"time"
)

func readLines(string fileName, map[string]int ){

}

func main(){
	fmt.Println("Getting K-Mers..")
	fmt.Println(time.Now())
	fileName := os.Args[1]
	_, err := os.Stat(fileName)
	if err !=nil{
		fmt.Errorf("File doesn't exis", err)
		os.Exit(1)
	}
	readLines(fileName)
	kmer := make(map[string]int)


}