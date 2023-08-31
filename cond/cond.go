package cond

import "fmt"

func CondIf() {
	k, x := 6, 4

	if y := x / 2; k > 5 {
		fmt.Printf("K value %d\n", k)
		fmt.Printf("Y value %d\n", y)
		fmt.Printf("X value %d\n", x)
	} else {
		fmt.Printf("K value is %d", k)
	}
}
