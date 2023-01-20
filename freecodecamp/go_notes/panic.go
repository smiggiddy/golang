package main

import (
	"fmt"
	"log"
)

func main() {
	// go doesn't have exceptions, but has panics
	defer fmt.Println("Test")
	panicking()
	fmt.Println("You made it through, basic program")
	// panic("Something bad happened with this")
	// panics happen after deferred statements in go
}
func panicking() {
	fmt.Println("Somethings about to go down bro")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			//through another panic if you need it to end for real
			// panic(err)
		}
	}()
	panic("Too many tech layoffs")
	fmt.Println("Keep focused though, this aint gonna run")
}
