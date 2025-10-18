package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello, Channels!"
	}()

	fmt.Println(<-stringStream)
}
