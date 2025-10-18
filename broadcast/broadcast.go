package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	var ready sync.WaitGroup // to ensure all goroutines are ready before broadcasting

	subscribe := func(c *sync.Cond, fn func()) {
		ready.Add(1) // indicate a new goroutine is being set up
		go func() {
			c.L.Lock()
			ready.Done() // indicate this goroutine is ready to receive broadcasts, the resource now is locked
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked")
		clickRegistered.Done()
	})

	ready.Wait()

	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
