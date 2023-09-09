package main

import "fmt"

func sendMessage(message string, ch chan string) {
	ch <- message
}

func main() {
	ch := make(chan string, 2)
	go sendMessage("Hello", ch)
	go sendMessage("World", ch)

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
