package main

import (
	"fmt"
	"sync"
)

// This program is deterministic due to the use of sync.WaitGroup to ensure
// that the main goroutine waits for the spawned goroutine to finish before exiting.

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("hello and no race condition!")
	}()

	wg.Wait()
}
