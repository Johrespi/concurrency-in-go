package main

import (
	"context"
	"time"
)

func main() {
	timeout()
}

func timeout() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	done := make(chan any)

	go func() {
		time.Sleep(4 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		println("Goroutine completed")
	case <-ctx.Done():
		println("Timeout exceeded:", ctx.Err().Error())
	}
}
