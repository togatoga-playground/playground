package main

import (
	"fmt"
)

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, v := range integers {
				select {
				case <-done:
					return
				case intStream <- v:
				}
			}
		}()
		return intStream
	}

	operator := func(done <-chan interface{}, intStream <-chan int, f func(int) int) <-chan int {
		resultStream := make(chan int)
		go func() {
			defer close(resultStream)
			for value := range intStream {
				select {
				case <-done:
					return
				case resultStream <- f(value):
				}
			}
		}()
		return resultStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		return operator(done, intStream, func(x int) int { return x * multiplier })
	}
	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		return operator(done, intStream, func(x int) int { return x + additive })
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4, 5, 6, 7)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}

}
