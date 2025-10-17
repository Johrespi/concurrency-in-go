package main

import (
	"fmt"
)

func main() {

	go func() {
		fmt.Println("Hello from a goroutine using an anonymous function")
	}() // Note the () at the end to invoke the function immediately

	// Do other things in main goroutine
	anotherGoroutine := func() {
		fmt.Println("Hello from another goroutine using a function variable")
	}
	go anotherGoroutine()
}
