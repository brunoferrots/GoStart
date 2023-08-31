/* This is my first programming in the GO lang.*/

package main

import (
	"fmt"
	"learngo/slice"
)

//const age int = int(19.0)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	var number = 19
	var age *int
	age = &number
	*age = 26

	fmt.Println(*age)
	fmt.Println(number)

	p1 := Person{"Bruno", "Ferro"}
	p2 := Person{firstName: "Maria"}

	fmt.Println(p1, p2)

	//var fruits [5]string
	fruits := [5]string{"apple", "orange", "banana", "pineapple", "strawberry"}
	fmt.Println(fruits, fruits[0])

	var arr [4]int
	var slc []int = arr[1:4]
	fmt.Println(slc)

	slice.Slice()

}
