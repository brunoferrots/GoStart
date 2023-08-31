package def

import "fmt"

func Def() {
	fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("c")
	defer fmt.Println("d")
}
