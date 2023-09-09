package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{5, 9, 1, 6, 4}
	wg := &sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			arr[i] *= 3
			fmt.Println("Gouroutine", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(arr)
}
