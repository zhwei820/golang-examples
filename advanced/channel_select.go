// 协程goroutine

//      不由OS调度，而是用户层自行释放CPU，从而在执行体之间切换。Go在底层进行协助实现
//      涉及系统调用的地方由Go标准库协助释放CPU
//      总之，不通过OS进行切换，自行切换，系统运行开支大大降低

// 通道channel

// 并发编程的关键在于执行体之间的通信，go通过通过channel进行通信
// channel可以认为类似其他OS体系中的消息队列，只不过在go中原生支持，因而易用

// 消息队列有哪些值得关注的地方？常见问题包括创建、关闭或删除、阻塞、超时、优先级等，golang中也不例外。罗列如下：
//      可否探测队列是满或空？或者说是否可以不阻塞地尝试读写？
//      读阻塞和写阻塞时关闭会怎样？
//      关闭后未读取的消息会被抛弃？
//      往关闭的channel发送数据或读取数据会怎样？
//      怎样探测channel的关闭？
//      两个地方读或写阻塞同一个channel，有没有优先级？
//      是否可以设定阻塞的超时时间？
//      阻塞时怎样可以被弹出来？比如某些信号？
// 事实上，知道存在这些问题并进行分门别类是重要的，但知道这些问题的答案却不紧要，因为一般不会太过古怪，使用时临时试验一下即可。

// 已知的部分答案：
//      好像不能不阻塞地尝试读写
//      关闭会导致退出阻塞(似乎是一个不错的特性)
//      可以探测关闭
//      channel本身不能设定超时
// 了解这些似乎已经足够。

// 与众不同的地方值得我们重点留意，包括：
//      除基本读写方式外，还有哪些特别的读写方式？在阻塞、关闭、超时方面又有什么不同？发现了select、range两个关键字
//      推荐的多通道读
//      推荐的同步方法
//      推荐的超时方法

// range

// range可以在for循环中读取channel
// Go文档的翻译文是：对于信道，其迭代值产生为在该信道上发送的连续值，直到该信道被关闭。若该信道为 nil，则range表达式将永远阻塞
// 经过试验，range会阻塞，并且可以通过关闭channel来解除阻塞。

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

	}()

	for w := range ch {
		fmt.Println("fmt print", w)
		if w > 5 {
			//break // 在这里break循环也可以
			close(ch)
		}
	}
	fmt.Println("after range or close ch!")
	//////////////////////////////////////////////////////////////////////////////////
	// select

	// select可以实现无阻塞的多通道尝试读写，以及阻塞超时

	var c, c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1: //如果能走通任何case则随机走一个
		print("received ", i1, " from c1\n")
	case c2 <- i2:
		print("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			print("received ", i3, " from c3\n")
		} else {
			print("c3 is closed\n")
		}
	default: // 如果case都阻塞，则走default，如果无default，则阻塞在case
		// default中可以不读写任何通道，那么只要default提供不阻塞的出路，就相当于实现了对case的无阻塞尝试读写
		print("no communication\n")
	}

	// 实现阻塞超时的方法是，只要不给default出路，而在case中实现一个超时

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 这是等待1秒钟
		timeout <- true
	}()

	// 用timeout这个通道作为阻塞超时的出路
	select {
	case <-ch:
	// 处理从ch中读到的数据
	case <-timeout:
		// 如果case都阻塞了，那么1秒钟后会从这里找到出路
	}

}

// Golang的并发编程还有其他细节，但以上是最主要脉络。
