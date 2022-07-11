package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})

		go func() {
			fmt.Println("working...")
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
			fmt.Println("done...")
			close(orDone)
		}()
		return orDone
	}
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(time.Hour),
		sig(time.Hour),
		sig(time.Hour),
		sig(time.Hour),
		sig(time.Hour),
		sig(time.Hour),
		sig(time.Hour),
		sig(1*time.Minute),
		sig(1*time.Hour),
		sig(10*time.Minute),
	)
	fmt.Printf("done after: %v", time.Since(start))
}
