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

	var ready sync.WaitGroup

	subscribe := func(c *sync.Cond, fn func()) {
		ready.Add(1)
		go func() {
			c.L.Lock()
			ready.Done()
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
