package main

import "fmt"

func main() {

	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := range 5 {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()

	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")
}
