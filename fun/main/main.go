package main

import "fmt"

func sqr(x *int) (y int) {
	y = *x * *x
	return
}

func calc(p int) func(int, int) int {
	if p >= 0 {
		return func(i1, i2 int) int {
			return i1 + i2
		}
	}

	return func(i1, i2 int) int {
		return i1 - i2
	}
}

func fibonacci() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

func main() {
	//var value int
	fmt.Println("Digit a value for cal sqr it")
	//fmt.Scanln(&value)

	//fmt.Printf("The func. sqr result is %d \n", sqr(&value))

	add := calc(5)
	sub := calc(-5)

	fmt.Printf("Result values of func Add: %d and Sub: %d", add(3, 2), sub(3, 2))

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

}
