package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {

			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		defer close(c)

	}()
	return c
}

func merge(a, b <-chan int) <-chan int {

	c := make(chan int)
	go func() {
		var ok1, ok2 bool
		v := 0
		defer wg.Done()
		for {

			select {
			case v, ok1 = <-a:
				c <- v
			case v, ok2 = <-b:
				c <- v

			}

		}

	}()
	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

func main() {
	wg.Add(1)
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
