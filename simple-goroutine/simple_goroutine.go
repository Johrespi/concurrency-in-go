package main

import (
	"fmt"
)

func main() {
	go someBehaviour()

	fmt.Println("Hello from main")
	// continue doing other things
}

func someBehaviour(){
	fmt.Println("Hello from a goroutine")
}
