## Concurrency in Go

This repository contains examples of concurrency patterns in Go. It covers goroutines, channels, and other concurrency primitives provided by the Go programming language.

## What is Concurrency?
Concurrency is the ability of a program to manage multiple tasks at the same time. In Go, concurrency is achieved through goroutines and channels, which allow for efficient multitasking and communication between different parts of a program.

## How Goroutines work?
Goroutines are lightweight threads managed by the Go runtime. They are created using the `go` keyword followed by a function call. When a goroutine is created, it runs concurrently with other goroutines in the same program. They're not OS threads nor green threads; instead, the Go runtime multiplexes many goroutines onto a smaller number of OS threads.The Go's runtime is responsible for scheduling and managing these goroutines.

They follow the fork-join model, where a goroutine can spawn other goroutines and wait for them to complete before proceeding. This model allows for efficient use of system resources and enables developers to write concurrent programs without worrying about low-level thread management.

## Author
Johann Ramírez
