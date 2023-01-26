package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// creates go routine aka a green thread
	go sayHello()
	wg.Add(2)
	// time.Sleep(100 * time.Microsecond)

	var msg = "Yup"
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}

// how to check for race conditions?
// go run -race <program name>
