package main

import "time"

func main() {
	done := make(chan bool)
	go func() {
		println("Sleeping...")
		time.Sleep(5 * time.Second)
		println("Wakeup!!")
		close(done)

	}()
	for {
		select {
		case <-done:
			return
		default:
			println("waiting...")
		}
	}
}
