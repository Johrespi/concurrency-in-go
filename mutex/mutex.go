package main

import (
	"fmt"
	"sync"
)

func increment(mu *sync.Mutex, number *int) {
	mu.Lock()
	defer mu.Unlock()
	*number++
	fmt.Printf("Incrementing number: %d\n", *number)
}

func decrement(mu *sync.Mutex, number *int) {
	mu.Lock()
	defer mu.Unlock()
	*number--
	fmt.Printf("Decrementing number: %d\n", *number)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment(&mu, &counter)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement(&mu, &counter)
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter: %d", counter)
}
