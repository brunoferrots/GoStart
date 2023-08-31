package caseS

import "fmt"

func Case() {
	var i int
	fmt.Println("Digite a integer value between 1..5")
	fmt.Scanln(&i)

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Invalid value")
	}
}
