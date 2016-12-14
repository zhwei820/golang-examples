package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// var chan_int chan int
// var stop_bool chan bool

// func maxprocs1() {
// 	/*
// 	 * Channels are useful for communications between multiple
// 	 * go routines, even when the routines are not runned in parralel.
// 	 *
// 	 */

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("1")
// 		stop_bool <- true
// 	}()
// 	go func() {
// 		time.Sleep(1900 * time.Millisecond)
// 		fmt.Println("2")
// 		stop_bool <- true
// 	}()
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("3")
// 		stop_bool <- true
// 	}()

// 	<-stop_bool
// 	<-stop_bool
// 	<-stop_bool
// }
// func init() {
// 	chan_int = make(chan int, 3)
// 	stop_bool = make(chan bool, 3)
// }
// func waitforbuffer() {
// 	go func(i int) {
// 		for {
// 			time.Sleep(time.Duration(i) * time.Second)

// 			fmt.Println(<-chan_int)

// 			fmt.Println(<-chan_int)

// 			fmt.Println(<-chan_int)

// 			fmt.Println(<-chan_int)

// 			fmt.Println(<-chan_int)

// 			fmt.Println(<-chan_int)
// 		}
// 	}(2)
// 	chan_int <- 1
// 	chan_int <- 2
// 	chan_int <- 5
// 	chan_int <- 3
// }

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	fmt.Println("starting maxprocs1()")
// 	// maxprocs1()
// 	fmt.Println("starting waitforbuffer()")
// 	waitforbuffer()
// 	fmt.Println("finish")
// }
