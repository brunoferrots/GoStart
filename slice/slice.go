package slice

import "fmt"

func Slice() {
	s1 := make([]int, 5)
	s2 := make([]int, 0, 5)

	s1 = append(s1, 1)
	s2 = append(s2, 4)

	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)
}
