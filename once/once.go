package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	var number int

	increment := func() {
		number++
	}

	decrement := func() {
		number--
	}

	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count is: %d\n", number)

	// In this example, the output will always be "Count is: 1" because the increment function is executed first and only once.
	// The decrement function is ignored by sync.Once.
	// The once primitive ensures that only one of the functions is executed, it will ignore any other calls after the first one. no matter what they are.

}
