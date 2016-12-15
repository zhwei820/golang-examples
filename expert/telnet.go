package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(c net.Conn, msgchan chan<- string) {
	defer c.Close()
	fmt.Printf("Connection from %v established.\n", c.RemoteAddr())
	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("Message Received:", string(message))
	c.Write([]byte("PDFLDFSDFSJDKFENEFJNSDKFNSDNFSKDJHFSDKFHJDJFDJSFLKDSJ" + "\n"))
	c.Close()
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
	return
}

func printMessages(msgchan <-chan string) {
	var count int = 0
	for {
		msg := strings.TrimSpace(<-msgchan)
		count++
		fmt.Printf("Data %d: %s\n", count, msg)
	}
}

func main() {
	flag.Parse()
	port := ":" + flag.Arg(0)
	if port == ":" {
		port = ":2333"
	}
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	msgchan := make(chan string)
	go printMessages(msgchan)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, msgchan)
	}
}
