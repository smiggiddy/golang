package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan string, 5)

	for i := 0; i < 5; i++ {
		ch <- strconv.Itoa(i)
	}
	wg.Add(1)

	// close(ch)

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println(msg)
			default:

			}
		}
	}()
	wg.Wait()

	fmt.Println("Program is done")

}
