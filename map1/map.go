package map1

import "fmt"

func Map1() {
	m := make(map[int]string)
	m[1] = "Monkey"
	m[2] = "Tiger"
	m[3] = "Lion"

	fmt.Println(m)
	delete(m, 2)

	k, ok := m[1]
	fmt.Println(m)
	fmt.Println(k, ok)
}
