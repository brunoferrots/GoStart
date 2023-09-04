package main

import "fmt"

type Rectangule struct {
	with, height int
}

func (r *Rectangule) Area() int {
	return r.with * r.height
}

func main() {
	r := Rectangule{4, 6}

	fmt.Printf("Rectangule area is %d", r.Area())
}
