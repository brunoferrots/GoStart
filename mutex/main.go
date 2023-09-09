package main

import (
	"fmt"
	"sync"
	"time"
)

type MutualExclusion struct {
	mutex sync.Mutex
	value int
}

func (m *MutualExclusion) Double() {
	m.mutex.Lock()
	m.value *= 2
	m.mutex.Unlock()
}

func main() {
	m := MutualExclusion{value: 5}
	go m.Double()
	go m.Double()

	time.Sleep(time.Second)
	fmt.Println(m.value)
}
