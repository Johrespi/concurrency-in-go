package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := range 5 {
			channel <- i + 1
		}
	}()

	for i := range channel {
		fmt.Printf("%v\n", i)
	}
}
