package main

import (
	"fmt"
	"time"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i < 5; i++ {

				resultStream <- i
				fmt.Printf("Sent: %d\n", i)
			}
			fmt.Println("Sent all values")
		}()
		return resultStream
	}

	resultStream := chanOwner()

	time.Sleep(2 * time.Second)
	fmt.Println("resultStream")

	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving")
}
