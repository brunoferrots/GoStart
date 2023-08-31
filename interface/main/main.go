package main

import "fmt"

type Printer interface {
	Print()
}

type Point struct {
	fName, lName string
}

func (p *Point) Print() {
	fmt.Printf("(%v, %v)", p.fName, p.lName)
}

func main() {
	p := Point{fName: "Bruno", lName: "Ferro"}
	p.Print()
}
