package main

import (
	"fmt"
	"time"
)

func initTimeSeq() func() int {
	t := time.Now().UnixNano()
	return func() int {
		return int(time.Now().UnixNano() - t)
	}
}

func main() {
	timeSince := initTimeSeq()

	fmt.Println(timeSince())

	fmt.Println(timeSince())

	// time.Sleep(1 * time.Second)

	fmt.Println(timeSince())

	time.Sleep(12 * time.Millisecond)

	fmt.Println(timeSince())

	timeSince1 := initTimeSeq()
	fmt.Println(timeSince1())
	fmt.Println("")
	time.Sleep(12 * time.Millisecond)

	fmt.Println(timeSince())

	fmt.Println(timeSince())

	fmt.Println(timeSince())
}
