package main

import (
	"fmt"
)

// upper case letter means it will be exported
// internal constant use lowercase

const (
	a = iota
	b
)

func main() {
	const myConst int = 133343
	fmt.Printf("%v, %T\n", myConst, myConst)

	// constants should not be executable, compiler will throw an error
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v\n", b)
}
