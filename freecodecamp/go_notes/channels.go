package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// channels are used to transition between go routines
	fmt.Println("Channels in Go :)")

	// declaring a channel that only can send INT
	ch := make(chan int, 10)
	wg.Add(2)

	go func(ch <-chan int) {
		// i := <-ch
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

}
