// Say we hava an Iprocess interface with a process method,
// the Adapter class implements the process method and has an adaptee instance ass an attribute.
// The adaptee class has a convert method and an Adaptee instance variable

package main

import (
	"fmt"
)

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

type Adaptee struct {
	adapterType int
}

//	func (adap Adaptee) process() {
//		fmt.Println("Adaptee Process")
//	}
func (adap Adapter) process() {
	fmt.Println("Adapter Process")
	adap.adaptee.convert()
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func main() {
	fmt.Println("Adaptor Structal Design")
	var processor IProcess = Adapter{}
	processor.process()

}
