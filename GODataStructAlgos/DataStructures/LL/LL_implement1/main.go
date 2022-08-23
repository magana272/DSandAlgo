package main

func main() {
	var newLL NormalLL
	newLL.add(1)
	newLL.add(2)
	newLL.add(3)
	newLL.add(4)
	newLL.add(5)
	newLL.add(6)
	newLL.PrintLL()   // Print 1 2 3 4 5 6
	newLL.llReverse() //  In-place reverse
	newLL.PrintLL()   // Prints 6 5 4 3 2 1
}
