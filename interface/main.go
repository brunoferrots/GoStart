package main

import "fmt"

type Printer interface {
	Print()
}

type Point struct {
	fName, lName string
}

type Triangule struct {
	x, y int
}

type Retangule struct {
	x, y int
}

func (p *Point) Print() {
	fmt.Printf("(%v, %v) \n", p.fName, p.lName)
}

func area(i interface{}) int {
	if t, ok := i.(Triangule); ok {
		return t.x * t.y / 2
	}

	if r, ok := i.(Retangule); ok {
		return r.x * r.y
	}

	return 0
}

func main() {
	p := Point{fName: "Bruno", lName: "Ferro"}
	p.Print()

	t := Triangule{x: 4, y: 8}
	r := Retangule{x: 4, y: 8}
	i := 5

	fmt.Println("Triangule:", area(t))
	fmt.Println("Retangule:", area(r))
	fmt.Println("Any integer value:", area(i))

	//Example nil handle
	//var i Printer
	//i.Print()
}
