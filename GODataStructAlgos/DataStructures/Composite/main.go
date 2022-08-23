package main

import (
	"fmt"
)

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet Struct
type Leaflet struct {
	name string
}

// Leaflet class method perform
func (leaf *Leaflet) perform() {
	fmt.Println(leaf.name)
}

type Branch struct{
	leafs []Leaflet
	name
}
