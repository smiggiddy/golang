package main

import (
	"fmt"
)

func main() {
	fmt.Print("Basic\n")

	//for loop can use i = i + 2
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	s := [3]int{1, 2, 3}

	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)

	for i := 0; i < 10; i++ {
		for {
			fmt.Println("Inner loop")
			break
		}
		break

	}
}
