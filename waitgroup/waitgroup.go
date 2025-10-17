package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("First goroutine is sleeping...zzz")
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Second goroutine is sleeping...zzz")
		time.Sleep(1 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutines have finished execution.")

}
