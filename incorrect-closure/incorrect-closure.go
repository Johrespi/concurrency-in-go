package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Iterate over a slice of strings
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// This goroutine closes over the loop variable `salutation`
			// Since the goroutines may start executing after the loop finishes,
			// all goroutines may see the last value of `salutation`.
			// On many machines, this will print "good day" three times.
			fmt.Println(salutation)
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// ---------------------------------------
	// Explanation:
	// The goroutine is running a closure that has captured the loop variable `salutation` by reference.
	// As the loop iterates, `salutation` is assigned the next string in the slice.
	// Because goroutines may be scheduled at any time in the future,
	// it is indeterminate which value they will print.
	// The Go runtime moves `salutation` to the heap if the loop finishes before the goroutines run,
	// so the variable is still accessible and not garbage collected.
	// This is why you usually see "good day" printed multiple times.
	//
	// Correct approach:
	// Pass a copy of `salutation` to the goroutine so each iteration has its own value:
	// go func(s string) {
	//     defer wg.Done()
	//     fmt.Println(s)
	// }(salutation)
}
