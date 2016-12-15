package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex
var val = 0

func main() {
	m = new(sync.RWMutex)
	go read(1)
	time.Sleep(50 * time.Microsecond)
	go write(2)
	time.Sleep(50 * time.Microsecond)

	go read(3)
	time.Sleep(5 * time.Second)
}

func read(i int) {
	m.RLock()
	time.Sleep(1 * time.Second)
	println("val: ", val)
	time.Sleep(1 * time.Second)
	m.RUnlock()
}

func write(i int) {
	m.Lock()
	val = 10 + i
	time.Sleep(1 * time.Second)
	m.Unlock()
}
