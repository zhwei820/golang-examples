package main

import (
	"fmt"
)

func main() {
	stop := make(chan bool)

	go func() {
		fmt.Println("this is Println inside an anonymous goroutine")
		stop <- true
	}()
	for ii := 0; ii < 200; ii++ {
		func() {
			fmt.Println("d")
		}()
	}

	<-stop
}
