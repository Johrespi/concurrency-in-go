package main

import "fmt"

func main() {

	done := make(chan struct{})
	for {
		select {
		case <-done:
			return
		default:
			// Do non-preemptable work
			fmt.Println("Working...")
		}
	}
}
