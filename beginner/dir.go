package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println(strings.Replace(dir, " ", "\\ ", -1))
	fmt.Println(dir)
}
