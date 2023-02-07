package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	ch := make(chan string, 20)

	for i := 0; i < 5; i++ {
		ch <- strconv.Itoa(i)
	}
	var wg = sync.WaitGroup{}
	wg.Add(1)

	go func(ch <-chan string) {
		for {
			select {
			case msg := <-ch:
				fmt.Println(msg)
			default:
			}
		}
	}(ch)

	wg.Wait()
	wg.Done()

}
