package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan any, strings <-chan string) <-chan any {
		terminated := make(chan any)
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan any)
	terminated := doWork(done, nil)

	go func() {
		// Cancel operation after 1 second
		time.Sleep(1 * time.Second)
		fmt.Println("cancelling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("done")
}
