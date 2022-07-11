package main

import (
	"fmt"
	"math/rand"
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

	// repeat
	_ = func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, value := range values {
					select {
					case <-done:
						return
					case valueStream <- value:
					}
				}
			}
		}()
		return valueStream
	}
	repatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	{
		done := make(chan interface{})
		defer close(done)
		rand := func() interface{} {
			return rand.Int()
		}
		for num := range take(done, repatFn(done, rand), 10) {
			fmt.Println(num)
		}
	}
}
