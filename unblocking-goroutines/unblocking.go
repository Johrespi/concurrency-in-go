package main

import (
	"fmt"
	"sync"
)

func main() {

	begin := make(chan int)
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("Goroutine %v started\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()

}
