package main

import (
	"fmt"
	"sync"
)

// This program is still undeterministic, because the order of execution of goroutines is not guaranteed.
// However, it uses a RWMutex to allow multiple readers or a single writer at a time,
// preventing data races

var number int = 5000

func main() {

	var rw sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(3)

	// A reader
	go func(number *int) {
		defer wg.Done()
		defer rw.RUnlock()
		rw.RLock()
		fmt.Printf("Im reading: %d\n", *number)
	}(&number)

	// Another reader
	go func(number *int) {
		defer wg.Done()
		defer rw.RUnlock()
		rw.RLock()
		fmt.Printf("Im also reading: %d\n", *number)
	}(&number)

	// A writer
	go func(number *int) {
		defer wg.Done()
		defer rw.Unlock()
		rw.Lock()
		*number += 5000
		fmt.Printf("I wrote: %d\n", *number)
	}(&number)

	wg.Wait()
	fmt.Println("End of goroutines")

}
